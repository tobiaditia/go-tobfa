package web

import "time"

type BusinessTransactionResponse struct {
	Id                        int                             `json:"id"`
	BusinessId                int                             `json:"businessId"`
	BusinessTransactionTypeId int                             `json:"businessTransactionTypeId"`
	BusinessTransactionItemId int                             `json:"businessTransactionItemId"`
	Total                     int                             `json:"total"`
	Multiplier                int                             `json:"multiplier"`
	Date                      time.Time                       `json:"date"`
	Description               string                          `json:"description"`
	BusinessTransactionItem   BusinessTransactionItemResponse `json:"businessTransactionItem"`
}
