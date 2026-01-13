package web

import (
	"encoding/json"
	"net"
	"net/http"
	"net/url"
	"strings"

	"dev-machine-proxy/internal/config"
	"dev-machine-proxy/internal/discovery"
	"dev-machine-proxy/internal/system"
	"dev-machine-proxy/internal/terminal"
)

// Handler serves the web dashboard
type Handler struct {
	discoverer  *discovery.Discoverer
	configMgr   *config.Manager
	sysMonitor  *system.Monitor
	termHandler *terminal.Handler
	mux         *http.ServeMux
}

// NewHandler creates a new web handler
func NewHandler(d *discovery.Discoverer, cfg *config.Manager, mon *system.Monitor) *Handler {
	h := &Handler{
		discoverer:  d,
		configMgr:   cfg,
		sysMonitor:  mon,
		termHandler: terminal.NewHandler(),
		mux:         http.NewServeMux(),
	}

	h.mux.HandleFunc("/", h.handleIndex)
	h.mux.HandleFunc("/config", h.handleConfigPage)
	h.mux.HandleFunc("/favicon.ico", h.handleFavicon)
	h.mux.HandleFunc("/api/services", h.handleAPIServices)
	h.mux.HandleFunc("/api/config", h.handleAPIConfig)
	h.mux.HandleFunc("/api/themes", h.handleAPIThemes)
	h.mux.HandleFunc("/api/stats", h.handleAPIStats)
	h.mux.HandleFunc("/ws/terminal", h.termHandler.ServeWS)

	return h
}

// ServeHTTP implements http.Handler
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mux.ServeHTTP(w, r)
}

// handleAPIServices returns services as JSON
func (h *Handler) handleAPIServices(w http.ResponseWriter, r *http.Request) {
	services := h.discoverer.GetServices()
	services = adjustServiceURLs(services, r)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(services)
}

// handleAPIConfig handles GET and POST for config
func (h *Handler) handleAPIConfig(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		json.NewEncoder(w).Encode(h.configMgr.Get())

	case http.MethodPost:
		var cfg config.Config
		if err := json.NewDecoder(r.Body).Decode(&cfg); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := h.configMgr.Update(cfg); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(cfg)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// handleAPIThemes returns available themes
func (h *Handler) handleAPIThemes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(config.AvailableThemes())
}

// handleIndex serves the main dashboard page
func (h *Handler) handleIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// Inject custom head HTML from config
	cfg := h.configMgr.Get()
	html := strings.Replace(indexHTML, "{{CUSTOM_HEAD_HTML}}", cfg.CustomHeadHTML, 1)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(html))
}

// handleConfigPage serves the config page
func (h *Handler) handleConfigPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(configHTML))
}

// handleFavicon serves the favicon with theme colors
func (h *Handler) handleFavicon(w http.ResponseWriter, r *http.Request) {
	cfg := h.configMgr.Get()
	colors := getThemeColors(cfg.Theme)
	svg := generateFaviconSVG(colors.bg, colors.primary, colors.secondary)
	w.Header().Set("Content-Type", "image/svg+xml")
	w.Write([]byte(svg))
}

// handleAPIStats returns system stats history
func (h *Handler) handleAPIStats(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(h.sysMonitor.GetHistory())
}

func adjustServiceURLs(services []discovery.Service, r *http.Request) []discovery.Service {
	host := requestHostname(r)
	if host == "" {
		return services
	}

	adjusted := make([]discovery.Service, len(services))
	copy(adjusted, services)

	for i := range adjusted {
		if adjusted[i].URL == "" {
			continue
		}
		if updated, ok := replaceLocalhostURL(adjusted[i].URL, host); ok {
			adjusted[i].URL = updated
		}
	}

	return adjusted
}

func requestHostname(r *http.Request) string {
	if r == nil || r.Host == "" {
		return ""
	}

	if strings.Contains(r.Host, ":") {
		if host, _, err := net.SplitHostPort(r.Host); err == nil {
			return host
		}
	}

	return r.Host
}

func replaceLocalhostURL(rawURL, host string) (string, bool) {
	parsed, err := url.Parse(rawURL)
	if err != nil || parsed.Host == "" {
		return "", false
	}

	if !isLoopbackHost(parsed.Hostname()) {
		return "", false
	}

	if port := parsed.Port(); port != "" {
		parsed.Host = net.JoinHostPort(host, port)
	} else {
		parsed.Host = host
	}

	return parsed.String(), true
}

func isLoopbackHost(host string) bool {
	if host == "localhost" {
		return true
	}

	ip := net.ParseIP(host)
	return ip != nil && ip.IsLoopback()
}

// themeColors holds the colors for favicon generation
type themeColors struct {
	bg        string
	primary   string
	secondary string
}

// getThemeColors returns the colors for a given theme
func getThemeColors(theme string) themeColors {
	themes := map[string]themeColors{
		"cyberpunk":            {bg: "#1a1a2e", primary: "#00d9ff", secondary: "#00ff88"},
		"catppuccin-mocha":     {bg: "#1e1e2e", primary: "#cba6f7", secondary: "#f5c2e7"},
		"catppuccin-macchiato": {bg: "#24273a", primary: "#c6a0f6", secondary: "#f5bde6"},
		"solarized-dark":       {bg: "#002b36", primary: "#268bd2", secondary: "#2aa198"},
		"solarized-light":      {bg: "#fdf6e3", primary: "#268bd2", secondary: "#2aa198"},
		"dracula":              {bg: "#282a36", primary: "#bd93f9", secondary: "#ff79c6"},
		"nord":                 {bg: "#2e3440", primary: "#88c0d0", secondary: "#81a1c1"},
		"gruvbox-dark":         {bg: "#282828", primary: "#fabd2f", secondary: "#fe8019"},
		"tokyo-night":          {bg: "#1a1b26", primary: "#7aa2f7", secondary: "#bb9af7"},
		"one-dark":             {bg: "#282c34", primary: "#61afef", secondary: "#c678dd"},
		"monokai":              {bg: "#272822", primary: "#a6e22e", secondary: "#f92672"},
	}

	if colors, ok := themes[theme]; ok {
		return colors
	}
	return themes["cyberpunk"]
}

// generateFaviconSVG creates a favicon SVG with the given colors
func generateFaviconSVG(bg, primary, secondary string) string {
	return `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 64 64">
  <rect width="64" height="64" rx="8" fill="` + bg + `"/>
  <path d="M16 20 L28 32 L16 44" stroke="` + primary + `" stroke-width="4" stroke-linecap="round" stroke-linejoin="round" fill="none"/>
  <line x1="32" y1="44" x2="48" y2="44" stroke="` + secondary + `" stroke-width="4" stroke-linecap="round"/>
</svg>`
}
