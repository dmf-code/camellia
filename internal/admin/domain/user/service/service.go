package service

import "camellia/internal/admin/interfaces/dto"

func NewUserService(dto dto.UserDTO) (user UserService, err error) {

	return UserService{}, nil
}
