package db

import (
	"fmt"
	"full_domain/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=12345 dbname=full_domain port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Connection with DB failed.")
	}
	db.AutoMigrate(entity.User{})
	return db
}

