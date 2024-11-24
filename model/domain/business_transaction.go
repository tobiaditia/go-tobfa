package domain

type BusinessTransaction struct {
	Id                        int
	BusinessId                int
	BusinessTransactionTypeId int
	BusinessTransactionItemId int
	Total                     int
	Multiplier                int
	Date                      string
	Description               string
	CreatedAt                 string
	UpdatedAt                 string
}
