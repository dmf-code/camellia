package service

import (
	"camellia/internal/admin/interfaces/dto"
	"camellia/tool/database"
)

type UserDomainService struct {
	UserDTO dto.UserDTO
}

func (src *UserDomainService) Login() error {

	return nil
}

func (src *UserDomainService) Register() error {

	db := database.GetInstance()

	db.Table("")

	return nil
}
