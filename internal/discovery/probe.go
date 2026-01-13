package discovery

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
	"time"
)

// HTTPProbeResult contains information gathered from probing an HTTP endpoint
type HTTPProbeResult struct {
	IsHTTP      bool
	StatusCode  int
	Title       string
	Server      string
	PoweredBy   string
	ContentType string
	IsHTTPS     bool
}

// ProbeHTTP attempts to connect to a port via HTTP and gather information
func ProbeHTTP(port int) HTTPProbeResult {
	result := HTTPProbeResult{}

	// Try HTTPS first, then HTTP
	for _, scheme := range []string{"https", "http"} {
		url := fmt.Sprintf("%s://localhost:%d/", scheme, port)

		client := &http.Client{
			Timeout: 2 * time.Second,
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse // Don't follow redirects
			},
		}

		resp, err := client.Get(url)
		if err != nil {
			continue
		}
		defer resp.Body.Close()

		result.IsHTTP = true
		result.StatusCode = resp.StatusCode
		result.IsHTTPS = scheme == "https"

		// Extract headers
		result.Server = resp.Header.Get("Server")
		result.PoweredBy = resp.Header.Get("X-Powered-By")
		result.ContentType = resp.Header.Get("Content-Type")

		// Try to extract title from HTML
		if strings.Contains(result.ContentType, "text/html") {
			body, err := io.ReadAll(io.LimitReader(resp.Body, 64*1024)) // Read up to 64KB
			if err == nil {
				result.Title = extractTitle(string(body))
			}
		}

		return result
	}

	return result
}

// extractTitle pulls the <title> tag content from HTML
func extractTitle(html string) string {
	re := regexp.MustCompile(`(?i)<title[^>]*>([^<]+)</title>`)
	matches := re.FindStringSubmatch(html)
	if len(matches) > 1 {
		return strings.TrimSpace(matches[1])
	}
	return ""
}

// GuessServiceFromProbe attempts to identify a service based on HTTP probe results
func GuessServiceFromProbe(probe HTTPProbeResult, port int) string {
	// Check server header
	serverLower := strings.ToLower(probe.Server)
	poweredByLower := strings.ToLower(probe.PoweredBy)

	serverHints := map[string]string{
		"nginx":      "Nginx",
		"apache":     "Apache",
		"express":    "Express.js",
		"kestrel":    "ASP.NET",
		"gunicorn":   "Python (Gunicorn)",
		"uvicorn":    "Python (Uvicorn)",
		"werkzeug":   "Flask",
		"jetty":      "Jetty",
		"tomcat":     "Tomcat",
		"openresty":  "OpenResty",
		"caddy":      "Caddy",
		"traefik":    "Traefik",
	}

	for hint, name := range serverHints {
		if strings.Contains(serverLower, hint) || strings.Contains(poweredByLower, hint) {
			return name
		}
	}

	// Check title for common apps
	titleLower := strings.ToLower(probe.Title)
	titleHints := map[string]string{
		"grafana":    "Grafana",
		"prometheus": "Prometheus",
		"portainer":  "Portainer",
		"jenkins":    "Jenkins",
		"gitlab":     "GitLab",
		"gitea":      "Gitea",
		"nextcloud":  "Nextcloud",
		"home assistant": "Home Assistant",
		"jupyter":    "Jupyter",
		"pgadmin":    "pgAdmin",
		"adminer":    "Adminer",
		"phpMyAdmin": "phpMyAdmin",
		"mailhog":    "MailHog",
		"rabbitmq":   "RabbitMQ",
		"kibana":     "Kibana",
		"swagger":    "Swagger UI",
		"redoc":      "ReDoc",
	}

	for hint, name := range titleHints {
		if strings.Contains(titleLower, strings.ToLower(hint)) {
			return name
		}
	}

	// If we got a title, use it as a fallback
	if probe.Title != "" && len(probe.Title) < 50 {
		return probe.Title
	}

	return ""
}
