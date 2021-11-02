package service

import (
	"camellia/internal/admin/domain/userSystem/service"
	"camellia/internal/admin/interfaces/assembler"
	resp "camellia/tool/response"
	"github.com/gin-gonic/gin"
)

type UserApplicationService struct {
}

func (src *UserApplicationService) Login(c *gin.Context) {
	userDTO, err := assembler.UserCreateToDTO(c)

	if err != nil {
		resp.Error(c, resp.ParameterError, err.Error())
		return
	}

	userDomainService := service.UserDomainService{
		UserDTO: userDTO,
	}

	err = userDomainService.Login()

	if err != nil {

		resp.Error(c, resp.BadRequest, resp.CodeText(resp.BadRequest))

		return
	}

	resp.Success(c, resp.Ok, resp.CodeText(resp.Ok))
}

