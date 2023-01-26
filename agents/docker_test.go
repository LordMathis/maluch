package agents

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/docker/docker/api/types"
	"gotest.tools/v3/assert"
)

func (d *DockerAgent) Cleanup() {
	os.RemoveAll(d.workdir)
}

func TestBuildImage(t *testing.T) {
	workdir, _ := filepath.Abs("./testdata")

	docker_agent := NewDockerAgent(workdir)
	tags := [1]string{"test_image"}

	ctx := context.Background()

	err := docker_agent.BuildImage(ctx, "Dockerfile", tags[:])

	assert.NilError(t, err)
}

func TestRunContainer(t *testing.T) {
	workdir, _ := filepath.Abs("./testdata")
	docker_agent := NewDockerAgent(workdir)
	ctx := context.Background()

	image := "busybox"

	_, err := docker_agent.client.ImagePull(ctx, image, types.ImagePullOptions{})
	assert.NilError(t, err)

	command := [2]string{"echo", "Hello World!"}
	err = docker_agent.RunContainer(ctx, image, command[:])

	assert.NilError(t, err)
}
