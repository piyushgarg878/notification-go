package utils

import (
	"fmt"
	"log"
	"net/smtp"
	"notification-service-go/internal/config"
)

func SendEmail(to, subject, message string) error {
	fmt.Printf("[EMAIL] To: %s | Subject: %s | Message: %s\n", to, subject, message)
	cfg := config.LoadConfig()
	auth := smtp.PlainAuth("", cfg.EmailFrom, cfg.EmailPassword, cfg.SmtpHost)
	msg := []byte(fmt.Sprintf(
		"To: %s\r\nSubject: %s\r\n\r\n%s\r\n",
		to, subject, message,
	))

	err := smtp.SendMail(cfg.SmtpHost+":"+cfg.SmtpPort, auth, cfg.EmailFrom, []string{to}, msg)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Email sent successfully!")
	
	return nil
}

