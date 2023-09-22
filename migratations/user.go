package main

import (
	"fmt"

	"go-admin/app/admin/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=123456 dbname=admin port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database: " + err.Error())
	}

	err = db.AutoMigrate(&models.FrontUser{})
	if err != nil {
		panic("failed to migrate database: " + err.Error())
	}

	fmt.Println("Database migration completed.")
}
