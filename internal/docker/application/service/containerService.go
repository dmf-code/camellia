package service

import (
	"camellia/internal/docker/infrastructure/util"
	"camellia/internal/docker/interfaces/dto"
	"camellia/tool/exception"
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/network"
)

type ContainerService struct {

}

func (srv *ContainerService) List() []types.Container {
	cli := util.NewDockerApi()
	bg := context.Background()

	options := types.ContainerListOptions{
		Quiet:   false,
		Size:    false,
		All:     false,
		Latest:  false,
		Since:   "",
		Before:  "",
		Limit:   0,
		Filters: filters.Args{},
	}

	list, err := cli.ContainerList(bg, options)

	exception.GetIns().Throw(err)

	return list
}

func (srv *ContainerService) New() {
	cli := util.NewDockerApi()
	bg := context.Background()
	imageName := "bootstrap"

	config := container.Config{
		Hostname:        "",
		Domainname:      "",
		User:            "",
		AttachStdin:     false,
		AttachStdout:    false,
		AttachStderr:    false,
		ExposedPorts:    nil,
		Tty:             false,
		OpenStdin:       false,
		StdinOnce:       false,
		Env:             nil,
		Cmd:             nil,
		Healthcheck:     nil,
		ArgsEscaped:     false,
		Image:           imageName,
		Volumes:         nil,
		WorkingDir:      "",
		Entrypoint:      nil,
		NetworkDisabled: false,
		MacAddress:      "",
		OnBuild:         nil,
		Labels:          nil,
		StopSignal:      "",
		StopTimeout:     nil,
		Shell:           nil,
	}

	var binds []string

	binds = []string{"/work/apps/php/learnku:/var/www/html"}

	hostConfig := container.HostConfig{
		Binds:           binds,
		ContainerIDFile: "",
		LogConfig:       container.LogConfig{},
		NetworkMode:     "",
		PortBindings:    nil,
		RestartPolicy:   container.RestartPolicy{},
		AutoRemove:      false,
		VolumeDriver:    "",
		VolumesFrom:     nil,
		CapAdd:          nil,
		CapDrop:         nil,
		CgroupnsMode:    "",
		DNS:             nil,
		DNSOptions:      nil,
		DNSSearch:       nil,
		ExtraHosts:      nil,
		GroupAdd:        nil,
		IpcMode:         "",
		Cgroup:          "",
		Links:           nil,
		OomScoreAdj:     0,
		PidMode:         "",
		Privileged:      false,
		PublishAllPorts: false,
		ReadonlyRootfs:  false,
		SecurityOpt:     nil,
		StorageOpt:      nil,
		Tmpfs:           nil,
		UTSMode:         "",
		UsernsMode:      "",
		ShmSize:         0,
		Sysctls:         nil,
		Runtime:         "",
		ConsoleSize:     [2]uint{},
		Isolation:       "",
		Resources:       container.Resources{},
		Mounts:          nil,
		MaskedPaths:     nil,
		ReadonlyPaths:   nil,
		Init:            nil,
	}

	netWorkingConfig := network.NetworkingConfig{}

	//platform := v1.Platform{
	//	Architecture: "",
	//	OS:           "",
	//	OSVersion:    "",
	//	OSFeatures:   nil,
	//	Variant:      "",
	//}

	containerName := "test"

	body, err := cli.ContainerCreate(
		bg,
		&config,
		&hostConfig,
		&netWorkingConfig,
		nil,
		containerName)

	fmt.Println(body)

	exception.GetIns().Throw(err)

	err = cli.ContainerStart(bg, body.ID, types.ContainerStartOptions{})

	exception.GetIns().Throw(err)
}

func (srv *ContainerService) Stop(dto *dto.ContainerDTO) {
	cli := util.NewDockerApi()
	bg := context.Background()
	err := cli.ContainerStop(
		bg,
		dto.Id,
		nil)

	exception.GetIns().Throw(err)

}

func (srv *ContainerService) Start(dto *dto.ContainerDTO) {

	cli := util.NewDockerApi()
	bg := context.Background()
	err := cli.ContainerStart(
		bg,
		dto.Id,
		types.ContainerStartOptions{
			CheckpointID:  "",
			CheckpointDir: "",
		})

	exception.GetIns().Throw(err)
}

func (srv *ContainerService) Remove(dto *dto.ContainerDTO) {
	cli := util.NewDockerApi()
	bg := context.Background()

	err := cli.ContainerRemove(
		bg,
		dto.Id,
		types.ContainerRemoveOptions{
			Force:         true,
		})

	exception.GetIns().Throw(err)
}

