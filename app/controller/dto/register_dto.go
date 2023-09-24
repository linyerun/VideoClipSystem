package dto

// swagger:model RegisterDto
type RegisterDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	AuthCode string `json:"auth_code"`
}
