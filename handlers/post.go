package handlers

import (
	"sber/database"
	"sber/models"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

// CreateTags		godoc
// @Summary Добавление новой задачи
// @Description Добавляет новую задачу ToDo.
// @Tags Post
// @Accept json
// @Produce json
// @Param task body models.ToDoRequest true "Новая задача"
// @Success 200 {object} models.ToDo
// @Router /add [post]
func Add(c *fiber.Ctx) error { // Add new ToDo
	query := new(models.ToDoRequest)
	if err := c.BodyParser(&query); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "bad request",
		})
	}
	// Check Title if exists
	emptyQuery := new(models.ToDo)
	check := database.Db.Where("title = ?", query.Title).First(&emptyQuery)
	// if check.Error != nil {
	// 	return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
	// 		"message": "something with gorm - where, first",
	// 	})
	// }
	if check.RowsAffected != 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"todo":    emptyQuery,
			"message": "this title already exisits",
		})
	}

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
	queryToAdd := models.ToDo{
		Title:       query.Title,
		Description: query.Description,
		Date:        DateAdd,
		Flag:        query.Flag,
	}
	if checker := database.Db.Create(&queryToAdd); checker.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"message": "something with gorm",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"todo":    queryToAdd,
		"message": "todo created",
	})
}
