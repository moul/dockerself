package dockerself

import (
	"archive/tar"
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/strslice"
	"github.com/docker/docker/client"
)

func Dockerize(image string) error {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return err
	}

	// create a container
	containerConfig := &container.Config{
		Image:        image,
		Tty:          true,
		OpenStdin:    true,
		AttachStdout: true,
		AttachStderr: true,
		Entrypoint:   strslice.StrSlice{"/dockerself"},
		Cmd:          append(os.Args),
	}
	hostConfig := &container.HostConfig{
		//AutoRemove: true,
	}
	cont, err := cli.ContainerCreate(ctx, containerConfig, hostConfig, nil, "")
	if err != nil {
		return err
	}

	// create a tarball containing "self" binary
	var buf bytes.Buffer
	binaryPath, err := os.Executable()
	if err != nil {
		return err
	}
	binary, err := ioutil.ReadFile(binaryPath)
	if err != nil {
		return err
	}
	tw := tar.NewWriter(&buf)
	if err := tw.WriteHeader(&tar.Header{
		Name: "dockerself",
		Mode: 0755,
		Size: int64(len(binary)),
	}); err != nil {
		return err
	}
	if _, err := tw.Write(binary); err != nil {
		return err
	}
	if err := tw.Close(); err != nil {
		return err
	}

	// copy and extract tarball into the container
	if err := cli.CopyToContainer(
		ctx,
		cont.ID,
		"/",
		&buf,
		types.CopyToContainerOptions{},
	); err != nil {
		return err
	}

	// start the container
	ctxCancel, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := cli.ContainerStart(ctxCancel, cont.ID, types.ContainerStartOptions{}); err != nil {
		return err
	}

	// attach to container
	resp, err := cli.ContainerAttach(ctx, cont.ID, types.ContainerAttachOptions{
		Stream: true,
		Stdin:  true,
		Stdout: true,
		Stderr: true,
		Logs:   false,
	})

	fmt.Println(resp)

	_, errC := cli.ContainerWait(ctx, cont.ID, "")
	if err := <-errC; err != nil {
		return err
	}

	// exit to avoid running the same code again
	return nil
}
