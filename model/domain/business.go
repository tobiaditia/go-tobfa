package domain

import (
	"database/sql"
	"time"
)

type Business struct {
	Id                 int
	UserId             int
	Name               string
	Address            sql.NullString
	BusinessCategoryId int
	CountryId          int
	ProvinceId         int
	CityId             int
	DistrictId         int
	VillageId          int
	CreatedAt          time.Time
	UpdatedAt          time.Time
}
