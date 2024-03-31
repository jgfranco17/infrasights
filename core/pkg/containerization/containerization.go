package containerization

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func CountDockerImages() (int, error) {
	// Create a new Docker client
	dockerClient, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return 0, err
	}

	// Fetch the list of Docker images
	images, err := dockerClient.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		return 0, err
	}

	return len(images), nil
}
