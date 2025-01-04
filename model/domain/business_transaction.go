package domain

import "time"

type BusinessTransaction struct {
	Id                        int
	BusinessId                int
	BusinessTransactionTypeId int
	BusinessTransactionItemId int
	Total                     int
	Multiplier                int
	Date                      time.Time
	Description               string
	CreatedAt                 string
	UpdatedAt                 string
}
