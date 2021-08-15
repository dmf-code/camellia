package service

import (
	"camellia/internal/dockerfile/domain/php/entity"
	"camellia/internal/dockerfile/domain/php/service"
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/jhoonb/archivex"
	"io/ioutil"
	"os"
	"time"
)

type PHPService struct {
	PHPDomainService service.PHPDomainService
}

// 构建 dockerfile 文件
func (srv *PHPService) Build(dockerfile *entity.Dockerfile) {

	path, _ := os.Getwd()
	fmt.Println(path)
	srv.PHPDomainService.ReadDockerfile(path+"/template/php/Dockerfile")

	bg := context.Background()

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	// tar 文件创建
	tar := new(archivex.TarFile)
	tar.Create(path+"/template/archieve")
	tar.AddAll(path+"/template/php", false)
	tar.Close()
	dockerBuildContext, err := os.Open(path+"/template/archieve")
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

	if err != nil {
		fmt.Println("error")
		fmt.Printf("%s", err.Error())
		return
	}

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

