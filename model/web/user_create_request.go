package web

type UserCreateRequest struct {
	Name      string `validate:"required,alphanum,max=100,min=3" json:"name" example:"tobfa"`
	Email     string `validate:"required,email" json:"email" example:"me@tobfa.id"`
	Handphone string `validate:"required,min=11,max=14" json:"handphone" example:"08510101010"`
	Password  string `validate:"required,len=8" json:"password" example:"12345678"`
	// ConfirmPassword string `validate:"required,len=8,eqfield=Password" json:"password" example:"12345678"`
}
