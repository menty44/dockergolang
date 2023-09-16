package main

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"golang.org/x/net/context"

	_ "github.com/docker/docker/api/types/container"
	//"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func main() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	images, err := cli.ImageList(ctx, types.ImageListOptions{})
	if err != nil {
		panic(err)
	}

	for _, image := range images {
		fmt.Println(image.ID)
	}
	//fmt.Printf("kakakkaka")
	//ctx := context.Background()
	//cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	//if err != nil {
	//	panic(err)
	//}
	//defer cli.Close()
	//
	//options := types.ContainerLogsOptions{ShowStdout: true}
	//// Replace this ID with a container that really exists
	//out, err := cli.ContainerLogs(ctx, "1afd5705259e", options)
	//if err != nil {
	//	panic(err)
	//}
	//
	//io.Copy(os.Stdout, out)
}
