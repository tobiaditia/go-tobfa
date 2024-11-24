package web

type BusinessTransactionUpdateRequest struct {
	Id                        int
	BusinessId                int    `validate:"required,max=100,min=3" json:"businessId"`
	BusinessTransactionTypeId int    `validate:"required" json:"businessTransactionTypeId"`
	BusinessTransactionItemId int    `validate:"required" json:"businessTransactionItemId"`
	Total                     int    `validate:"required" json:"total"`
	Multiplier                int    `validate:"required" json:"multiplier"`
	Date                      string `validate:"required" json:"date"`
	Description               string `validate:"required" json:"description"`
}
