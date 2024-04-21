package main

import (
	"github.com/gofiber/fiber/v2"
	"go-apps.com/file-upload2.0/src/db"
)

const port = ":8000"

func main() {
	app := fiber.New()
	db.InitDB()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	err := app.Listen(port)
	if err != nil {
		panic(err)
	}

}
