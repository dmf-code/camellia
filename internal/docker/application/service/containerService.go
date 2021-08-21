package service

import (
	"camellia/internal/docker/infrastructure/util"
	"camellia/tool/exception"
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/network"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
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

	hostConfig := container.HostConfig{
		Binds:           nil,
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

	platform := v1.Platform{
		Architecture: "",
		OS:           "",
		OSVersion:    "",
		OSFeatures:   nil,
		Variant:      "",
	}

	containerName := "test"
	body, err := cli.ContainerCreate(
		bg,
		&config,
		&hostConfig,
		&netWorkingConfig,
		&platform,
		containerName)

	exception.GetIns().Throw(err)

	fmt.Println(body)

	err = cli.ContainerStart(bg, body.ID, types.ContainerStartOptions{})

	exception.GetIns().Throw(err)
}

