package web

type LoginRequest struct {
	Email    string `validate:"required,email" json:"email" example:"me@tobfa.id"`
	Password string `validate:"required,len=8" json:"password" example:"12345678"`
}
