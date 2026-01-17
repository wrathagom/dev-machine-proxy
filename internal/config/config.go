package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// Config holds user-configurable settings
type Config struct {
	Title           string          `json:"title"`
	Theme           string          `json:"theme"`
	RefreshInterval int             `json:"refreshInterval"` // in seconds
	TerminalFont    string          `json:"terminalFont"`    // CSS font-family for terminal
	CustomHeadHTML  string          `json:"customHeadHtml"`  // Custom HTML to inject in <head> (for fonts, etc.)
	Sections        SectionSettings `json:"sections"`        // Which sections to show
	SectionOrder    []string        `json:"sectionOrder"`    // Order of sections on dashboard
}

// SectionSettings controls visibility of dashboard sections
type SectionSettings struct {
	Performance bool `json:"performance"`
	AIUsage     bool `json:"aiUsage"`
	Projects    bool `json:"projects"`
	Services    bool `json:"services"`
	Terminal    bool `json:"terminal"`
	DailyTasks  bool `json:"dailyTasks"`
}

// DailyTask represents a recurring daily task
type DailyTask struct {
	ID            string          `json:"id"`
	Name          string          `json:"name"`
	CreatedAt     time.Time       `json:"createdAt"`
	CurrentStreak int             `json:"currentStreak"`
	LongestStreak int             `json:"longestStreak"`
	Completions   map[string]bool `json:"completions"` // date string (YYYY-MM-DD) -> completed
}

// DailyTasksData holds all daily tasks data
type DailyTasksData struct {
	Tasks []DailyTask `json:"tasks"`
}

// DefaultConfig returns the default configuration
func DefaultConfig() Config {
	return Config{
		Title:           "Dev Machine Services",
		Theme:           "cyberpunk",
		RefreshInterval: 30,
		TerminalFont:    "MesloLGS NF",
		CustomHeadHTML: `<!-- MesloLGS NF for Powerlevel10k terminal icons -->
<style>
@font-face {
    font-family: 'MesloLGS NF';
    src: url('https://cdn.jsdelivr.net/gh/romkatv/powerlevel10k-media/MesloLGS%20NF%20Regular.ttf') format('truetype');
}
</style>`,
		Sections: SectionSettings{
			Performance: true,
			AIUsage:     true,
			Projects:    true,
			Services:    true,
			Terminal:    true,
			DailyTasks:  true,
		},
		SectionOrder: []string{"performance", "aiUsage", "projects", "services", "dailyTasks", "terminal"},
	}
}

// DefaultSectionOrder returns the default order of sections
func DefaultSectionOrder() []string {
	return []string{"performance", "aiUsage", "projects", "services", "dailyTasks", "terminal"}
}

// Manager handles loading and saving configuration
type Manager struct {
	config     Config
	path       string
	dailyTasks DailyTasksData
	tasksPath  string
	mu         sync.RWMutex
}

// NewManager creates a config manager
func NewManager() *Manager {
	m := &Manager{
		config:    DefaultConfig(),
		path:      getConfigPath(),
		tasksPath: getDailyTasksPath(),
	}
	m.Load()
	m.LoadDailyTasks()
	return m
}

// getConfigPath returns the path to the config file
func getConfigPath() string {
	configDir, err := os.UserConfigDir()
	if err != nil {
		configDir = os.Getenv("HOME")
	}
	return filepath.Join(configDir, "dev-machine-proxy", "config.json")
}

// Load reads config from disk
func (m *Manager) Load() error {
	m.mu.Lock()
	defer m.mu.Unlock()

	data, err := os.ReadFile(m.path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil // Use defaults
		}
		return err
	}

	return json.Unmarshal(data, &m.config)
}

// Save writes config to disk
func (m *Manager) Save() error {
	m.mu.RLock()
	data, err := json.MarshalIndent(m.config, "", "  ")
	m.mu.RUnlock()
	if err != nil {
		return err
	}

	// Ensure directory exists
	dir := filepath.Dir(m.path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	return os.WriteFile(m.path, data, 0644)
}

// Get returns the current config
func (m *Manager) Get() Config {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.config
}

// Update applies changes to the config
func (m *Manager) Update(cfg Config) error {
	m.mu.Lock()
	m.config = cfg
	m.mu.Unlock()
	return m.Save()
}

// AvailableThemes returns the list of supported themes
func AvailableThemes() []Theme {
	return []Theme{
		{ID: "cyberpunk", Name: "Cyberpunk", Description: "Neon blues and cyans on dark background"},
		{ID: "catppuccin-mocha", Name: "Catppuccin Mocha", Description: "Soothing pastel theme with warm tones"},
		{ID: "catppuccin-macchiato", Name: "Catppuccin Macchiato", Description: "Catppuccin with slightly lighter base"},
		{ID: "solarized-dark", Name: "Solarized Dark", Description: "Classic Solarized dark palette"},
		{ID: "solarized-light", Name: "Solarized Light", Description: "Classic Solarized light palette"},
		{ID: "dracula", Name: "Dracula", Description: "Dark theme with vibrant colors"},
		{ID: "nord", Name: "Nord", Description: "Arctic, north-bluish color palette"},
		{ID: "gruvbox-dark", Name: "Gruvbox Dark", Description: "Retro groove with warm colors"},
		{ID: "tokyo-night", Name: "Tokyo Night", Description: "Clean dark theme inspired by Tokyo lights"},
		{ID: "one-dark", Name: "One Dark", Description: "Atom's iconic One Dark theme"},
		{ID: "monokai", Name: "Monokai", Description: "Classic Sublime Text theme"},
	}
}

// Theme represents a color theme option
type Theme struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// getDailyTasksPath returns the path to the daily tasks file
func getDailyTasksPath() string {
	configDir, err := os.UserConfigDir()
	if err != nil {
		configDir = os.Getenv("HOME")
	}
	return filepath.Join(configDir, "dev-machine-proxy", "daily-tasks.json")
}

// LoadDailyTasks reads daily tasks from disk
func (m *Manager) LoadDailyTasks() error {
	m.mu.Lock()
	defer m.mu.Unlock()

	data, err := os.ReadFile(m.tasksPath)
	if err != nil {
		if os.IsNotExist(err) {
			m.dailyTasks = DailyTasksData{Tasks: []DailyTask{}}
			return nil
		}
		return err
	}

	if err := json.Unmarshal(data, &m.dailyTasks); err != nil {
		return err
	}

	// Update streaks on load
	m.updateAllStreaks()
	return nil
}

// SaveDailyTasks writes daily tasks to disk
func (m *Manager) SaveDailyTasks() error {
	m.mu.RLock()
	data, err := json.MarshalIndent(m.dailyTasks, "", "  ")
	m.mu.RUnlock()
	if err != nil {
		return err
	}

	dir := filepath.Dir(m.tasksPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	return os.WriteFile(m.tasksPath, data, 0644)
}

// GetDailyTasks returns all daily tasks with today's completion status
func (m *Manager) GetDailyTasks() []DailyTask {
	m.mu.RLock()
	defer m.mu.RUnlock()

	// Update streaks before returning
	m.updateAllStreaksLocked()

	result := make([]DailyTask, len(m.dailyTasks.Tasks))
	copy(result, m.dailyTasks.Tasks)
	return result
}

// AddDailyTask adds a new daily task
func (m *Manager) AddDailyTask(name string) (DailyTask, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	task := DailyTask{
		ID:          generateID(),
		Name:        name,
		CreatedAt:   time.Now(),
		Completions: make(map[string]bool),
	}

	m.dailyTasks.Tasks = append(m.dailyTasks.Tasks, task)

	if err := m.saveDailyTasksLocked(); err != nil {
		return DailyTask{}, err
	}

	return task, nil
}

// UpdateDailyTask updates an existing task's name
func (m *Manager) UpdateDailyTask(id, name string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i := range m.dailyTasks.Tasks {
		if m.dailyTasks.Tasks[i].ID == id {
			m.dailyTasks.Tasks[i].Name = name
			return m.saveDailyTasksLocked()
		}
	}
	return nil
}

// DeleteDailyTask removes a daily task
func (m *Manager) DeleteDailyTask(id string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i := range m.dailyTasks.Tasks {
		if m.dailyTasks.Tasks[i].ID == id {
			m.dailyTasks.Tasks = append(m.dailyTasks.Tasks[:i], m.dailyTasks.Tasks[i+1:]...)
			return m.saveDailyTasksLocked()
		}
	}
	return nil
}

// ToggleDailyTaskCompletion toggles completion for today
func (m *Manager) ToggleDailyTaskCompletion(id string) (bool, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	today := TodayString()

	for i := range m.dailyTasks.Tasks {
		if m.dailyTasks.Tasks[i].ID == id {
			if m.dailyTasks.Tasks[i].Completions == nil {
				m.dailyTasks.Tasks[i].Completions = make(map[string]bool)
			}

			// Toggle completion
			completed := !m.dailyTasks.Tasks[i].Completions[today]
			m.dailyTasks.Tasks[i].Completions[today] = completed

			// Update streak
			m.updateStreak(&m.dailyTasks.Tasks[i])

			if err := m.saveDailyTasksLocked(); err != nil {
				return false, err
			}
			return completed, nil
		}
	}
	return false, nil
}

// updateAllStreaks updates streaks for all tasks (must hold lock)
func (m *Manager) updateAllStreaks() {
	for i := range m.dailyTasks.Tasks {
		m.updateStreak(&m.dailyTasks.Tasks[i])
	}
}

// updateAllStreaksLocked updates streaks without acquiring lock (caller must hold RLock)
func (m *Manager) updateAllStreaksLocked() {
	// Note: This is called from GetDailyTasks which has RLock
	// We can't modify during RLock, so this is just for reading
}

// updateStreak calculates the current streak for a task
func (m *Manager) updateStreak(task *DailyTask) {
	if task.Completions == nil {
		task.CurrentStreak = 0
		return
	}

	streak := 0
	date := time.Now()

	// Check today first - if not completed, start from yesterday
	todayStr := date.Format("2006-01-02")
	if !task.Completions[todayStr] {
		date = date.AddDate(0, 0, -1)
	}

	// Count consecutive days backwards
	for {
		dateStr := date.Format("2006-01-02")
		if !task.Completions[dateStr] {
			break
		}
		streak++
		date = date.AddDate(0, 0, -1)
	}

	task.CurrentStreak = streak
	if streak > task.LongestStreak {
		task.LongestStreak = streak
	}
}

// saveDailyTasksLocked saves without acquiring lock (caller must hold lock)
func (m *Manager) saveDailyTasksLocked() error {
	data, err := json.MarshalIndent(m.dailyTasks, "", "  ")
	if err != nil {
		return err
	}

	dir := filepath.Dir(m.tasksPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	return os.WriteFile(m.tasksPath, data, 0644)
}

// TodayString returns today's date as YYYY-MM-DD
func TodayString() string {
	return time.Now().Format("2006-01-02")
}

// generateID creates a simple unique ID
func generateID() string {
	return time.Now().Format("20060102150405.000000000")
}
