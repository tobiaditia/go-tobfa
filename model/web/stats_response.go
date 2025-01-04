package web

import "time"

type StatsResponse struct {
	Date  time.Time `json:"date"`
	Total int       `json:"total"`
}
