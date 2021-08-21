package controller

import (
	"camellia/internal/docker/application/service"
	resp "camellia/tool/response"
	"github.com/gin-gonic/gin"
)

type ContainerController struct {
}

func (ctr *ContainerController) List(ctx *gin.Context) {
	containerService := service.ContainerService{}
	list := containerService.List()

	resp.Success(ctx, "ok", list)
}

func (ctr *ContainerController) NewContainer(ctx *gin.Context) {
	containerService := service.ContainerService{}
	containerService.New()
}
