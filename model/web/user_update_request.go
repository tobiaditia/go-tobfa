package web

type UserUpdateRequest struct {
	Id        int
	Name      string `validate:"required,max=100,min=3" json:"name"`
	Handphone string `validate:"required,min=11,max=14" json:"handphone"`
}
