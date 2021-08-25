package controller

import (
	"camellia/internal/docker/application/service"
	"camellia/internal/docker/interfaces/assembler"
	resp "camellia/tool/response"
	"github.com/gin-gonic/gin"
)

type ContainerController struct {
}

func (ctrl *ContainerController) List(ctx *gin.Context) {
	containerService := service.ContainerService{}
	list := containerService.List()

	resp.Success(ctx, "ok", list)
}

func (ctrl *ContainerController) NewContainer(ctx *gin.Context) {
	containerService := service.ContainerService{}
	containerService.New()
}

func (ctrl *ContainerController) Stop(ctx *gin.Context) {
	dto := assembler.ToContainerDTO(ctx)
	containerService := service.ContainerService{}

	containerService.Stop(dto)
}

func (ctrl *ContainerController) Start(ctx *gin.Context) {

	dto := assembler.ToContainerDTO(ctx)
	containerService := service.ContainerService{}

	containerService.Start(dto)
}

func (ctrl *ContainerController) Remove(ctx *gin.Context) {

	dto := assembler.ToContainerDTO(ctx)

	containerService := service.ContainerService{}

	containerService.Remove(dto)
}
