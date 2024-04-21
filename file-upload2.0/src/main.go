package main

import (
	"fmt"

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

	app.Post("/upload", func(c *fiber.Ctx) error {
		file, err := c.FormFile("file")
		if err != nil {
			// TODO render error page with error message
			fmt.Printf("Failed to get the file: %v\n", err)
			return c.SendString(err.Error())
		}
		fmt.Println("NAME ", file.Filename)
		fmt.Println("SIZE ", file.Size)
		return c.Render("index", fiber.Map{"Title": "File Upload"})
	})

	err := app.Listen(port)
	if err != nil {
		panic(err)
	}

}
