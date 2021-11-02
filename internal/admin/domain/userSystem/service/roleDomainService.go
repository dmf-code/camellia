package service

import (
	"camellia/internal/admin/interfaces/assembler"
	"camellia/internal/admin/interfaces/dto"
	"camellia/tool/database"
	"fmt"
)

type RoleDomainService struct {
	RoleDTO dto.RoleDTO
}

func (srv *RoleDomainService) Create() error {
	db := database.GetInstance()
	rolePO := assembler.RoleDTOToPO(&srv.RoleDTO)
	res := db.Table("role").Create(&rolePO)

	if res.Error != nil {
		return res.Error
	}

	fmt.Println(res)

	return nil
}
