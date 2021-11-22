package assembler

import (
	"camellia/internal/admin/domain/userSystem/repository/po"
	"camellia/internal/admin/interfaces/dto"
	"fmt"
	"github.com/gin-gonic/gin"
)

func RoleCreateToDTO(c *gin.Context) (dto.RoleDTO, error) {
	var obj dto.RoleDTO
	if err := c.ShouldBind(&obj); err != nil {
		fmt.Println(err)
		return dto.RoleDTO{}, err
	}

	return obj, nil
}

func RoleDTOToPO(roleDTO *dto.RoleDTO) *po.RolePO {
	return &po.RolePO{
		Pid:                 roleDTO.Pid,
		Name:                roleDTO.Name,
		Memo:                roleDTO.Memo,
		Sequence:            0,
	}
}
