package domain

import "time"

type BusinessTransaction struct {
	Id                        int
	BusinessId                int
	BusinessTransactionTypeId int
	BusinessTransactionItemId int
	Total                     int
	Quantity                  int
	Date                      time.Time
	Description               string
	CreatedAt                 string
	UpdatedAt                 string
}
