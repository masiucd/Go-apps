# Send Email in Golang

Sending emails in Golang is easy using the built-in smtp package.

## Code Walkthrough

```go
// sendEmail sends an email using the built-in smtp package in Golang.
// It takes in the recipient email address, the subject of the email, and the body of the email as arguments.
// The sender email address and password are hardcoded in the function for simplicity.
// Returns an error if there was an issue sending the email.
func sendEmail(to, subject, body string) error {
// SMTP server configuration
smtpServer := "smtp.gmail.com"
smtpPort := "587"

// Sender email address and password
from := "sender@gmail.com"
password := "password"

// Message to send
message := []byte("To: " + to + "\r\n" +
"Subject: " + subject + "\r\n" +
"\r\n" +
body + "\r\n")

// Authentication
auth := smtp.PlainAuth("", from, password, smtpServer)

// Sending email
err := smtp.SendMail(smtpServer+":"+smtpPort, auth, from, []string{to}, message)
if err != nil {
return err
}

return nil
}
```
