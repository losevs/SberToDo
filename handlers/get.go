package handlers

import (
	"fmt"
	"sber/database"
	"sber/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// @Summary Show all ToDo's
// @Description Show all.
// @Success 200 {array} models.ToDo{}
// @Router /show [get]
func Show(c *fiber.Ctx) error { // Show all ToDo's
	query := []models.ToDo{}
	if check := database.DB.Db.Find(&query); check.RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "todo list is empty",
		})
	}
	return c.Status(fiber.StatusOK).JSON(query)
}

// @Summary Изменение флага выполнения задачи ToDo
// @Description Меняет флаг выполнения задачи ToDo по ее заголовку.
// @Accept json
// @Produce json
// @Param title path string true "Заголовок задачи"
// @Success 200 {object} models.ToDo
// @Router /flag/{title} [get]
func ChangeFlag(c *fiber.Ctx) error { // Change Flag true/false by title
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

// TruePagTags		godoc
// @Summary Отображение списка задач с фильтром по флагу выполнения и пагинацией
// @Description Получает список задач ToDo с фильтром по флагу выполнения и пагинацией.
// @Produce json
// @Param flag query bool false "Флаг выполнения задачи (true/false)"
// @Param page query int false "Номер страницы для пагинации"
// @Success 200 {array} ToDoTask
// @Router /true [get]
func TruePag(c *fiber.Ctx) error { // Pagination true
	needPage, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err,
		})
	}
	todos := []models.ToDo{}
	if check := database.DB.Db.Offset((needPage-1)*3).Limit(3).Where("flag = ?", true).Find(&todos); check.RowsAffected == 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": fmt.Sprintf("there is no true todos on page=%d", needPage),
		})
	}
	return c.Status(fiber.StatusOK).JSON(todos)
}

// FalsePagTags		godoc
// @Summary Отображение списка задач с фильтром по флагу выполнения и пагинацией
// @Description Получает список задач ToDo с фильтром по флагу выполнения и пагинацией.
// @Produce json
// @Param flag query bool false "Флаг выполнения задачи (true/false)"
// @Param page query int false "Номер страницы для пагинации"
// @Success 200 {array} ToDoTask
// @Router /false [get]
func FalsePag(c *fiber.Ctx) error { // Pagination false
	needPage, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err,
		})
	}
	todos := []models.ToDo{}
	if check := database.DB.Db.Offset((needPage-1)*3).Limit(3).Where("flag = ?", false).Find(&todos); check.RowsAffected == 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": fmt.Sprintf("there is no false todos on page=%d", needPage),
		})
	}
	return c.Status(fiber.StatusOK).JSON(todos)
}

// FlagAscTags		godoc
// @Summary Отображение списка задач с фильтром по флагу выполнения и сортировкой по дате
// @Description Получает список задач ToDo с фильтром по флагу выполнения и сортировкой по дате.
// @Produce json
// @Param flag path bool true "Флаг выполнения задачи (true/false)"
// @Success 200 {array} ToDoTask
// @Router /date/{flag} [get]
func FlagAsc(c *fiber.Ctx) error { // Flag asc order
	needFlag, err := strconv.ParseBool(c.Params("flag", "false"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err,
		})
	}
	query := []models.ToDo{}
	if check := database.DB.Db.Where("flag = ?", needFlag).Order("date asc").Find(&query); check.RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": fmt.Sprintf("todo list with flag=%t is empty", needFlag),
		})
	}
	return c.Status(fiber.StatusOK).JSON(query)
}
