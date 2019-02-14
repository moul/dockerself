package main

import (
	"flag"
	"fmt"
	"os"

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
		return
	}
	fmt.Println(os.Environ())
}
