package service

import (
	"camellia/internal/docker/infrastructure/util"
	"camellia/tool/exception"
	"context"
	"github.com/docker/docker/api/types"
)

func RemoveContainer(containerId string, options types.ContainerRemoveOptions) {
	cli := util.NewDockerApi()
	bg := context.Background()
	err := cli.ContainerRemove(bg, containerId, options)

	exception.GetIns().Throw(err)
}
