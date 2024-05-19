package main

import (
	"src/github.com/mustafaakilll/go-site-exam/db"
	"src/github.com/mustafaakilll/go-site-exam/internal/user"

	"github.com/gofiber/fiber/v2"
)

func main() {
	db.Connect()

	database := db.DB
	userRepository := user.NewUserRepository(database)
	userService := user.NewUserService(userRepository)
	userHandler := user.NewUserHandler(userService)

	app := fiber.New()
	api := app.Group("/api/v1")

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Hello, World!"})
	})

	userApi := api.Group("/users")
	userApi.Get("/", userHandler.GetUsers)
	userApi.Get("/:id", userHandler.GetUserByID)
	userApi.Post("/", userHandler.CreateUser)
	userApi.Put("/:id", userHandler.UpdateUser)
	userApi.Delete("/:id", userHandler.DeleteUser)
	userApi.Get("/email/:email", userHandler.GetUserByEmail)
	userApi.Get("/name/:username", userHandler.GetUserByUsername)
	userApi.Get("/teacher/:id", userHandler.SetTeacher)

	app.Listen(":8081")
}
