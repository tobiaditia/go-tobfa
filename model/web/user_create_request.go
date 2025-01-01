package web

type UserCreateRequest struct {
	Name      string `validate:"required,max=100,min=3" json:"name" example:"tobfa"`
	Email     string `validate:"required,email" json:"email" example:"me@tobfa.id"`
	Password  string `validate:"required,len=8" json:"password" example:"12345678"`
	Handphone string `validate:"required,min=11,max=14" json:"handphone" example:"08510101010"`
}
