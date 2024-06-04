package service

import (
	"do-list/internal/database"
	"do-list/internal/entities"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetUsersGroupsAll(pg *database.DB) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		var usersGroups entities.UsersGroups
		result, err := pg.GetUsersGroups(&usersGroups)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		if result == 0 {
			return c.Status(500).SendString("GetUsersGroups faild")
		}

		return c.Status(200).JSON(&usersGroups)
	}

}

func CreateUserGroupOne(pg *database.DB) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		var g entities.UserGroup
		if err := c.BodyParser(&g); err != nil {
			log.Println(err)
			return c.Status(500).SendString(err.Error())
		}

		if err := g.ValidateCreateUserGroup(); err != nil {
			log.Println("Create user_group validation error", err)
			return c.Status(500).SendString(err.Error())
		}

		userID, err := GetUserIdFromHeader(c)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		g.User_id, _ = uuid.Parse(userID)

		if err := pg.CreateUserGroup(&g); err != nil {
			log.Println(err)
			return c.Status(500).SendString(err.Error())
		}

		return c.Status(200).JSON(g)
	}
}

func DeleteUserGroupOne(pg *database.DB) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		var g entities.UserGroup
		if err := c.BodyParser(&g); err != nil {
			log.Println(err)
			return c.Status(500).SendString(err.Error())
		}

		if err := g.ValidateDeleteUserGroup(); err != nil {
			log.Println("DElete user_group validation error:", err)
			return c.Status(500).SendString(err.Error())
		}

		result, err := pg.DeleteUserGroup(&g)
		if err != nil {
			log.Println(err)
			return c.Status(500).SendString(err.Error())
		}

		if result == 0 {
			return c.Status(400).SendString("not deleted group:")
		}
		log.Println("deleted successfully:")

		return c.Status(200).JSON(g)
	}
}
