package app

import (
	"database/sql"
	"fmt"
	"go-tobfa/helper"
	"time"

	"github.com/spf13/viper"
)

func NewDB(config *viper.Viper) *sql.DB {

	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true",
		config.GetString("DB_USER"),
		config.GetString("DB_PASS"),
		config.GetString("DB_HOST"),
		config.GetInt("DB_PORT"),
		config.GetString("DB_NAME"),
	)
	// "username:password@tcp(hostname:port)/database?parseTime=true"

	db, err := sql.Open("mysql", connectionString)
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
