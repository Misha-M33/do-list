package service

import (
	"do-list/internal/database"
	"do-list/internal/entities"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetGroupsAll(pg *database.DB) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		var groups entities.Groups
		result, err := pg.GetGroups(&groups)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		if result == 0 {
			return c.Status(500).SendString("GetGroups faild")
		}

		return c.Status(200).JSON(&groups)
	}

}

func CreateGroupOne(pg *database.DB) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		var g entities.Group
		if err := c.BodyParser(&g); err != nil {
			log.Println(err)
			return c.Status(500).SendString(err.Error())
		}

		if err := g.ValidateCreateGroup(); err != nil {
			log.Println("Create group validation error", err)
			return c.Status(400).SendString(err.Error())
		}

		userID, err := GetUserIdFromHeader(c)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		g.Owner_id, _ = uuid.Parse(userID)

		if err := pg.CreateGroup(&g); err != nil {
			log.Println(err)
			return c.Status(500).SendString(err.Error())
		}

		return c.Status(200).JSON(g)
	}
}

func UpdateGroupOne(pg *database.DB) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		var g entities.Group
		if err := c.BodyParser(&g); err != nil {
			log.Println(err)
			return c.Status(500).SendString(err.Error())
		}

		if err := g.ValidateUpdateGroup(); err != nil {
			log.Println("Update group validation error", err)
			return c.Status(400).SendString(err.Error())
		}

		userID, err := GetUserIdFromHeader(c)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		g.Owner_id, _ = uuid.Parse(userID)

		if err := pg.UpdateGroup(&g); err != nil {
			log.Println(err)
			return c.Status(500).SendString(err.Error())
		}

		return c.Status(200).JSON(g)
	}
}

func DeleteGroupOne(pg *database.DB) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		var g entities.Group
		if err := c.BodyParser(&g); err != nil {
			return c.Status(500).SendString(err.Error())
		}

		if err := g.ValidateDeleteGroup(); err != nil {
			log.Println("Delete group validation error", err)
			return c.Status(400).SendString(err.Error())
		}

		userID, err := GetUserIdFromHeader(c)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		g.Owner_id, _ = uuid.Parse(userID)

		result, err := pg.DeleteGroup(&g)
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
