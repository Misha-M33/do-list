package service

import (
	"do-list/internal/database"
	"do-list/internal/entities"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateTaskOne(pg *database.DB) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		var t entities.Task
		if err := c.BodyParser(&t); err != nil {
			log.Println(err)
			return c.Status(500).SendString(err.Error())
		}

		if err := t.ValidateCreateTask(); err != nil {
			log.Println("Create task has error", err)
			return c.Status(400).SendString(err.Error())
		}

		userID, err := GetUserIdFromHeader(c)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		t.Creator, _ = uuid.Parse(userID)

		if err := pg.CreateTask(&t); err != nil {
			log.Println(err)
			return c.Status(500).SendString(err.Error())
		}

		return c.Status(200).JSON(t)
	}
}

func GetTasksAll(pg *database.DB) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		var tasks entities.Tasks
		err := pg.GetTasks(&tasks)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		// if result == 0 {
		// 	return c.Status(500).SendString("GetGroups faild")
		// }

		return c.Status(200).JSON(&tasks)
	}

}

func GetTasksByGroupId(pg *database.DB) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		var t entities.Task
		if err := c.BodyParser(&t); err != nil {
			log.Println("BodyParser has error:", err)
			return c.Status(500).SendString(err.Error())
		}

		if err := t.ValidateGetTaskByGroupId(); err != nil {
			log.Println("Get task by group id has error:", err)
			return c.Status(400).SendString(err.Error())
		}

		var tasks entities.Tasks
		if err := pg.GetTaskByGroup(&t, &tasks); err != nil {
			log.Println(err)
			return c.Status(500).SendString(err.Error())
		}

		return c.Status(200).JSON(&tasks)
	}
}

func GetTasksByUserId(pg *database.DB) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		var u entities.User
		if err := c.BodyParser(&u); err != nil {
			log.Println(err)
			return c.Status(500).SendString(err.Error())
		}

		userID, err := GetUserIdFromHeader(c)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		u.Id, _ = uuid.Parse(userID)

		var tasks entities.Tasks
		if err := pg.TasksByUserId(&u, &tasks); err != nil {
			log.Println(err)
			return c.Status(500).SendString(err.Error())
		}

		return c.Status(200).JSON(tasks)
	}
}

func GetListTasks(pg *database.DB) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		var tasks entities.Tasks
		if err := pg.GetTasks(&tasks); err != nil {
			log.Println(err)
			return c.Status(500).SendString(err.Error())
		}

		return c.Status(200).JSON(tasks)
	}
}

func DeleteTaskById(pg *database.DB) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		var t entities.Task
		if err := c.BodyParser(&t); err != nil {
			log.Println(err)
			return c.Status(500).SendString(err.Error())
		}

		if err := t.ValidateDeleteTaskById(); err != nil {
			log.Println("Delete task by id has error:", err)
			return c.Status(400).SendString(err.Error())
		}

		if err := pg.DeleteTask(&t); err != nil {
			log.Println(err)
			return c.Status(500).SendString(err.Error())
		}

		return c.Status(200).JSON(t.Id)
	}
}

func UpdateTaskById(pg *database.DB) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		var t entities.Task
		if err := c.BodyParser(&t); err != nil {
			log.Println(err)
			return c.Status(500).SendString(err.Error())
		}

		if err := t.ValidateUpdateTaskById(); err != nil {
			log.Println("Update task by id has error:", err)
			return c.Status(400).SendString(err.Error())
		}

		if err := pg.UpdateTask(&t); err != nil {
			log.Println(err)
			return c.Status(500).SendString(err.Error())
		}

		return c.Status(200).JSON(t)
	}
}
