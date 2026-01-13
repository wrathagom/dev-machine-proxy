package discovery

import (
	"fmt"
	"log"
	"sort"
	"sync"
)

// Discoverer orchestrates service discovery from multiple sources
type Discoverer struct {
	projectsDir string
	services    []Service
	mu          sync.RWMutex
}

// New creates a new Discoverer
func New(projectsDir string) *Discoverer {
	return &Discoverer{
		projectsDir: projectsDir,
	}
}

// Discover runs all discovery mechanisms and returns discovered services
func (d *Discoverer) Discover() ([]Service, error) {
	var allErrors []error

	// Step 1: Get all listening ports
	log.Println("  Scanning listening ports...")
	listeningPorts, err := GetListeningPorts()
	if err != nil {
		allErrors = append(allErrors, fmt.Errorf("port scan: %w", err))
	}
	log.Printf("  Found %d listening ports", len(listeningPorts))

	// Step 2: Get Docker containers
	log.Println("  Querying Docker...")
	containers, err := GetDockerContainers()
	if err != nil {
		log.Printf("  Docker not available: %v", err)
		// Not fatal - Docker might not be running
	} else {
		log.Printf("  Found %d containers with exposed ports", len(containers))
	}

	// Step 3: Scan projects directory
	var projectMatches []ProjectMatch
	if d.projectsDir != "" {
		log.Printf("  Scanning projects in %s...", d.projectsDir)
		ports := make([]int, len(listeningPorts))
		for i, p := range listeningPorts {
			ports[i] = p.Port
		}
		projectMatches, err = ScanProjectsForPorts(d.projectsDir, ports)
		if err != nil {
			log.Printf("  Project scan error: %v", err)
		} else {
			log.Printf("  Found %d project/port matches", len(projectMatches))
		}
	}

	// Step 4: Build service list
	services := make([]Service, 0, len(listeningPorts))

	for _, lp := range listeningPorts {
		svc := Service{
			Port:     lp.Port,
			Protocol: "tcp",
			Process:  lp.Process,
			Source:   "port-scan",
		}

		var dockerName string

		// Check if this port belongs to a Docker container (highest priority for naming)
		if container := GetContainerByPort(containers, lp.Port); container != nil {
			svc.Source = "docker"
			svc.Container = container.Name
			svc.Image = container.Image
			dockerName = GuessServiceFromContainer(container)
			svc.Tags = append(svc.Tags, "docker")

			// For Docker containers, check if there's a compose project directory
			if projectDir, ok := container.Labels["com.docker.compose.project.working_dir"]; ok {
				svc.ProjectPath = projectDir
				svc.Tags = append(svc.Tags, "project")
			}
		}

		// Only check project folder matches for non-Docker services
		if svc.Source != "docker" {
			if match := FindProjectForPort(projectMatches, lp.Port); match != nil {
				svc.ProjectPath = match.ProjectPath
				if svc.Name == "" {
					svc.Name = match.ProjectName
				}
				svc.Description = fmt.Sprintf("Found in %s: %s", match.File, truncate(match.Context, 60))
				svc.Tags = append(svc.Tags, "project")
			}
		}

		// Apply known port names (only if we don't have a better name)
		if knownName, ok := KnownPorts[lp.Port]; ok {
			if svc.Name == "" && dockerName == "" {
				svc.Name = knownName
			}
			svc.Tags = append(svc.Tags, "known-port")
		}

		// HTTP probe for likely HTTP ports or unknown services
		if HTTPPorts[lp.Port] || svc.Name == "" {
			probe := ProbeHTTP(lp.Port)
			if probe.IsHTTP {
				svc.IsHTTP = true
				scheme := "http"
				if probe.IsHTTPS {
					scheme = "https"
				}
				svc.URL = fmt.Sprintf("%s://localhost:%d", scheme, lp.Port)

				// Only use probe name if we don't have a Docker-derived name
				// (Docker image/container name is more accurate than HTTP server header)
				if dockerName == "" {
					if probeName := GuessServiceFromProbe(probe, lp.Port); probeName != "" {
						svc.Name = probeName
					}
				}

				if probe.Server != "" && svc.Description == "" {
					svc.Description = fmt.Sprintf("Server: %s", probe.Server)
				}

				svc.Tags = append(svc.Tags, "http")
			}
		}

		// Apply Docker name (highest priority - do this late so it wins)
		if dockerName != "" {
			svc.Name = dockerName
		}

		// If we still don't have a name, use the process name
		if svc.Name == "" && svc.Process != "" {
			svc.Name = svc.Process
		}

		// Last resort: just use the port
		if svc.Name == "" {
			svc.Name = fmt.Sprintf("Port %d", lp.Port)
		}

		services = append(services, svc)
	}

	// Sort by port number
	sort.Slice(services, func(i, j int) bool {
		return services[i].Port < services[j].Port
	})

	// Store results
	d.mu.Lock()
	d.services = services
	d.mu.Unlock()

	if len(allErrors) > 0 {
		return services, fmt.Errorf("discovery completed with errors: %v", allErrors)
	}

	return services, nil
}

// GetServices returns the last discovered services
func (d *Discoverer) GetServices() []Service {
	d.mu.RLock()
	defer d.mu.RUnlock()

	// Return a copy to avoid races
	result := make([]Service, len(d.services))
	copy(result, d.services)
	return result
}

func truncate(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-3] + "..."
}
