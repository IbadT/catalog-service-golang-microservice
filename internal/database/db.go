package database

import (
	"log"

	"github.com/IbadT/catalog-service-golang-microservice.git/internal/catalog"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	dsn := "host=auth_db user=postgres password=postgres dbname=auth_mic port=5432 sslmode=disable"

	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return nil, err
	}

	err = DB.AutoMigrate(
		catalog.Product{},
		catalog.Category{},
	)
	if err != nil {
		log.Fatalf("Failed to migrate models: %v", err)
	}

	return DB, nil
}
