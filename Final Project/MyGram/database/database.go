package database

import (
	"MyGram/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "postgresGO"
	dbPort   = "5432"
	dbName   = "db_mygram"
	db       *gorm.DB
	err      error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbName, dbPort)
	dsn := config
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("error connecting to database :", err)
	}

	fmt.Println("Successfully connected to database")
	err := db.Debug().AutoMigrate(models.User{}, models.Photo{}, models.Comment{}, models.SocialMedia{})
	if err != nil {
		return
	}
}

func GetDB() *gorm.DB {
	return db
}
