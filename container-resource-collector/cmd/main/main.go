package main

import (
	"context"
	"log"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/htchan/ContainerResourceCollector/internal/container"
)

const INTERVAL = time.Second * 5

var containerMap = make(map[string]*container.Container)

func updateContainers(ctx context.Context, cli *client.Client) {
	log.Println("going to list container")
	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		log.Printf("list container error: %v\n", err)
		return
	}
	log.Printf("container length: %d\n", len(containers))

	containerIDs := make([]string, 0)
	for containerID := range containerMap {
		containerIDs = append(containerIDs, containerID)
	}

	for _, c := range containers {
		// TODO: find not existing container
		for i, id := range containerIDs {
			if id == c.ID {
				containerIDs = append(containerIDs[:i], containerIDs[i+1:]...)
				break
			}
		}

		// add new container to map
		_, exist := containerMap[c.ID]
		if !exist {
			containerMap[c.ID] = &container.Container{
				ContainerID: c.ID,
				Project:     c.Labels["com.docker.compose.project"],
				Service:     c.Labels["com.docker.compose.service"],
			}

			go func(id string) {
				for {
					if _, exist := containerMap[id]; !exist {
						break
					}

					containerMap[id].ReadResource(ctx, cli)
					time.Sleep(INTERVAL)
					log.Println(containerMap[id].Resources.CPU)
					log.Println(containerMap[id].Resources.Memory)
					log.Println(containerMap[id].Resources.Network)
					log.Println(containerMap[id].Resources.UpdatedAt)
				}
			}(c.ID)
		}
	}

	for _, id := range containerIDs {
		delete(containerMap, id)
	}
}

func main() {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	background := context.Background()

	// go func() {
	for {
		updateContainers(background, cli)
		time.Sleep(INTERVAL)
	}
	// }()
}
