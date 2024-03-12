package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"assignment2/models"
)

var (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	dbname   = "belajar-golang"
	password = "alvaro123"
	db       *gorm.DB
	err      error
)

func StartDB() {
	config := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbname, password)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to database", err)
	}

	db.Debug().AutoMigrate(models.Order{}, models.Item{})
}

func GetDB() *gorm.DB {
	return db
}
