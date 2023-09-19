package handlers

import (
	"fmt"
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

func Del(c *fiber.Ctx) error {
	needTitle := c.Params("title", "")
	if needTitle == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "bad params request",
		})
	}
	emptyEx := new(models.ToDo)
	if check := database.DB.Db.Where("title = ?", needTitle).Delete(emptyEx); check.RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": fmt.Sprintf("there is no todo with title = %s", needTitle),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "todo deleted successfully",
	})
}

func ChangeFlag(c *fiber.Ctx) error {
	needTitle := c.Params("title", "")
	if needTitle == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "bad params request",
		})
	}
	emptyEx := new(models.ToDo)
	if check := database.DB.Db.Where("title = ?", needTitle).First(&emptyEx); check.RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": fmt.Sprintf("there is no todo with title = %s", needTitle),
		})
	}
	emptyEx.Flag = !emptyEx.Flag
	database.DB.Db.Where("title = ?", needTitle).Save(&emptyEx)
	return c.Status(fiber.StatusOK).JSON(emptyEx)
}
