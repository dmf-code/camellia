package service

import "camellia/internal/admin/interfaces/dto"

func NewUserService(dto dto.UserDTO) *UserService {
	return &UserService{UserDTO: dto}
}

func NewRoleService(dto dto.RoleDTO) *RoleService {
	return &RoleService{RoleDTO:dto}
}
