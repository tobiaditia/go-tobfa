package web

type BusinessTransactionCreateRequest struct {
	BusinessId                int    `validate:"required" json:"businessId"`
	BusinessTransactionTypeId int    `validate:"required" json:"businessTransactionTypeId"`
	BusinessTransactionItemId int    `validate:"required" json:"businessTransactionItemId"`
	Total                     int    `validate:"required" json:"total"`
	Multiplier                int    `validate:"required" json:"multiplier"`
	Date                      string `validate:"required,datetime" json:"date"`
	Description               string `validate:"required" json:"description"`
}
