package system

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

// Stats holds current system statistics
type Stats struct {
	CPUPercent    float64   `json:"cpuPercent"`
	MemoryTotal   uint64    `json:"memoryTotal"`
	MemoryUsed    uint64    `json:"memoryUsed"`
	MemoryPercent float64   `json:"memoryPercent"`
	Timestamp     time.Time `json:"timestamp"`
}

// History holds historical stats for charting
type History struct {
	Stats     []Stats `json:"stats"`
	MaxPoints int     `json:"-"`
}

// Monitor continuously collects system stats
type Monitor struct {
	history History
	mu      sync.RWMutex

	// For CPU calculation
	prevIdle  uint64
	prevTotal uint64
}

// NewMonitor creates a new system monitor
func NewMonitor(maxPoints int) *Monitor {
	return &Monitor{
		history: History{
			Stats:     make([]Stats, 0, maxPoints),
			MaxPoints: maxPoints,
		},
	}
}

// Start begins collecting stats at the specified interval
func (m *Monitor) Start(interval time.Duration) {
	// Collect initial stats
	m.collect()

	go func() {
		ticker := time.NewTicker(interval)
		for range ticker.C {
			m.collect()
		}
	}()
}

// GetHistory returns the stats history
func (m *Monitor) GetHistory() History {
	m.mu.RLock()
	defer m.mu.RUnlock()

	// Return a copy
	statsCopy := make([]Stats, len(m.history.Stats))
	copy(statsCopy, m.history.Stats)
	return History{
		Stats:     statsCopy,
		MaxPoints: m.history.MaxPoints,
	}
}

// GetCurrent returns the most recent stats
func (m *Monitor) GetCurrent() Stats {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if len(m.history.Stats) == 0 {
		return Stats{}
	}
	return m.history.Stats[len(m.history.Stats)-1]
}

func (m *Monitor) collect() {
	stats := Stats{
		Timestamp: time.Now(),
	}

	// Get CPU usage
	stats.CPUPercent = m.getCPUPercent()

	// Get memory usage
	stats.MemoryTotal, stats.MemoryUsed, stats.MemoryPercent = getMemoryStats()

	m.mu.Lock()
	m.history.Stats = append(m.history.Stats, stats)
	if len(m.history.Stats) > m.history.MaxPoints {
		m.history.Stats = m.history.Stats[1:]
	}
	m.mu.Unlock()
}

// getCPUPercent calculates CPU usage from /proc/stat
func (m *Monitor) getCPUPercent() float64 {
	file, err := os.Open("/proc/stat")
	if err != nil {
		return 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if !scanner.Scan() {
		return 0
	}

	line := scanner.Text()
	if !strings.HasPrefix(line, "cpu ") {
		return 0
	}

	fields := strings.Fields(line)
	if len(fields) < 5 {
		return 0
	}

	// Parse CPU times: user, nice, system, idle, iowait, irq, softirq, steal
	var total, idle uint64
	for i := 1; i < len(fields); i++ {
		val, err := strconv.ParseUint(fields[i], 10, 64)
		if err != nil {
			continue
		}
		total += val
		if i == 4 || i == 5 { // idle + iowait
			idle += val
		}
	}

	// Calculate percentage
	var cpuPercent float64
	if m.prevTotal > 0 {
		totalDelta := total - m.prevTotal
		idleDelta := idle - m.prevIdle
		if totalDelta > 0 {
			cpuPercent = float64(totalDelta-idleDelta) / float64(totalDelta) * 100
		}
	}

	m.prevTotal = total
	m.prevIdle = idle

	return cpuPercent
}

// getMemoryStats reads memory info from /proc/meminfo
func getMemoryStats() (total, used uint64, percent float64) {
	file, err := os.Open("/proc/meminfo")
	if err != nil {
		return 0, 0, 0
	}
	defer file.Close()

	var memTotal, memAvailable uint64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if len(fields) < 2 {
			continue
		}

		val, err := strconv.ParseUint(fields[1], 10, 64)
		if err != nil {
			continue
		}

		switch fields[0] {
		case "MemTotal:":
			memTotal = val * 1024 // Convert from KB to bytes
		case "MemAvailable:":
			memAvailable = val * 1024
		}

		if memTotal > 0 && memAvailable > 0 {
			break
		}
	}

	if memTotal > 0 {
		used = memTotal - memAvailable
		percent = float64(used) / float64(memTotal) * 100
	}

	return memTotal, used, percent
}
