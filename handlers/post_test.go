package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	// "sber/models"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	// "github.com/stretchr/testify/require"
)

func TestAdd(t *testing.T) {
	arg := struct {
		Title       string `json:"title"`
		Description string `json:"desc"`
		Date        string `json:"date"`
		Flag        bool   `json:"flag"`
	}{
		Title:       "checkFstt",
		Description: "descFst",
		Date:        "01.01.2001",
		Flag:        false,
	}

	app := fiber.New()
	app.Post("/add", Add)

	marshalled, _ := json.Marshal(arg)
	req, _ := http.NewRequest(http.MethodPost, "/add", bytes.NewBuffer(marshalled))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req, -1)

	bod, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()

	t.Log(resp.StatusCode)
	t.Log(string(bod))
	assert.Equal(t, 200, resp.StatusCode)
}

func TestShow(t *testing.T) {
	app := fiber.New()
	app.Get("/show", Show)
	req, _ := http.NewRequest(http.MethodGet, "/show", nil)
	resp, _ := app.Test(req, -1)

	t.Log(resp.StatusCode)
	t.Log(resp.Body)
	assert.Equal(t, 200, resp.StatusCode)
}
