package database

import (
	"fmt"
	"log"
	"os"
	"sber/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Db *gorm.DB

func init() {
	err := godotenv.Load("C:\\Users\\Owner\\Documents\\GoLang\\SberToDo\\.env")
	if err != nil {
		fmt.Println(err)
	}
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	log.Println("connected")
	Db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("running migrations")
	Db.AutoMigrate(&models.ToDo{})
}
