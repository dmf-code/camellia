package dto

type UserDTO struct {
	Username string `json:"username"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Password string `json:"password"`
	RepeatPassword string `json:"repeat_password"`
	Code string `json:"code"`
}
