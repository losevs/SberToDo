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

func Show(c *fiber.Ctx) error {
	query := []models.ToDo{}
	if check := database.DB.Db.Order("date asc").Find(&query); check.RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "todo list is empty",
		})
	}
	return c.Status(fiber.StatusOK).JSON(query)
}

func Add(c *fiber.Ctx) error {
	query := models.ToDoRequest{}
	if err := c.BodyParser(&query); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err,
		})
	}
	// Check Title if exists
	emptyQuery := new(models.ToDo)
	if check := database.DB.Db.Where("title = ?", query.Title).First(&emptyQuery); check.RowsAffected != 0 {
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
	database.DB.Db.Create(&queryToAdd)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"todo":    queryToAdd,
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

func PatchToDo(c *fiber.Ctx) error {
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
