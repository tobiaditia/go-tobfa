package web

type BusinessTransactionUpdateRequest struct {
	Id                        int
	BusinessId                int    `validate:"required" json:"businessId"`
	BusinessTransactionTypeId int    `validate:"required" json:"businessTransactionTypeId"`
	BusinessTransactionItemId int    `validate:"required" json:"businessTransactionItemId"`
	Total                     int    `validate:"required,numeric,min=0" json:"total"`
	Quantity                  int    `validate:"required,numeric,min=0" json:"quantity"`
	Date                      string `validate:"required" json:"date"`
	Description               string `validate:"required" json:"description"`
}
