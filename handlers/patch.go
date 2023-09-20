package handlers

import (
	"fmt"
	"sber/database"
	"sber/models"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

// PatchTags		godoc
// @Summary Изменение задачи по заголовку
// @Description Изменяет задачу ToDo по заголовку.
// @Param title path string true "Заголовок задачи"
// @Accept json
// @Produce json
// @Param task body ToDoTask true "Обновленная задача"
// @Success 200 {object} ToDoTask
// @Router /change/{title} [patch]
func PatchToDo(c *fiber.Ctx) error { // Update ToDo by Title
	needTitle := c.Params("title", "")
	if needTitle == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "bad params request",
		})
	}
	query := new(models.ToDoRequest)
	if err := c.BodyParser(&query); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err,
		})
	}
	FoundToDo := new(models.ToDo)
	if check := database.DB.Db.Where("title = ?", needTitle).First(&FoundToDo); check.RowsAffected == 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": fmt.Sprintf("there is no todo with title = %s", needTitle),
		})
	}
	if query.Description != "" {
		database.DB.Db.Model(&FoundToDo).Where("title = ?", needTitle).Update("description", query.Description)
	}
	if query.Date != "" {
		// Check Date
		splitDate := strings.Split(query.Date, ".")
		if len(splitDate) != 3 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "bad date request: need dd.mm.yyyy",
			})
		}
		//check year
		if len(splitDate[2]) != 4 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "bad year request",
			})
		}
		day, checkNum1 := strconv.Atoi(splitDate[0])
		if day < 1 || day > 31 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "bad date request: wrong day number",
			})
		}
		month, checkNum2 := strconv.Atoi(splitDate[1])
		if month < 1 || month > 12 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "bad date request: wrong month number",
			})
		}
		year, checkNum3 := strconv.Atoi(splitDate[2])
		if checkNum1 != nil || checkNum2 != nil || checkNum3 != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "bad date request: need dd.mm.yyyy in numbers",
			})
		}
		DateAdd := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
		database.DB.Db.Model(&FoundToDo).Where("title = ?", needTitle).Update("date", DateAdd)
	}
	return c.Status(fiber.StatusOK).JSON(FoundToDo)
}
