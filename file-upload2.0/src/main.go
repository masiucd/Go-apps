package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"go-apps.com/file-upload2.0/src/db"
)

const port = ":8000"

func main() {
	engine := html.New("./src/views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
		// ViewsLayout: "layouts/main",
	})
	db.InitDB()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index",
			fiber.Map{
				"Title": "File Upload",
			},
		)
	})

	app.Get("/upload", func(c *fiber.Ctx) error {
		return c.Render("upload",
			fiber.Map{
				"Title": "Upload File",
			},
		)

	})

	err := app.Listen(port)
	if err != nil {
		panic(err)
	}

}
