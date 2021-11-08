package service

import (
	"camellia/internal/admin/interfaces/dto"
	"camellia/tool/database"
)

type UserDOService struct {
	UserDTO dto.UserDTO
}

func (src *UserDOService) Login() error {

	return nil
}

func (src *UserDOService) Register() error {

	db := database.GetInstance()

	db.Table("")

	return nil
}
