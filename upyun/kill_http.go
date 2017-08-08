package main

import (
	"fmt"

	"github.com/fsouza/go-dockerclient"
)

func main() {
	endpoint := "unix:///var/run/docker.sock"
	client, err := docker.NewClient(endpoint)
	if err != nil {
		panic(err)
	}
	// START OMIT
	containers, err := client.ListContainers(docker.ListContainersOptions{All: false})
	if err != nil {
		panic(err)
	}
	for _, cont := range containers {
		if cont.Image == "nginx:latest" {
			fmt.Println("ID: ", cont.ID)
			fmt.Println("Image: ", cont.Image)
			err := client.StopContainer(cont.ID, 30)
			if err != nil {
				panic(err)
			}
		}
	}
	// END OMIT
}
