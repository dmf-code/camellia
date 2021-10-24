package service

import (
	"camellia/internal/admin/domain/permission/service"
	"camellia/internal/admin/interfaces/assembler"
	resp "camellia/tool/response"
	"github.com/gin-gonic/gin"
)

type RoleApplicationService struct {
	response resp.Response
}

func (srv *RoleApplicationService) Create(c *gin.Context) {

	roleDTO, err := assembler.RoleCreateToDTO(c)

	if err != nil {
		resp.Error(c, resp.ParameterError, resp.CodeText(resp.ParameterError))
		return
	}

	roleDomainService := service.RoleDomainService{
		RoleDTO: roleDTO,
	}

	err = roleDomainService.Create()

	if err != nil {

		resp.Error(c, resp.BadRequest, resp.CodeText(resp.BadRequest))

		return
	}

	resp.Success(c, resp.Ok, resp.CodeText(resp.Ok))
}
