package main

import (
	"context"
	"strings"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/pieterclaerhout/go-formatter"
	"github.com/pieterclaerhout/go-log"
)

func main() {

	log.PrintColors = true
	log.PrintTimestamp = false

	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{
		All: true,
	})
	log.CheckError(err)

	for _, container := range containers {

		if container.State == "running" {
			continue
		}

		containerID := container.ID[:10]

		log.Info("> Removing container:", containerID, container.Image)

	}

	images, err := cli.ImageList(context.Background(), types.ImageListOptions{
		All: false,
	})

	for _, image := range images {
		for _, tag := range image.RepoTags {

			tagParts := strings.SplitN(tag, ":", 2)

			log.Infof(
				"%-40s | %-30s | %s | %s | %10s",
				tagParts[0],
				tagParts[1],
				image.ID[7:19],
				time.Unix(image.Created, 0),
				formatter.FileSize(image.Size),
			)
		}
	}

	// log.InfoDump(images, "images")

}
