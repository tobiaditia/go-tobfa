package web

type BusinessTransactionStatsResponse struct {
	Averange         int                      `json:"averange"`
	BusinessCategory BusinessCategoryResponse `json:"businessCategory"`
	Stats            []StatsResponse          `json:"stats"`
}
