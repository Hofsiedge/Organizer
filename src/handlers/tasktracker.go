package handlers

import (
	"fmt"
	db "github.com/Hofsiedge/Organizer/src/database"
	"github.com/Hofsiedge/Organizer/src/utils"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

// TODO: return JSON errors
func GetTask(ac *utils.AppContext, c *fiber.Ctx) error {
	dbPool, sugar := ac.DbPool, ac.Logger
	taskId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		err = fmt.Errorf("Invalid URL parameter 'id' in GetTask: %v\n", err)
		sugar.Error(err)
		return fiber.NewError(fiber.StatusBadRequest, "Invalid task id: id must be an integer")
	}

	task, err := db.GetTask(dbPool, taskId)
	if err != nil {
		err = fmt.Errorf("DB error: %v\n", err)
		sugar.Error(err)
		return fiber.NewError(fiber.StatusNotFound, "Task not found")
	}

	return c.JSON(task)
}

type taskPostStructure struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	OwnerId     int64  `json:"ownerId"` // TODO: extract from JWT
}

// TODO: return JSON errors
// TODO: check if there is a way to make BodyParser fail on empty JSON
func CreateTask(ac *utils.AppContext, c *fiber.Ctx) error {
	dbPool, sugar := ac.DbPool, ac.Logger
	b := new(taskPostStructure)
	if err := c.BodyParser(b); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	if b.OwnerId == 0 || b.Name == "" {
		err := fmt.Errorf("missing or incorrectly formatted parameters." +
			"Body should contain `ownerId`, `name` and `description`")
		sugar.Error(err)
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	task, err := db.CreateTask(dbPool, b.Name, b.Description, b.OwnerId)
	if err != nil {
		err = fmt.Errorf("DB error: %v\n", err)
		sugar.Error(err)
		return fiber.NewError(fiber.StatusBadRequest, "Could not create a task")
	}

	return c.JSON(fiber.Map{
		"id": task,
	})
}
