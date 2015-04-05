package dockermonitor

import (
	"github.com/fsouza/go-dockerclient"
)

func GetClient() (*docker.Client, error) {
	endpoint := "unix:///var/run/docker.sock"
	return docker.NewClient(endpoint)
}

func GetImages() ([]docker.APIImages, error) {
	endpoint := "unix:///var/run/docker.sock"
	client, _ := docker.NewClient(endpoint)
	return client.ListImages(docker.ListImagesOptions{All: false})
}
