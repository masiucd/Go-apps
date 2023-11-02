package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/wneessen/go-mail"
)

func main() {
	m := mail.NewMsg()
	if err := m.From(config("EMAIL")); err != nil {
		log.Fatalf("failed to set From address: %s", err)
	}
	if err := m.To(config("EMAIL")); err != nil {
		log.Fatalf("failed to set To address: %s", err)
	}

	m.Subject("Hello Gopher!")
	m.SetBodyString(mail.TypeTextPlain, "This is the body of the email")

	c, err := mail.NewClient("smtp.gmail.com", mail.WithPort(25), mail.WithSMTPAuth(mail.SMTPAuthPlain),
		mail.WithUsername(config("EMAIL")), mail.WithPassword(config("APP_PASS")),
	)
	if err != nil {
		log.Fatalf("failed to create client: %s", err)
	}
	if err := c.DialAndSend(m); err != nil {
		log.Fatalf("failed to send mail: %s", err)
	}
}

// Config func to get env value
func config(key string) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error loading .env file")
	}
	return os.Getenv(key)
}
