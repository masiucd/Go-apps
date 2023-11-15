package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"net/smtp"
	"os"
)

func main() {
	var to string
	fmt.Println("Enter your email")
	fmt.Scan(&to)

	var message string
	fmt.Println("Enter your message")
	fmt.Scan(&message)

	from := loadConfig("EMAIL")
	password := loadConfig("APP_PASS")

	smtpHost := loadConfig("SMTP_HOST")
	port := loadConfig("SMTP_PORT")

	headers := map[string]string{
		"From":         from,
		"To":           to,
		"Subject":      "Test email",
		"Content-Type": "text/html; charset=UTF-8",
	}

	var body string
	for k, v := range headers {
		body += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	body += "\r\n" + message

	auth := smtp.PlainAuth("", from, password, smtpHost)
	err := smtp.SendMail(smtpHost+":"+port, auth, from, []string{to}, []byte(body))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email sent successfully!")

}

func loadConfig(key string) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error loading .env file")
	}
	return os.Getenv(key)
}
