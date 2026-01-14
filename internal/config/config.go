package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sync"
)

// Config holds user-configurable settings
type Config struct {
	Title           string          `json:"title"`
	Theme           string          `json:"theme"`
	RefreshInterval int             `json:"refreshInterval"` // in seconds
	TerminalFont    string          `json:"terminalFont"`    // CSS font-family for terminal
	CustomHeadHTML  string          `json:"customHeadHtml"`  // Custom HTML to inject in <head> (for fonts, etc.)
	Sections        SectionSettings `json:"sections"`        // Which sections to show
}

// SectionSettings controls visibility of dashboard sections
type SectionSettings struct {
	Performance bool `json:"performance"`
	Projects    bool `json:"projects"`
	Services    bool `json:"services"`
	Terminal    bool `json:"terminal"`
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
			Projects:    true,
			Services:    true,
			Terminal:    true,
		},
	}
}

// Manager handles loading and saving configuration
type Manager struct {
	config Config
	path   string
	mu     sync.RWMutex
}

// NewManager creates a config manager
func NewManager() *Manager {
	m := &Manager{
		config: DefaultConfig(),
		path:   getConfigPath(),
	}
	m.Load()
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
