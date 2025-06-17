package database

import "log"

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	dsn := "host=auth_db user=postgres password=postgres dbname=auth_mic port=5432 sslmode=disable"

	DB, err = gorm.Open(postgres.Open(dsn), gorm.Config{})
	if err != nil {
		log.Fatalf("")
	}
}
