package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hidalgo27/art-digital.git/data"
	"github.com/hidalgo27/art-digital.git/db"
)

func main() {
	app := fiber.New()

	db.Init()
	db.DB.AutoMigrate(&data.User{})
	db.DB.Create(&data.User{
		Name:  "John Smith",
		Email: "john.smith@gmail.com",
	})
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Listen(":3000")
}
