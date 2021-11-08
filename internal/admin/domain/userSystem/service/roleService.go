package service

import (
	"camellia/internal/admin/interfaces/assembler"
	"camellia/internal/admin/interfaces/dto"
	"camellia/tool/database"
	"fmt"
)

type RoleDOService struct {
	RoleDTO dto.RoleDTO
}

func (srv *RoleDOService) Create() error {
	db := database.GetInstance()
	rolePO := assembler.RoleDTOToPO(&srv.RoleDTO)
	res := db.Table("role").Create(&rolePO)

	if res.Error != nil {
		return res.Error
	}

	fmt.Println(res)

	return nil
}
