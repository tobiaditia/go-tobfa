package web

type UserUpdatePasswordRequest struct {
	Password string `validate:"required,len=8" json:"password" example:"12345678"`
}
