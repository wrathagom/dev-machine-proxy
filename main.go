package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"dev-machine-proxy/internal/config"
	"dev-machine-proxy/internal/discovery"
	"dev-machine-proxy/internal/system"
	"dev-machine-proxy/internal/usage"
	"dev-machine-proxy/internal/web"
)

func main() {
	port := flag.Int("port", 9999, "Port to serve the dashboard on")
	projectsDir := flag.String("projects", "", "Directory containing project folders to scan for port references")
	refreshInterval := flag.Duration("refresh", 30*time.Second, "How often to refresh service discovery")
	flag.Parse()

	// Load configuration
	configMgr := config.NewManager()
	log.Printf("Config loaded from %s", configMgr.Get().Title)

	// Start system monitor (collect stats every 2 seconds, keep 60 data points = 2 minutes)
	sysMonitor := system.NewMonitor(60)
	sysMonitor.Start(2 * time.Second)
	log.Println("System monitor started")

	// Start AI usage monitor (collect every 5 minutes, keep 7 days of history)
	usageMonitor := usage.NewMonitor()
	usageMonitor.Start(5 * time.Minute)
	log.Println("AI usage monitor started")

	// Create the service discoverer
	disc := discovery.New(*projectsDir)

	// Initial discovery
	log.Println("Starting initial service discovery...")
	services, err := disc.Discover()
	if err != nil {
		log.Printf("Warning: initial discovery had errors: %v", err)
	}
	log.Printf("Discovered %d services", len(services))

	// Start background refresh
	go func() {
		ticker := time.NewTicker(*refreshInterval)
		for range ticker.C {
			log.Println("Refreshing service discovery...")
			if _, err := disc.Discover(); err != nil {
				log.Printf("Warning: discovery refresh had errors: %v", err)
			}
		}
	}()

	// Set up web server
	handler := web.NewHandler(disc, configMgr, sysMonitor, usageMonitor, *projectsDir)

	addr := fmt.Sprintf(":%d", *port)
	log.Printf("Starting dashboard on http://localhost%s", addr)
	if err := http.ListenAndServe(addr, handler); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
