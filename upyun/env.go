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
			dk, err := client.InspectContainer(cont.ID)
			if err != nil {
				panic(err)
			}
			fmt.Println("ID: ", cont.ID)
			fmt.Println("Image: ", cont.Image)
			fmt.Println("Env: ", dk.Config.Env)
		}
	}
	// END OMIT
}
