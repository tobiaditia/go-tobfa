package web

type BusinessTransactionCreateRequest struct {
	BusinessId                int    `validate:"required" json:"businessId"`
	BusinessTransactionTypeId int    `validate:"required" json:"businessTransactionTypeId"`
	BusinessTransactionItemId int    `validate:"required" json:"businessTransactionItemId"`
	Total                     int    `validate:"required,numeric,min=0" json:"total"`
	Quantity                  int    `validate:"required,numeric,min=0" json:"quantity"`
	Date                      string `validate:"required,datetime" json:"date"`
	Description               string `validate:"required" json:"description"`
}
