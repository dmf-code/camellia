package service

import (
	"camellia/internal/admin/domain/userSystem/service"
	"camellia/internal/admin/interfaces/assembler"
	resp "camellia/tool/response"
	"github.com/gin-gonic/gin"
)

type RoleApplicationService struct {
}

func (srv *RoleApplicationService) Create(c *gin.Context) {

	roleDTO, err := assembler.RoleCreateToDTO(c)

	if err != nil {
		resp.Error(c, resp.ParameterError, err.Error())
		return
	}

	roleDomainService := service.NewRoleService(roleDTO)

	err = roleDomainService.Create()

	if err != nil {
		resp.Error(c, resp.BadRequest, resp.CodeText(resp.BadRequest))
		return
	}

	resp.Success(c, resp.Ok, resp.CodeText(resp.Ok))
}
