package discovery

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

// ProjectMatch represents a port reference found in a project
type ProjectMatch struct {
	ProjectPath string
	ProjectName string
	Port        int
	File        string
	Context     string // The line or context where we found the port
}

// ScanProjectsForPorts scans a directory of projects for port references
func ScanProjectsForPorts(projectsDir string, ports []int) ([]ProjectMatch, error) {
	if projectsDir == "" {
		return nil, nil
	}

	// Convert ports to a set for quick lookup
	portSet := make(map[int]bool)
	for _, p := range ports {
		portSet[p] = true
	}

	var matches []ProjectMatch

	// List top-level directories (each is a project)
	entries, err := os.ReadDir(projectsDir)
	if err != nil {
		return nil, fmt.Errorf("reading projects directory: %w", err)
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		projectPath := filepath.Join(projectsDir, entry.Name())
		projectMatches := scanProject(projectPath, entry.Name(), portSet)
		matches = append(matches, projectMatches...)
	}

	return matches, nil
}

// scanProject searches a single project directory for port references
func scanProject(projectPath, projectName string, ports map[int]bool) []ProjectMatch {
	var matches []ProjectMatch

	// Files likely to contain port configurations
	configFiles := []string{
		"docker-compose.yml",
		"docker-compose.yaml",
		"compose.yml",
		"compose.yaml",
		".env",
		".env.local",
		".env.development",
		"package.json",
		"Makefile",
		"Dockerfile",
		"config.json",
		"config.yaml",
		"config.yml",
		"settings.json",
		"appsettings.json",
		"appsettings.Development.json",
	}

	// Check each config file
	for _, configFile := range configFiles {
		filePath := filepath.Join(projectPath, configFile)
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			continue
		}

		fileMatches := scanFileForPorts(filePath, projectPath, projectName, ports)
		matches = append(matches, fileMatches...)
	}

	// Also check common config directories
	configDirs := []string{"config", ".config", "conf"}
	for _, dir := range configDirs {
		dirPath := filepath.Join(projectPath, dir)
		if info, err := os.Stat(dirPath); err == nil && info.IsDir() {
			filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
				if err != nil || info.IsDir() {
					return nil
				}
				if isConfigFile(info.Name()) {
					fileMatches := scanFileForPorts(path, projectPath, projectName, ports)
					matches = append(matches, fileMatches...)
				}
				return nil
			})
		}
	}

	return matches
}

// isConfigFile checks if a filename looks like a config file
func isConfigFile(name string) bool {
	extensions := []string{".yml", ".yaml", ".json", ".env", ".toml", ".conf", ".cfg"}
	for _, ext := range extensions {
		if strings.HasSuffix(strings.ToLower(name), ext) {
			return true
		}
	}
	return false
}

// scanFileForPorts reads a file and looks for port number references
func scanFileForPorts(filePath, projectPath, projectName string, ports map[int]bool) []ProjectMatch {
	var matches []ProjectMatch

	file, err := os.Open(filePath)
	if err != nil {
		return nil
	}
	defer file.Close()

	// Regex patterns for port references
	portPatterns := []*regexp.Regexp{
		regexp.MustCompile(`[Pp]ort["\s:=]+(\d{2,5})`),            // port: 3000, PORT=3000, "port": 3000
		regexp.MustCompile(`(\d{2,5}):(\d{2,5})`),                 // 8080:80 (docker port mapping)
		regexp.MustCompile(`localhost:(\d{2,5})`),                 // localhost:3000
		regexp.MustCompile(`127\.0\.0\.1:(\d{2,5})`),              // 127.0.0.1:3000
		regexp.MustCompile(`0\.0\.0\.0:(\d{2,5})`),                // 0.0.0.0:3000
		regexp.MustCompile(`["\s:=](\d{4,5})["\s,\n\r]`),          // Generic 4-5 digit numbers that might be ports
	}

	relPath, _ := filepath.Rel(projectPath, filePath)
	scanner := bufio.NewScanner(file)
	lineNum := 0

	for scanner.Scan() {
		lineNum++
		line := scanner.Text()

		for _, pattern := range portPatterns {
			submatches := pattern.FindAllStringSubmatch(line, -1)
			for _, submatch := range submatches {
				for i := 1; i < len(submatch); i++ {
					port, err := strconv.Atoi(submatch[i])
					if err != nil {
						continue
					}

					// Only include if this port is in our listening ports
					if !ports[port] {
						continue
					}

					// Avoid duplicates in same file
					isDupe := false
					for _, m := range matches {
						if m.Port == port && m.File == relPath {
							isDupe = true
							break
						}
					}
					if isDupe {
						continue
					}

					matches = append(matches, ProjectMatch{
						ProjectPath: projectPath,
						ProjectName: projectName,
						Port:        port,
						File:        relPath,
						Context:     strings.TrimSpace(line),
					})
				}
			}
		}
	}

	return matches
}

// FindProjectForPort finds the best project match for a given port
func FindProjectForPort(matches []ProjectMatch, port int) *ProjectMatch {
	for i, m := range matches {
		if m.Port == port {
			return &matches[i]
		}
	}
	return nil
}
