package mailer

import (
	"bytes"
	"errors"
	"fmt"

	"text/template"

	gomail "gopkg.in/mail.v2"
)

type mailtrapClient struct {
	fromEmail string
	apiKey    string
}

func NewMailTrapClient(apiKey, fromEmail string) (mailtrapClient, error) {
	if apiKey == "" {
		return mailtrapClient{}, errors.New("api key is required")
	}

	return mailtrapClient{
		fromEmail: fromEmail,
		apiKey:    apiKey,
	}, nil
}

func (m mailtrapClient) Send(templateFile, username, email string, data any, isSandbox bool) (int, error) {
	// Template parsing and building
	tmpl, err := template.ParseFS(FS, "templates/"+templateFile)
	if err != nil {
		fmt.Printf("Error parsing template (file: %s): %v\n", templateFile, err)
		return -1, err
	}

	subject := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(subject, "subject", data)
	if err != nil {
		fmt.Printf("Error executing subject template (file: %s): %v\n", templateFile, err)
		return -1, err
	}

	body := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(body, "body", data)
	if err != nil {
		fmt.Printf("Error executing body template (file: %s): %v\n", templateFile, err)
		return -1, err
	}

	message := gomail.NewMessage()
	message.SetHeader("From", m.fromEmail)
	message.SetHeader("Reply-To", email)
	message.SetHeader("Subject", subject.String())
	message.AddAlternative("text/html", body.String())

	// Log email sending details
	fmt.Printf("Sending email (from: %s, to: %s, subject: %s)\n", m.fromEmail, email, subject.String())

	dialer := gomail.NewDialer("live.smtp.mailtrap.io", 587, "api", m.apiKey)

	// Log the dialer settings (excluding sensitive information like API key)
	fmt.Printf("Dialer configured (smtpServer: live.smtp.mailtrap.io, port: 587)\n")

	// Attempt to send the email and log any errors
	if err := dialer.DialAndSend(message); err != nil {
		fmt.Printf("Error sending email to %s: %v\n", email, err)
		return -1, err
	}

	// Log success
	fmt.Printf("Email sent successfully to %s\n", email)
	return 200, nil
}
