package usage

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
	"time"
)

const (
	defaultInterval   = 5 * time.Minute
	historyRetention  = 7 * 24 * time.Hour
	maxHistorySamples = 7 * 24 * 12 // 5-minute samples over 7 days
	commandTimeout    = 10 * time.Second
)

type UsagePoint struct {
	Timestamp time.Time `json:"timestamp"`
	Value     float64   `json:"value"`
}

type UsageHistory struct {
	ClaudeFiveHour []UsagePoint `json:"claudeFiveHour"`
	ClaudeSevenDay []UsagePoint `json:"claudeSevenDay"`
	CodexPrimary   []UsagePoint `json:"codexPrimary"`
	CodexSecondary []UsagePoint `json:"codexSecondary"`
}

type WindowUsage struct {
	UsedPercent    float64   `json:"usedPercent"`
	ResetAt        time.Time `json:"resetAt"`
	ResetInSeconds int64     `json:"resetInSeconds"`
}

type Forecast struct {
	RatePerHour      float64    `json:"ratePerHour"`
	HoursToExhaust   float64    `json:"hoursToExhaust"`
	ProjectedAtReset float64    `json:"projectedAtReset"`
	ExhaustAt        *time.Time `json:"exhaustAt,omitempty"`
	WillExhaust      bool       `json:"willExhaust"`
	BasisHours       float64    `json:"basisHours"`
}

type WindowStatus struct {
	Current  WindowUsage `json:"current"`
	Forecast Forecast    `json:"forecast"`
}

type ClaudeUsage struct {
	FiveHour WindowStatus `json:"fiveHour"`
	SevenDay WindowStatus `json:"sevenDay"`
}

type CodexUsage struct {
	Primary   WindowStatus  `json:"primary"`
	Secondary *WindowStatus `json:"secondary,omitempty"`
}

type UsageSnapshot struct {
	Timestamp time.Time    `json:"timestamp"`
	Claude    *ClaudeUsage `json:"claude,omitempty"`
	Codex     *CodexUsage  `json:"codex,omitempty"`
	Errors    []string     `json:"errors,omitempty"`
}

type UsageResponse struct {
	Latest  UsageSnapshot `json:"latest"`
	History UsageHistory  `json:"history"`
}

type Monitor struct {
	mu      sync.RWMutex
	history UsageHistory
	latest  UsageSnapshot
	path    string
}

func NewMonitor() *Monitor {
	m := &Monitor{
		path: getUsagePath(),
	}
	_ = m.load()
	return m
}

func (m *Monitor) Start(interval time.Duration) {
	if interval <= 0 {
		interval = defaultInterval
	}
	m.collect()
	go func() {
		ticker := time.NewTicker(interval)
		for range ticker.C {
			m.collect()
		}
	}()
}

func (m *Monitor) GetResponse() UsageResponse {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return UsageResponse{
		Latest:  m.latest,
		History: m.history,
	}
}

func (m *Monitor) collect() {
	now := time.Now()
	var errorsList []string

	claudeUsage, claudeErr := fetchClaudeUsage()
	if claudeErr != nil {
		errorsList = append(errorsList, fmt.Sprintf("claude-usage: %v", claudeErr))
	}

	codexUsage, codexErr := fetchCodexUsage()
	if codexErr != nil {
		errorsList = append(errorsList, fmt.Sprintf("codex-usage: %v", codexErr))
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	if claudeUsage != nil {
		m.history.ClaudeFiveHour = appendUsagePoint(m.history.ClaudeFiveHour, now, claudeUsage.FiveHour.Current.UsedPercent)
		m.history.ClaudeSevenDay = appendUsagePoint(m.history.ClaudeSevenDay, now, claudeUsage.SevenDay.Current.UsedPercent)
	}

	if codexUsage != nil {
		m.history.CodexPrimary = appendUsagePoint(m.history.CodexPrimary, now, codexUsage.Primary.Current.UsedPercent)
		if codexUsage.Secondary != nil {
			m.history.CodexSecondary = appendUsagePoint(m.history.CodexSecondary, now, codexUsage.Secondary.Current.UsedPercent)
		}
	}

	m.history = pruneHistory(m.history, now)

	if claudeUsage != nil {
		claudeUsage.FiveHour.Forecast = forecastWindow(m.history.ClaudeFiveHour, claudeUsage.FiveHour.Current, 3*time.Hour)
		claudeUsage.SevenDay.Forecast = forecastWindow(m.history.ClaudeSevenDay, claudeUsage.SevenDay.Current, 48*time.Hour)
	}
	if codexUsage != nil {
		codexUsage.Primary.Forecast = forecastWindow(m.history.CodexPrimary, codexUsage.Primary.Current, 3*time.Hour)
		if codexUsage.Secondary != nil {
			codexUsage.Secondary.Forecast = forecastWindow(m.history.CodexSecondary, codexUsage.Secondary.Current, 48*time.Hour)
		}
	}

	m.latest = UsageSnapshot{
		Timestamp: now,
		Claude:    claudeUsage,
		Codex:     codexUsage,
		Errors:    errorsList,
	}

	_ = m.save()
}

func appendUsagePoint(points []UsagePoint, ts time.Time, value float64) []UsagePoint {
	return append(points, UsagePoint{
		Timestamp: ts,
		Value:     value,
	})
}

func pruneHistory(history UsageHistory, now time.Time) UsageHistory {
	cutoff := now.Add(-historyRetention)
	history.ClaudeFiveHour = prunePoints(history.ClaudeFiveHour, cutoff)
	history.ClaudeSevenDay = prunePoints(history.ClaudeSevenDay, cutoff)
	history.CodexPrimary = prunePoints(history.CodexPrimary, cutoff)
	history.CodexSecondary = prunePoints(history.CodexSecondary, cutoff)
	return history
}

func prunePoints(points []UsagePoint, cutoff time.Time) []UsagePoint {
	pruned := points[:0]
	for _, pt := range points {
		if pt.Timestamp.After(cutoff) {
			pruned = append(pruned, pt)
		}
	}
	if len(pruned) > maxHistorySamples {
		pruned = pruned[len(pruned)-maxHistorySamples:]
	}
	return pruned
}

func forecastWindow(points []UsagePoint, current WindowUsage, lookback time.Duration) Forecast {
	if len(points) < 2 {
		return Forecast{RatePerHour: 0, HoursToExhaust: -1, ProjectedAtReset: current.UsedPercent}
	}

	cutoff := points[len(points)-1].Timestamp.Add(-lookback)
	windowPoints := points
	for i, pt := range points {
		if pt.Timestamp.After(cutoff) {
			windowPoints = points[i:]
			break
		}
	}
	if len(windowPoints) < 2 {
		windowPoints = points
	}

	first := windowPoints[0]
	last := windowPoints[len(windowPoints)-1]
	durationHours := last.Timestamp.Sub(first.Timestamp).Hours()
	if durationHours <= 0 {
		return Forecast{RatePerHour: 0, HoursToExhaust: -1, ProjectedAtReset: current.UsedPercent}
	}

	rate := (last.Value - first.Value) / durationHours
	now := last.Timestamp
	hoursToReset := current.ResetAt.Sub(now).Hours()
	projectedAtReset := current.UsedPercent + rate*hoursToReset

	forecast := Forecast{
		RatePerHour:      rate,
		HoursToExhaust:   -1,
		ProjectedAtReset: projectedAtReset,
		BasisHours:       durationHours,
	}

	if rate <= 0 {
		return forecast
	}

	if current.UsedPercent >= 100 {
		forecast.HoursToExhaust = 0
		forecast.ExhaustAt = &now
		forecast.WillExhaust = true
		return forecast
	}

	hoursToExhaust := (100 - current.UsedPercent) / rate
	forecast.HoursToExhaust = hoursToExhaust
	exhaustAt := now.Add(time.Duration(hoursToExhaust * float64(time.Hour)))
	forecast.ExhaustAt = &exhaustAt
	forecast.WillExhaust = exhaustAt.Before(current.ResetAt)

	return forecast
}

func fetchClaudeUsage() (*ClaudeUsage, error) {
	output, err := runCommand("claude-usage")
	if err != nil {
		return nil, err
	}

	var resp struct {
		FiveHour *struct {
			Utilization float64   `json:"utilization"`
			ResetsAt    time.Time `json:"resets_at"`
		} `json:"five_hour"`
		SevenDay *struct {
			Utilization float64   `json:"utilization"`
			ResetsAt    time.Time `json:"resets_at"`
		} `json:"seven_day"`
	}

	if err := decodeJSONPrefix(output, &resp); err != nil {
		return nil, err
	}
	if resp.FiveHour == nil || resp.SevenDay == nil {
		return nil, errors.New("missing usage data")
	}

	fiveHour := WindowUsage{
		UsedPercent:    resp.FiveHour.Utilization,
		ResetAt:        resp.FiveHour.ResetsAt,
		ResetInSeconds: int64(time.Until(resp.FiveHour.ResetsAt).Seconds()),
	}
	sevenDay := WindowUsage{
		UsedPercent:    resp.SevenDay.Utilization,
		ResetAt:        resp.SevenDay.ResetsAt,
		ResetInSeconds: int64(time.Until(resp.SevenDay.ResetsAt).Seconds()),
	}

	return &ClaudeUsage{
		FiveHour: WindowStatus{Current: fiveHour},
		SevenDay: WindowStatus{Current: sevenDay},
	}, nil
}

func fetchCodexUsage() (*CodexUsage, error) {
	output, err := runCommand("codex-usage")
	if err != nil {
		return nil, err
	}

	var resp struct {
		RateLimit struct {
			PrimaryWindow *struct {
				UsedPercent        float64 `json:"used_percent"`
				ResetAt            int64   `json:"reset_at"`
				ResetAfterSeconds  int64   `json:"reset_after_seconds"`
				LimitWindowSeconds int64   `json:"limit_window_seconds"`
			} `json:"primary_window"`
			SecondaryWindow *struct {
				UsedPercent        float64 `json:"used_percent"`
				ResetAt            int64   `json:"reset_at"`
				ResetAfterSeconds  int64   `json:"reset_after_seconds"`
				LimitWindowSeconds int64   `json:"limit_window_seconds"`
			} `json:"secondary_window"`
		} `json:"rate_limit"`
	}

	if err := decodeJSONPrefix(output, &resp); err != nil {
		return nil, err
	}
	if resp.RateLimit.PrimaryWindow == nil {
		return nil, errors.New("missing primary window")
	}

	primaryReset := time.Unix(resp.RateLimit.PrimaryWindow.ResetAt, 0)
	primary := WindowUsage{
		UsedPercent:    resp.RateLimit.PrimaryWindow.UsedPercent,
		ResetAt:        primaryReset,
		ResetInSeconds: resp.RateLimit.PrimaryWindow.ResetAfterSeconds,
	}

	codex := &CodexUsage{
		Primary: WindowStatus{Current: primary},
	}

	if resp.RateLimit.SecondaryWindow != nil {
		secondaryReset := time.Unix(resp.RateLimit.SecondaryWindow.ResetAt, 0)
		secondary := WindowUsage{
			UsedPercent:    resp.RateLimit.SecondaryWindow.UsedPercent,
			ResetAt:        secondaryReset,
			ResetInSeconds: resp.RateLimit.SecondaryWindow.ResetAfterSeconds,
		}
		codex.Secondary = &WindowStatus{Current: secondary}
	}

	return codex, nil
}

func runCommand(name string, args ...string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), commandTimeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, name, args...)
	output, err := cmd.CombinedOutput()
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}
	if err != nil {
		return nil, fmt.Errorf("%v: %s", err, string(output))
	}
	return output, nil
}

func decodeJSONPrefix(output []byte, target interface{}) error {
	trimmed := bytes.TrimSpace(output)
	start := bytes.IndexByte(trimmed, '{')
	if start == -1 {
		return errors.New("no JSON found")
	}

	decoder := json.NewDecoder(bytes.NewReader(trimmed[start:]))
	if err := decoder.Decode(target); err != nil {
		return err
	}
	return nil
}

func getUsagePath() string {
	configDir, err := os.UserConfigDir()
	if err != nil {
		configDir = os.Getenv("HOME")
	}
	return filepath.Join(configDir, "dev-machine-proxy", "usage-history.json")
}

func (m *Monitor) load() error {
	data, err := os.ReadFile(m.path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	return json.Unmarshal(data, &m.history)
}

func (m *Monitor) save() error {
	data, err := json.MarshalIndent(m.history, "", "  ")
	if err != nil {
		return err
	}

	dir := filepath.Dir(m.path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	return os.WriteFile(m.path, data, 0644)
}
