package handlers

import (
	"sber/database"
	"sber/models"

	"github.com/gofiber/fiber/v2"
)

func Show(c *fiber.Ctx) error {
	query := []models.ToDo{}
	if check := database.DB.Db.Find(&query); check.RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "todo list is empty",
		})
	}
	return c.Status(fiber.StatusOK).JSON(query)
}

func Add(c *fiber.Ctx) error {
	query := models.ToDo{}
	if err := c.BodyParser(&query); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err,
		})
	}
	database.DB.Db.Create(&query)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"todo":    query,
		"message": "todo created",
	})
}
