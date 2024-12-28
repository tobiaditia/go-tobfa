package web

type BusinessStatsGetRequest struct {
	DateStarted                string `validate:"required" json:"dateStarted"`
	DateEnded                  string `validate:"required" json:"dateEnded"`
	BusinessIds                []int  `validate:"required" json:"businessIds"`
	BusinessTransactionTypeIds []int  `validate:"required" json:"businessTransactionTypeIds"`
	BusinessTransactionItemIds []int  `validate:"required" json:"businessTransactionItemIds"`
}
