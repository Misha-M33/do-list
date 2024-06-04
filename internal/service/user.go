package service

import (
	"do-list/internal/database"
	"do-list/internal/entities"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

const DateTime = "2006-01-02 15:04:05" // for printing UpdatedAt in this format

func CreateUser(pg *database.DB) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		var u entities.User
		if err := c.BodyParser(&u); err != nil {
			return c.Status(400).SendString(err.Error())
		}

		pass := GetHashPassword(u.Password)
		err := pg.CreateUser(pass, &u)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		return c.Status(200).JSON(u)
	}
}

func GetUsersAll(pg *database.DB) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		var users entities.Users
		result, err := pg.GetUsers(&users)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		if result == 0 {
			return c.Status(500).SendString("GetUsers faild")
		}

		return c.Status(200).JSON(&users)
	}

}

func CreateTabletUsers(pg *database.DB) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		var user entities.User
		err := pg.TableUsers(&user)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		return c.Status(200).JSON(&user)
	}

}

func RegistrationUser(pg *database.DB) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		var u entities.User
		if err := c.BodyParser(&u); err != nil {
			return c.Status(400).SendString(err.Error())
		}

		if err := u.ValidateRegistration(); err != nil {
			log.Println("Create user validation error", err)
			return c.Status(400).SendString(err.Error())
		}

		id, hash, deleted, block, err := pg.GetUserByEmail(&u)
		if err != nil {
			log.Println(err)
			return c.Status(500).SendString(err.Error())
		}

		if deleted {
			log.Println("User deleted")
			return c.Status(500).SendString(err.Error())
		}
		if block {
			log.Println("User blocked")
			return c.Status(500).SendString(err.Error())
		}
		if ComparePassword(u.Password, hash) {
			accessToken, err := CreateAccessToken(id, c)
			if err != nil {
				c.Status(500).SendString(err.Error())
			}
			refreshToken, err := CreateRefreshToken(id, c)
			if err != nil {
				c.Status(500).SendString(err.Error())
			}
			tokens := &Tokens{
				AccessToken:  accessToken,
				RefreshToken: refreshToken,
			}
			if err := pg.SaveRefreshToken(id, refreshToken); err != nil {
				return c.Status(500).SendString(err.Error())
			}
			return c.Status(200).JSON(tokens)
		}
		return c.Status(500).SendString(err.Error())
	}
}

func UpdateUserById(pg *database.DB) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		var u entities.User
		if err := c.BodyParser(&u); err != nil {
			return c.Status(500).SendString(err.Error())
		}
		if err := u.ValidateUpdate(); err != nil {
			log.Println("Update user validation error", err)
			return c.Status(400).SendString(err.Error())
		}

		userID, err := GetUserIdFromHeader(c)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		u.Id, _ = uuid.Parse(userID)

		if err := pg.UpdateUser(&u); err != nil {
			return c.Status(400).SendString(err.Error())
		}

		return c.Status(200).JSON(u)
	}
}

func DeleteUserOne(pg *database.DB) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		var user entities.User
		if err := c.BodyParser(&user); err != nil {
			log.Println(err)
			return c.Status(500).SendString(err.Error())
		}

		userID, err := GetUserIdFromHeader(c)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		user.Id, _ = uuid.Parse(userID)

		result, err := pg.DeleteUser(&user)
		if err != nil {
			log.Println(err)
			return c.Status(500).SendString(err.Error())
		}

		if result == 0 {
			return c.Status(500).SendString("DeleteUser faild")
		}

		return c.Status(200).JSON(&user)
	}
}

func BlockUserOne(pg *database.DB) func(c *fiber.Ctx) error {
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

		result, err := pg.BlockUser(&u)
		if err != nil {
			log.Println(err)
			return c.Status(500).SendString(err.Error())
		}
		if result == 0 {
			return c.Status(500).SendString("BlockUserOne faild")
		}

		return c.Status(200).JSON(u)
	}
}

func EditPasswordUser(pg *database.DB) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		var u entities.User
		if err := c.BodyParser(&u); err != nil {
			log.Println(err)
			return c.Status(500).SendString(err.Error())
		}

		if err := u.ValidatePassword(); err != nil {
			log.Println("Update user validation error", err)
			return c.Status(400).SendString(err.Error())
		}

		userID, err := GetUserIdFromHeader(c)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		u.Id, _ = uuid.Parse(userID)

		if err := pg.EditPassword(GetHashPassword(u.Password), &u); err != nil {
			log.Println(err)
			return c.Status(500).SendString(err.Error())
		}

		return c.Status(200).JSON(&u)
	}
}

func GetUsersByGroup(pg *database.DB) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		var ug entities.UserGroup
		if err := c.BodyParser(&ug); err != nil {
			log.Println(err)
			return c.Status(500).SendString(err.Error())
		}

		userID, err := GetUserIdFromHeader(c)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		ug.User_id, _ = uuid.Parse(userID)

		var u entities.Users
		if err := pg.UsersGroup(&ug, &u); err != nil {
			log.Println(err)
			return c.Status(500).SendString(err.Error())
		}

		return c.Status(200).JSON(u)
	}
}
