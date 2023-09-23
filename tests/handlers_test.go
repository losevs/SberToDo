package tests

import (
	"io"
	"net/http"
	"sber/server"
	"testing"

	// "github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

// func TestAdd(t *testing.T) {
// 	arg := models.ToDoRequest{
// 		Title:       "checkFstt",
// 		Description: "descFst",
// 		Date:        "01.01.2001",
// 		Flag:        false,
// 	}

// 	app := fiber.New()
// 	app.Post("/add", Add)

// 	marshalled, _ := json.Marshal(arg)
// 	req, _ := http.NewRequest(http.MethodPost, "/add", bytes.NewBuffer(marshalled))
// 	req.Header.Set("Content-Type", "application/json")
// 	resp, _ := app.Test(req, -1)

// 	bod, _ := io.ReadAll(resp.Body)
// 	defer resp.Body.Close()

// 	t.Log(resp.StatusCode)
// 	t.Log(string(bod))
// 	assert.Equal(t, 200, resp.StatusCode)
// }

// err := godotenv.Load("C:\\Users\\Owner\\Documents\\GoLang\\SberToDo\\.env")
// if err != nil {
// 	t.Fatal("Error loading .env file")
// }

var app = server.New()

func TestShow(t *testing.T) {
	tests := []struct {
		testName     string
		route        string
		method       string
		body         io.Reader
		expectedCode int
		// expectedBody string
	}{
		{
			testName:     "valid",
			route:        "/show",
			method:       "GET",
			body:         nil,
			expectedCode: 200,
		},
	}
	for _, test := range tests {
		req, _ := http.NewRequest(
			test.method,
			test.route,
			test.body,
		)

		res, err := app.Test(req, -1)
		assert.Nilf(t, err, test.testName)
		assert.Equalf(t, test.expectedCode, res.StatusCode, test.testName)
	}
}

// req, _ := http.NewRequest(http.MethodGet, "/show", nil)
// 	resp, _ := app.Test(req, -1)

// 	bod, _ := io.ReadAll(resp.Body)
// 	defer resp.Body.Close()

// 	t.Log(resp.StatusCode)
// 	t.Log(string(bod))

// 	assert.Equal(t, 200, resp.StatusCode)

// func TestShow(t *testing.T) {
// 	app := fiber.New()
// 	app.Get("/show", Show)
// 	req, _ := http.NewRequest(http.MethodGet, "/show", nil)
// 	resp, _ := app.Test(req, -1)

// 	t.Log(resp.StatusCode)
// 	t.Log(resp.Body)
// 	assert.Equal(t, 200, resp.StatusCode)
// }

//----
// main_test.go
// package handlers

// import (
// 	"log"
// 	"os"
// 	"sber/database"
// 	"sber/models"
// 	"testing"

// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// const (
// 	dsn = "host=localhost user=postgres password=Sergey26 dbname=TestSB port=5432 sslmode=disable"
// )

// var testQueries *database.Dbinstance

// func TestMain(m *testing.M) {
// 	conn, err := gorm.Open(postgres.Open(dsn))
// 	if err != nil {
// 		log.Fatal("cannot connect to db:", err)
// 	}
// 	conn.AutoMigrate(&models.ToDo{}, &models.ToDoRequest{})
// 	testQueries = &database.Dbinstance{
// 		Db: conn,
// 	}
// 	os.Exit(m.Run())
// }
