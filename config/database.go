package config

import (
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
	dsn := "host=localhost user=dbwriter password=dbwriter dbname=timeslotdb port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
