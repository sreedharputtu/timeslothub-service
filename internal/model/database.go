package model

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "dbwriter"
	password = ""
	db       = "timeslotdb"
)

func Database() (*gorm.DB, error) {
	sqlUrl := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, db)
	return gorm.Open(postgres.Open(sqlUrl), &gorm.Config{})
}
