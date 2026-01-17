package system

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

// Process holds information about a running process
type Process struct {
	PID        int     `json:"pid"`
	Name       string  `json:"name"`
	CPUPercent float64 `json:"cpuPercent"`
	MemPercent float64 `json:"memPercent"`
	MemoryMB   float64 `json:"memoryMB"`
}

// Stats holds current system statistics
type Stats struct {
	CPUPercent    float64   `json:"cpuPercent"`
	MemoryTotal   uint64    `json:"memoryTotal"`
	MemoryUsed    uint64    `json:"memoryUsed"`
	MemoryPercent float64   `json:"memoryPercent"`
	Timestamp     time.Time `json:"timestamp"`
	TopCPU        []Process `json:"topCPU"`
	TopMemory     []Process `json:"topMemory"`
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

	// For per-process CPU calculation
	prevProcCPU map[int]uint64
	prevTime    time.Time
}

// NewMonitor creates a new system monitor
func NewMonitor(maxPoints int) *Monitor {
	return &Monitor{
		history: History{
			Stats:     make([]Stats, 0, maxPoints),
			MaxPoints: maxPoints,
		},
		prevProcCPU: make(map[int]uint64),
		prevTime:    time.Now(),
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

	// Get top processes
	stats.TopCPU, stats.TopMemory = m.getTopProcesses(3, stats.MemoryTotal)

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

// getTopProcesses returns top N processes by CPU and memory
func (m *Monitor) getTopProcesses(n int, totalMem uint64) (topCPU, topMem []Process) {
	now := time.Now()
	elapsed := now.Sub(m.prevTime).Seconds()
	if elapsed <= 0 {
		elapsed = 1
	}

	// Get number of CPU cores for percentage calculation
	numCPU := float64(getNumCPU())

	procDir, err := os.Open("/proc")
	if err != nil {
		return nil, nil
	}
	defer procDir.Close()

	entries, err := procDir.Readdirnames(-1)
	if err != nil {
		return nil, nil
	}

	var processes []Process
	newProcCPU := make(map[int]uint64)

	for _, entry := range entries {
		pid, err := strconv.Atoi(entry)
		if err != nil {
			continue // Not a PID directory
		}

		proc := Process{PID: pid}

		// Read process name from /proc/[pid]/comm
		commPath := filepath.Join("/proc", entry, "comm")
		if data, err := os.ReadFile(commPath); err == nil {
			proc.Name = strings.TrimSpace(string(data))
		} else {
			continue
		}

		// Read CPU time from /proc/[pid]/stat
		statPath := filepath.Join("/proc", entry, "stat")
		if data, err := os.ReadFile(statPath); err == nil {
			fields := strings.Fields(string(data))
			if len(fields) >= 15 {
				utime, _ := strconv.ParseUint(fields[13], 10, 64)
				stime, _ := strconv.ParseUint(fields[14], 10, 64)
				totalTime := utime + stime
				newProcCPU[pid] = totalTime

				// Calculate CPU percentage
				if prevTime, ok := m.prevProcCPU[pid]; ok {
					ticksPerSec := 100.0 // Usually 100 Hz (sysconf(_SC_CLK_TCK))
					cpuTime := float64(totalTime-prevTime) / ticksPerSec
					proc.CPUPercent = (cpuTime / elapsed) * 100 / numCPU
				}
			}
		}

		// Read memory from /proc/[pid]/statm
		statmPath := filepath.Join("/proc", entry, "statm")
		if data, err := os.ReadFile(statmPath); err == nil {
			fields := strings.Fields(string(data))
			if len(fields) >= 2 {
				// RSS is in pages, convert to bytes (page size is typically 4096)
				rss, _ := strconv.ParseUint(fields[1], 10, 64)
				memBytes := rss * 4096
				proc.MemoryMB = float64(memBytes) / (1024 * 1024)
				if totalMem > 0 {
					proc.MemPercent = float64(memBytes) / float64(totalMem) * 100
				}
			}
		}

		processes = append(processes, proc)
	}

	// Update previous CPU times
	m.prevProcCPU = newProcCPU
	m.prevTime = now

	// Sort by CPU and get top N
	sort.Slice(processes, func(i, j int) bool {
		return processes[i].CPUPercent > processes[j].CPUPercent
	})
	if len(processes) > n {
		topCPU = make([]Process, n)
		copy(topCPU, processes[:n])
	} else {
		topCPU = processes
	}

	// Sort by memory and get top N
	sort.Slice(processes, func(i, j int) bool {
		return processes[i].MemPercent > processes[j].MemPercent
	})
	if len(processes) > n {
		topMem = make([]Process, n)
		copy(topMem, processes[:n])
	} else {
		topMem = processes
	}

	return topCPU, topMem
}

// getNumCPU returns the number of CPU cores
func getNumCPU() int {
	file, err := os.Open("/proc/cpuinfo")
	if err != nil {
		return 1
	}
	defer file.Close()

	count := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), "processor") {
			count++
		}
	}
	if count == 0 {
		return 1
	}
	return count
}

// formatBytes formats bytes to human readable string
func formatBytes(b uint64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := uint64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(b)/float64(div), "KMGTPE"[exp])
}
