package main

import (
	"fmt"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	from := config("EMAIL")
	pass := config("APP_PASS")
	to := []string{config("EMAIL")}
	host := "smtp.gmail.com"
	port := "587"
	address := host + ":" + port

	subject := "Test email from Go!"
	body := "This is the body of an email message."
	message := []byte(subject + "\n" + "https://www.google.com" + "\n" + body)

	auth := smtp.PlainAuth("Hello", from, pass, host)

	err := smtp.SendMail(address, auth, from, to, message)

	if err != nil {
		panic(err)
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
