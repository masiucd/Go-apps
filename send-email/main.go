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
=======
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
>>>>>>> cf5d338ee8e5ae8e9a7f42931cbd201d517a952f
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error loading .env file")
	}
	return os.Getenv(key)
}
