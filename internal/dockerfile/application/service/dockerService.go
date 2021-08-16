package service

import (
	"camellia/internal/dockerfile/domain/php/entity"
	"camellia/internal/dockerfile/domain/php/service"
	"camellia/tool/exception"
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/jhoonb/archivex"
	"io/ioutil"
	"os"
	"time"
)

type DockerService struct {
	DockerDomainService service.DockerDomainService
}

// 构建 dockerfile 文件
func (srv *DockerService) Build(dockerfile *entity.Dockerfile) {

	path, _ := os.Getwd()

	srv.DockerDomainService.ReadDockerfile(path+"/template/php/Dockerfile")

	bg := context.Background()

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	exception.GetIns().Throw(err)

	// tar 文件创建
	tar := new(archivex.TarFile)

	err = tar.Create(path+"/template/archieve")
	exception.GetIns().Throw(err)

	err = tar.AddAll(path+"/template/php", false)
	exception.GetIns().Throw(err)

	err = tar.Close()
	exception.GetIns().Throw(err)

	dockerBuildContext, err := os.Open(path+"/template/archieve")
	exception.GetIns().Throw(err)

	defer dockerBuildContext.Close()

	options := types.ImageBuildOptions{
		SuppressOutput: true,
		Remove:         true,
		ForceRemove:    true,
		PullParent:     true,
		Tags:           []string{"bootstrap"},
		BuildArgs: 		dockerfile.Args,
		Dockerfile: 	"Dockerfile",
	}

	buildResponse, err := cli.ImageBuild(bg, dockerBuildContext, options)

	exception.GetIns().Throw(err)

	defer buildResponse.Body.Close()

	time.Sleep(5000 * time.Millisecond)
	fmt.Printf("********* %s **********\n", buildResponse.OSType)
	response, err := ioutil.ReadAll(buildResponse.Body)
	if err != nil {
		fmt.Printf("%s", err.Error())
	}
	fmt.Println(string(response))

	return
}

func (srv *DockerService) Run() {

}

