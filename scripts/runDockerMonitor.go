package main

import (
	"fmt"
	"github.com/longan/litchi/dockermonitor"
)

func main() {
	fmt.Println("Start docker monitor")

	client, _ := dockermonitor.GetClient()
	if client == nil {
		fmt.Println("Fail to get client")
	}

	images, _ := dockermonitor.GetImages()
	for _, img := range images {
		fmt.Println("ID: ", img.ID)
		fmt.Println("RepoTags: ", img.RepoTags)
		fmt.Println("Created: ", img.Created)
		fmt.Println("Size: ", img.Size)
		fmt.Println("VirtualSize: ", img.VirtualSize)
		fmt.Println("ParentId: ", img.ParentID)
	}
}
