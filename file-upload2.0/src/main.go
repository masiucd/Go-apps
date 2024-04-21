package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"go-apps.com/file-upload2.0/src/db"
)

const port = ":4000"

func main() {
	engine := html.New("./src/views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
		// ViewsLayout: "layouts/main",
	})

	app.Static("/public", "./src/public")

	// read the file contents in src/public by reading the dir
	files, _ := ioutil.ReadDir("./src/public")
	fmt.Println(files)

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
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		fmt.Println("NAME ", file.Filename)
		fmt.Println("SIZE ", file.Size)

		// open the file
		f, err := file.Open()
		if err != nil {
			fmt.Printf("Failed to open the file: %v\n", err)
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		// closing the file
		defer f.Close()

		fileBytes, err := io.ReadAll(f)
		if err != nil {
			fmt.Printf("Failed to read the file: %v\n", err)
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		// store the file in the database
		db.DB.Create(&db.Attachment{
			FileName: file.Filename,
			Blob:     fileBytes,
		})

		return c.Render("index", fiber.Map{"Title": "File Upload"})
	})

	// Route to get a file by id and serve it
	app.Get("/file/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		var attachment db.Attachment
		db.DB.First(&attachment, id)
		if attachment.ID == 0 {
			return c.Status(fiber.StatusNotFound).SendString("File not found")
		}

		// check if attachment is a pdf
		if strings.Contains(attachment.FileName, ".pdf") {
			// set the content type to pdf
			c.Set("Content-Type", "application/pdf")
		}

		return c.Send(attachment.Blob)
	})

	app.Get("/attachments", func(c *fiber.Ctx) error {
		var attachments []db.Attachment
		db.DB.Find(&attachments)

		type AttachmentUi struct {
			ID   uint
			Name string
		}
		var attachmentsUi []AttachmentUi
		for _, attachment := range attachments {
			attachmentsUi = append(attachmentsUi, AttachmentUi{attachment.ID, attachment.FileName})
		}
		return c.Status(fiber.StatusOK).Render("attachments", fiber.Map{
			"Title":       "Attachments",
			"Attachments": attachmentsUi,
		})
	})

	err := app.Listen(port)
	if err != nil {
		panic(err)
	}

}
