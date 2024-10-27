package web

type UserCreateRequest struct {
	Name      string `validate:"required,max=100,min=3" json:"name"`
	Email     string `validate:"required,email" json:"email"`
	Password  string `validate:"required,len=8" json:"password"`
	Handphone string `validate:"required,min=11,max=14" json:"handphone"`
}
