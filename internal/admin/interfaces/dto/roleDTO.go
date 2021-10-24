package dto


type RoleDTO struct {
	Pid string `json:"pid"`
	Name string `json:"name"`
	Memo string `json:"memo"`
	Sequence uint64 `json:"sequence"`
}

