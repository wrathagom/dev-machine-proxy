package projects

import (
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// Project represents a project folder with git status
type Project struct {
	Name         string    `json:"name"`
	Path         string    `json:"path"`
	IsGit        bool      `json:"isGit"`
	Branch       string    `json:"branch,omitempty"`
	ChangedFiles int       `json:"changedFiles,omitempty"`
	Unpushed     int       `json:"unpushed,omitempty"`
	Ahead        int       `json:"ahead,omitempty"`
	Behind       int       `json:"behind,omitempty"`
	LastModified time.Time `json:"lastModified"`
	Tags         []string  `json:"tags"`
}

// Scanner scans a directory for projects
type Scanner struct {
	projectsDir string
}

// NewScanner creates a new project scanner
func NewScanner(projectsDir string) *Scanner {
	return &Scanner{projectsDir: projectsDir}
}

// Scan finds all projects in the configured directory
func (s *Scanner) Scan() []Project {
	if s.projectsDir == "" {
		return nil
	}

	entries, err := os.ReadDir(s.projectsDir)
	if err != nil {
		return nil
	}

	var projects []Project
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		// Skip hidden directories
		if strings.HasPrefix(entry.Name(), ".") {
			continue
		}

		path := filepath.Join(s.projectsDir, entry.Name())
		project := s.analyzeProject(entry.Name(), path)
		projects = append(projects, project)
	}

	return projects
}

func (s *Scanner) analyzeProject(name, path string) Project {
	p := Project{
		Name: name,
		Path: path,
		Tags: []string{},
	}

	// Check if it's a git repo
	gitDir := filepath.Join(path, ".git")
	if info, err := os.Stat(gitDir); err == nil && info.IsDir() {
		p.IsGit = true
		p.Tags = append(p.Tags, "git")
		s.getGitStatus(&p)
	}

	// Get last modified time
	p.LastModified = s.getLastModified(path)

	// Detect project type
	s.detectProjectType(&p)

	return p
}

func (s *Scanner) getGitStatus(p *Project) {
	// Get current branch
	if out, err := runGitCommand(p.Path, "rev-parse", "--abbrev-ref", "HEAD"); err == nil {
		p.Branch = strings.TrimSpace(out)
	}

	// Get changed files count (staged + unstaged + untracked)
	if out, err := runGitCommand(p.Path, "status", "--porcelain"); err == nil {
		lines := strings.Split(strings.TrimSpace(out), "\n")
		if out != "" && len(lines) > 0 {
			p.ChangedFiles = len(lines)
		}
	}

	// Get unpushed commits
	if out, err := runGitCommand(p.Path, "log", "--oneline", "@{upstream}..HEAD"); err == nil {
		lines := strings.Split(strings.TrimSpace(out), "\n")
		if out != "" && len(lines) > 0 && lines[0] != "" {
			p.Unpushed = len(lines)
		}
	}

	// Get ahead/behind counts
	if out, err := runGitCommand(p.Path, "rev-list", "--left-right", "--count", "HEAD...@{upstream}"); err == nil {
		parts := strings.Fields(strings.TrimSpace(out))
		if len(parts) == 2 {
			p.Ahead, _ = strconv.Atoi(parts[0])
			p.Behind, _ = strconv.Atoi(parts[1])
		}
	}
}

func (s *Scanner) getLastModified(path string) time.Time {
	var latest time.Time

	// Check .git/index for git repos (most accurate for git activity)
	gitIndex := filepath.Join(path, ".git", "index")
	if info, err := os.Stat(gitIndex); err == nil {
		return info.ModTime()
	}

	// Fall back to checking the directory itself
	if info, err := os.Stat(path); err == nil {
		latest = info.ModTime()
	}

	return latest
}

func (s *Scanner) detectProjectType(p *Project) {
	// Check for various project files
	checks := map[string]string{
		"package.json":      "node",
		"go.mod":            "go",
		"Cargo.toml":        "rust",
		"requirements.txt":  "python",
		"pyproject.toml":    "python",
		"Gemfile":           "ruby",
		"composer.json":     "php",
		"pom.xml":           "java",
		"build.gradle":      "java",
		"docker-compose.yml": "docker",
		"docker-compose.yaml": "docker",
		"Dockerfile":        "docker",
	}

	for file, tag := range checks {
		if _, err := os.Stat(filepath.Join(p.Path, file)); err == nil {
			// Avoid duplicate tags
			if !contains(p.Tags, tag) {
				p.Tags = append(p.Tags, tag)
			}
		}
	}
}

func runGitCommand(dir string, args ...string) (string, error) {
	cmd := exec.Command("git", args...)
	cmd.Dir = dir
	out, err := cmd.Output()
	return string(out), err
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
