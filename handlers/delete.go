package handlers

import (
	"fmt"
	"sber/database"
	"sber/models"

	"github.com/gofiber/fiber/v2"
)

// DeleteTags		godoc
// @Summary Удаление задачи по заголовку
// @Description Удаляет задачу ToDo по заголовку.
// @Tags Delete
// @Param title path string true "Заголовок задачи"
// @Produce application/json
// @Success 200 {string} string "todo deleted successfully"
// @Router /del/{title} [delete]
func Del(c *fiber.Ctx) error { // Delete ToDo by Title
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
