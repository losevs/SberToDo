package handlers

import (
	"log"
	"os"
	"sber/database"
	"sber/models"
	"testing"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	dsn = "host=localhost user=postgres password=Sergey26 dbname=TestSB port=5432 sslmode=disable"
)

var testQueries *database.Dbinstance

func TestMain(m *testing.M) {
	conn, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	conn.AutoMigrate(&models.ToDo{}, &models.ToDoRequest{})
	testQueries = &database.Dbinstance{
		Db: conn,
	}
	os.Exit(m.Run())
}
