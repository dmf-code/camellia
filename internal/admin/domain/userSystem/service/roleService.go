package service

import (
	"camellia/internal/admin/interfaces/assembler"
	"camellia/internal/admin/interfaces/dto"
	"camellia/tool/database"
	"fmt"
)

type RoleService struct {
	RoleDTO dto.RoleDTO
}

func (srv *RoleService) Create() error {
	db := database.GetInstance()
	rolePO := assembler.RoleDTOToPO(&srv.RoleDTO)
	res := db.Create(&rolePO)

	if res.Error != nil {
		return res.Error
	}

	fmt.Println(res)

	return nil
}
