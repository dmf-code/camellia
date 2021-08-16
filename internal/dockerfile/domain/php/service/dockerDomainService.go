package service

import (
	"os"
)

type DockerDomainService struct {
	ArgStr     string
	Dockerfile string
}

func (srv *DockerDomainService) ReadDockerfile(path string) {
	file, _ := os.Open(path)
	defer file.Close()
}
