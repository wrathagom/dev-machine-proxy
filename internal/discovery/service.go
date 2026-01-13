package discovery

// Service represents a discovered service running on a port
type Service struct {
	Port        int      `json:"port"`
	Protocol    string   `json:"protocol"`    // tcp, udp
	Name        string   `json:"name"`        // Best guess at service name
	Description string   `json:"description"` // Additional context
	URL         string   `json:"url"`         // Clickable URL if HTTP-based
	Source      string   `json:"source"`      // How we discovered it: docker, process, port-scan
	Process     string   `json:"process"`     // Process name if available
	Container   string   `json:"container"`   // Docker container name if applicable
	Image       string   `json:"image"`       // Docker image if applicable
	ProjectPath string   `json:"projectPath"` // Path to project folder if found
	Tags        []string `json:"tags"`        // Additional tags for categorization
	IsHTTP      bool     `json:"isHttp"`      // Whether this appears to be an HTTP service
}

// KnownPorts maps common ports to their typical services
var KnownPorts = map[int]string{
	22:    "SSH",
	80:    "HTTP",
	443:   "HTTPS",
	1433:  "MSSQL",
	1521:  "Oracle DB",
	2375:  "Docker (unencrypted)",
	2376:  "Docker (TLS)",
	3000:  "Dev Server (Node/Rails/etc)",
	3306:  "MySQL",
	4200:  "Angular Dev Server",
	5000:  "Flask/ASP.NET Dev Server",
	5173:  "Vite Dev Server",
	5432:  "PostgreSQL",
	5672:  "RabbitMQ",
	6379:  "Redis",
	8000:  "Python Dev Server",
	8080:  "HTTP Alt / Tomcat",
	8081:  "HTTP Alt",
	8443:  "HTTPS Alt",
	8888:  "Jupyter Notebook",
	9000:  "PHP-FPM / Portainer",
	9090:  "Prometheus",
	9200:  "Elasticsearch",
	9999:  "Dev Machine Proxy",
	15672: "RabbitMQ Management",
	27017: "MongoDB",
}

// HTTPPorts are ports that are typically HTTP services worth probing
var HTTPPorts = map[int]bool{
	80: true, 443: true, 3000: true, 4200: true, 5000: true,
	5173: true, 8000: true, 8080: true, 8081: true, 8443: true,
	8888: true, 9000: true, 9090: true, 15672: true,
}
