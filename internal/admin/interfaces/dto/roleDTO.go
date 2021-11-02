package dto


type RoleDTO struct {
	Pid string `json:"pid" binding:"required"`
	Name string `json:"name" binding:"required"`
	Memo string `json:"memo"`
	Sequence uint64 `json:"sequence"`
}

