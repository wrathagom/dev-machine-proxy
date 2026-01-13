package discovery

import (
	"context"
	"fmt"
	"strings"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

// DockerContainer represents a running Docker container with exposed ports
type DockerContainer struct {
	ID     string
	Name   string
	Image  string
	Ports  []ContainerPort
	Labels map[string]string
}

// ContainerPort maps a container port to a host port
type ContainerPort struct {
	ContainerPort int
	HostPort      int
	Protocol      string
}

// GetDockerContainers lists all running containers and their port mappings
func GetDockerContainers() ([]DockerContainer, error) {
	ctx := context.Background()

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, fmt.Errorf("creating docker client: %w", err)
	}
	defer cli.Close()

	containers, err := cli.ContainerList(ctx, container.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("listing containers: %w", err)
	}

	var result []DockerContainer
	for _, c := range containers {
		dc := DockerContainer{
			ID:     c.ID[:12],
			Name:   strings.TrimPrefix(c.Names[0], "/"),
			Image:  c.Image,
			Labels: c.Labels,
		}

		for _, p := range c.Ports {
			if p.PublicPort > 0 {
				dc.Ports = append(dc.Ports, ContainerPort{
					ContainerPort: int(p.PrivatePort),
					HostPort:      int(p.PublicPort),
					Protocol:      p.Type,
				})
			}
		}

		if len(dc.Ports) > 0 {
			result = append(result, dc)
		}
	}

	return result, nil
}

// GetContainerByPort finds a container that has a specific host port mapped
func GetContainerByPort(containers []DockerContainer, port int) *DockerContainer {
	for i, c := range containers {
		for _, p := range c.Ports {
			if p.HostPort == port {
				return &containers[i]
			}
		}
	}
	return nil
}

// GuessServiceFromContainer tries to identify what service a container is running
func GuessServiceFromContainer(c *DockerContainer) string {
	// Check for common labels
	if name, ok := c.Labels["com.docker.compose.service"]; ok {
		return name
	}
	if name, ok := c.Labels["org.opencontainers.image.title"]; ok {
		return name
	}

	// Try to extract from image name
	image := c.Image

	// Remove registry prefix
	if idx := strings.LastIndex(image, "/"); idx != -1 {
		image = image[idx+1:]
	}

	// Remove tag
	if idx := strings.Index(image, ":"); idx != -1 {
		image = image[:idx]
	}

	return image
}
