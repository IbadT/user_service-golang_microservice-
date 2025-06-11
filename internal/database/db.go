package database

import (
	"log"

	"github.com/IbadT/user_service-golang_microservice-.git/internal/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	dsn := "host=db user=postgres password=postgres dbname=u_mic port=5432 sslmode=disable"

	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
		return nil, err
	}

	if err := DB.AutoMigrate(&user.User{}); err != nil {
		log.Fatalf("Could not migrate database: %v", err)
		return nil, err
	}

	return DB, nil
}
