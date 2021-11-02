package service

import "camellia/internal/admin/interfaces/dto"

type UserDomainService struct {
	UserDTO dto.UserDTO
}

func (src *UserDomainService) Login() error {

	return nil
}
