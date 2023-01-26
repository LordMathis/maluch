package agents

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"reflect"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
)

type DockerAgent struct {
	BaseAgent
	workdir string
	client  *client.Client
	// action  string
	// image   string
	// args    []string
}

const (
	Build Action = "build"
	Run          = "run"
)

func getClient() (*client.Client, error) {
	cli, err := client.NewClientWithOpts()
	if err != nil {
		return nil, err
	}

	return cli, nil

}

func NewDockerAgent(workdir string) *DockerAgent {

	client, err := getClient()
	if err != nil {
		log.Fatal(err)
	}

	action := map[Action]interface{}{
		Build: DockerAgent.BuildImage,
		Run:   DockerAgent.RunContainer,
	}

	return &DockerAgent{
		BaseAgent: BaseAgent{
			name:   "docker",
			action: action,
		},
		workdir: workdir,
		client:  client,
	}
}

func (d DockerAgent) RunAction(a Action) error {
	f := reflect.ValueOf(d.action[a])

	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	var res []reflect.Value
	res = f.Call(in)
	result = res[0].Interface()
	return
}

func (d DockerAgent) BuildImage(ctx context.Context, dockerfile string, image []string) error {

	tar, err := archive.TarWithOptions(d.workdir, &archive.TarOptions{})
	if err != nil {
		return err
	}

	opts := types.ImageBuildOptions{
		Dockerfile: dockerfile,
		Tags:       image,
		Remove:     true,
	}

	res, err := d.client.ImageBuild(ctx, tar, opts)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	scanner := bufio.NewScanner(res.Body)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	return nil
}

func (d DockerAgent) RunContainer(ctx context.Context, image string, command []string) error {

	resp, err := d.client.ContainerCreate(ctx, &container.Config{
		Image: image,
		Cmd:   command,
		Tty:   false,
	}, nil, nil, nil, "")
	if err != nil {
		return err
	}

	if err := d.client.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		return err
	}

	statusCh, errCh := d.client.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			return err
		}
	case <-statusCh:
	}

	_, err = d.client.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		return err
	}

	return nil
}
