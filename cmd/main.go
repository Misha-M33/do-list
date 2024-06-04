package main

import (
	"do-list/internal/database"
	"do-list/internal/service"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/joho/godotenv"
)

const port = ":3333"

func main() {
	godotenv.Load()
	log.Println("godotenv init = ok")

	custom := database.NewDatabase()
	err := custom.ConnectDB(
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	Secret_access := os.Getenv("ACCESS_KEY")
	// Secret_refresh := os.Getenv("REFRESH_KEY")

	if err != nil {
		log.Fatal("database do not connect", err)
	}

	log.Println("database is connect = ok")

	app := fiber.New()

	app.Post("/create/user", service.CreateUser(custom))
	app.Get("/create/table/users", service.CreateTabletUsers(custom))
	app.Get("/users", service.GetUsersAll(custom))
	app.Get("/groups", service.GetGroupsAll(custom))
	app.Get("/users/groups", service.GetUsersGroupsAll(custom))
	app.Get("/tasks", service.GetTasksAll(custom))
	app.Post("/registration", service.RegistrationUser(custom))

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(Secret_access)},
	}))
	app.Patch("/update/user", service.UpdateUserById(custom))
	app.Patch("/delete/user", service.DeleteUserOne(custom))
	app.Patch("/block/user", service.BlockUserOne(custom))
	app.Patch("/password/user", service.EditPasswordUser(custom))
	app.Post("/users/group", service.GetUsersByGroup(custom))

	app.Post("/create/task", service.CreateTaskOne(custom))
	app.Post("/get/tasks/by/group", service.GetTasksByGroupId(custom))
	app.Post("/get/tasks/by/user/id", service.GetTasksByUserId(custom))
	app.Get("/list/tasks", service.GetListTasks(custom))
	app.Delete("/delete/task", service.DeleteTaskById(custom))
	app.Patch("/task", service.UpdateTaskById(custom))

	app.Post("/create/group", service.CreateGroupOne(custom))
	app.Patch("/update/group", service.UpdateGroupOne(custom))
	app.Delete("/delete/group", service.DeleteGroupOne(custom))

	app.Post("/user/group", service.CreateUserGroupOne(custom))
	app.Delete("/user/group", service.DeleteUserGroupOne(custom))

	if app.Listen(port); err != nil {
		log.Fatal("Server Dead, read err: ", err)
	}

}
