package web

import "time"

type BusinessStatsResponse struct {
	Name                       string                `json:"name"`
	DateStarted                string                `json:"dateStarted"`
	DateEnded                  string                `json:"dateEnded"`
	BusinessTransactionTypeIds []int                 `json:"businessTransactionTypeIds"`
	BusinessTransactionItemIds []int                 `json:"businessTransactionItemIds"`
	Businesses                 []Business            `json:"businesses"`
	BusinessTransactionType    []TransactionTypeStat `json:"businessTransactionType"`
	Total                      int                   `json:"total"`
	TotalSell                  int                   `json:"totalSell"`
	TotalBuy                   int                   `json:"totalBuy"`
}

type Business struct {
	Id                       int                   `json:"id"`
	Name                     string                `json:"name"`
	BusinessTransactionItems []TransactionItemStat `json:"businessTransactionItem"`
	Total                    int                   `json:"total"`
	TotalSell                int                   `json:"totalSell"`
	TotalBuy                 int                   `json:"totalBuy"`
}

type TransactionItemStat struct {
	Name                     string                `json:"name"`
	BusinessTransactionTypes []TransactionTypeStat `json:"businessTransactionType"`
	Total                    int                   `json:"total"`
	TotalSell                int                   `json:"totalSell"`
	TotalBuy                 int                   `json:"totalBuy"`
}

type TransactionTypeStat struct {
	Name  string            `json:"name"`
	Stats []TransactionStat `json:"stats"`
}

type TransactionStat struct {
	Date  time.Time `json:"date"`
	Total int       `json:"total"`
}
