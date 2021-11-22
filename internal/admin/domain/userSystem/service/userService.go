package service

import (
	"camellia/internal/admin/interfaces/dto"
	"camellia/tool/database"
)

type UserService struct {
	UserDTO dto.UserDTO
}

func (src *UserService) Login() error {

	return nil
}

func (src *UserService) Register() error {

	db := database.GetInstance()

	db.Table("")

	return nil
}
