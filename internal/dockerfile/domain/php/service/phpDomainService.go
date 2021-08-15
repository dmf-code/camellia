package service

import (
	"os"
)

type PHPDomainService struct {
	ArgStr     string
	Dockerfile string
}

func (srv *PHPDomainService) ReadDockerfile(path string) {
	file, _ := os.Open(path)
	defer file.Close()
}
