package config

import (
	"GIN-CRUD-SAMPLE-PROJECT/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	db, err := gorm.Open(postgres.Open("postgres://server:postgres@localhost:5432/pg"), &gorm.Config{})
	if err != nil {
		panic(err) //Panic is a built-in function that stops the ordinary flow of control and begins panicking.
	}
	db.AutoMigrate(&models.User{})
	DB = db
}
