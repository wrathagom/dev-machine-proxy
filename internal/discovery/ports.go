package discovery

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ListeningPort represents a port that's currently listening
type ListeningPort struct {
	Port    int
	PID     int
	Process string
}

// GetListeningPorts reads /proc/net/tcp to find all listening TCP ports
func GetListeningPorts() ([]ListeningPort, error) {
	ports := []ListeningPort{}
	seen := make(map[int]bool)

	// Read TCP ports (IPv4)
	tcpPorts, err := parseProcNet("/proc/net/tcp")
	if err != nil {
		return nil, fmt.Errorf("reading /proc/net/tcp: %w", err)
	}
	for _, p := range tcpPorts {
		if !seen[p.Port] {
			ports = append(ports, p)
			seen[p.Port] = true
		}
	}

	// Read TCP6 ports (IPv6)
	tcp6Ports, err := parseProcNet("/proc/net/tcp6")
	if err != nil {
		// IPv6 might not be available, not fatal
		return ports, nil
	}
	for _, p := range tcp6Ports {
		if !seen[p.Port] {
			ports = append(ports, p)
			seen[p.Port] = true
		}
	}

	// Enrich with process names
	enrichWithProcessNames(ports)

	return ports, nil
}

// parseProcNet parses /proc/net/tcp or /proc/net/tcp6 format
func parseProcNet(path string) ([]ListeningPort, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var ports []ListeningPort
	scanner := bufio.NewScanner(file)

	// Skip header line
	scanner.Scan()

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if len(fields) < 10 {
			continue
		}

		// Field 1 is local_address (hex IP:port)
		// Field 3 is state (0A = LISTEN)
		// Field 7 is uid
		// Field 9 is inode

		state := fields[3]
		if state != "0A" { // 0A = TCP_LISTEN
			continue
		}

		// Parse local address
		localAddr := fields[1]
		parts := strings.Split(localAddr, ":")
		if len(parts) != 2 {
			continue
		}

		portHex := parts[1]
		port64, err := strconv.ParseInt(portHex, 16, 32)
		if err != nil {
			continue
		}

		inode := fields[9]
		pid := findPIDByInode(inode)

		ports = append(ports, ListeningPort{
			Port: int(port64),
			PID:  pid,
		})
	}

	return ports, scanner.Err()
}

// findPIDByInode looks through /proc/*/fd/* to find the process owning a socket inode
func findPIDByInode(inode string) int {
	procDirs, err := os.ReadDir("/proc")
	if err != nil {
		return 0
	}

	socketLink := fmt.Sprintf("socket:[%s]", inode)

	for _, dir := range procDirs {
		if !dir.IsDir() {
			continue
		}

		pid, err := strconv.Atoi(dir.Name())
		if err != nil {
			continue // Not a PID directory
		}

		fdPath := fmt.Sprintf("/proc/%d/fd", pid)
		fds, err := os.ReadDir(fdPath)
		if err != nil {
			continue // Can't read this process's fds
		}

		for _, fd := range fds {
			linkPath := fmt.Sprintf("%s/%s", fdPath, fd.Name())
			link, err := os.Readlink(linkPath)
			if err != nil {
				continue
			}

			if link == socketLink {
				return pid
			}
		}
	}

	return 0
}

// enrichWithProcessNames adds process names to listening ports
func enrichWithProcessNames(ports []ListeningPort) {
	for i := range ports {
		if ports[i].PID == 0 {
			continue
		}

		commPath := fmt.Sprintf("/proc/%d/comm", ports[i].PID)
		data, err := os.ReadFile(commPath)
		if err != nil {
			continue
		}

		ports[i].Process = strings.TrimSpace(string(data))
	}
}
