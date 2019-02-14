package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"

	"moul.io/dockerself"
)

func main() {
	dockerize := flag.Bool("dockerize", false, "dockerize")
	image := flag.String("image", "ubuntu", "docker image")
	flag.Parse()
	if *dockerize {
		err := dockerself.Dockerize(*image)
		if err != nil {
			panic(err)
		}
		fmt.Println("container exited.")
		return
	}

	fmt.Println("# dockerself-example")
	fmt.Printf("env: %+v\n", os.Environ())
	fmt.Println("starting interactive `/bin/sh` session...")
	cmd := exec.Command("/bin/sh")
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}
