package controller

import (
	"camellia/internal/docker/application/service"
	"camellia/internal/docker/domain/php/entity"
	"github.com/gin-gonic/gin"
)

type ImageController struct {
}

func (ctrl *ImageController) Build(ctx *gin.Context) {
	srv := service.ImageService{}

	srv.Build(entity.New())
}
