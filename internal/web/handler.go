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
