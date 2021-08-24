package assembler

import (
	"camellia/internal/docker/interfaces/dto"
	"github.com/gin-gonic/gin"
)

func ToContainerDTO(ctx *gin.Context) *dto.ContainerDTO {

	containerDTO := dto.ContainerDTO{Id: ctx.Query("id")}

	return &containerDTO
}
