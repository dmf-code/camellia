package util

import (
	"github.com/docker/docker/client"
	"sync"
)

var dockerApiClient *client.Client
var once sync.Once

func NewDockerApi() *client.Client {

	once.Do(func() {
		var err error
		dockerApiClient, err = client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
		if err != nil {
			panic("初始化docker api sdk失败")
		}
	})

	return dockerApiClient
}
