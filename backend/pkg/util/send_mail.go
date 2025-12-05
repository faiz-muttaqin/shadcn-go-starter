package util

import (
	"fmt"
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
	"gopkg.in/gomail.v2"
)

func SendEmail(subject, to, body string) error {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", os.Getenv("APP_NAME")+" <"+os.Getenv("CONFIG_SMTP_SENDER")+">")
	mailer.SetHeader("To", to)
	mailer.SetHeader("Subject", "[noreply] "+subject)
	mailer.SetBody("text/html", body)

	smtpPortStr := os.Getenv("CONFIG_SMTP_PORT")
	smtpPort, oops := strconv.Atoi(smtpPortStr)
	if oops != nil {
		smtpPort = 587
	}
	dialer := gomail.NewDialer(
		os.Getenv("CONFIG_SMTP_HOST"),
		smtpPort,
		os.Getenv("CONFIG_AUTH_EMAIL"),
		os.Getenv("CONFIG_AUTH_PASSWORD"),
	)

	return dialer.DialAndSend(mailer)
}

// SendEmailDynamic mengirim email dengan parameter dinamis
func SendEmailDynamic(to []string, cc []string, subject string, body string) error {
	mailer := gomail.NewMessage()

	// Set "From"
	fromEmail := os.Getenv("CONFIG_SMTP_SENDER")
	mailer.SetHeader("From", "Hommypay <"+fromEmail+">")

	// Set "To"
	if len(to) > 0 {
		mailer.SetHeader("To", to...)
	} else {
		return fmt.Errorf("recipient (to) list cannot be empty")
	}

	// Set "Cc" (hanya jika ada)
	if cc != nil {
		if len(cc) > 0 {
			mailer.SetHeader("Cc", cc...)
		}
	}

	// Set "Subject"
	mailer.SetHeader("Subject", subject)

	// Set body
	mailer.SetBody("text/html", body)

	// SMTP config
	smtpHost := os.Getenv("CONFIG_SMTP_HOST")
	smtpPortStr := os.Getenv("CONFIG_SMTP_PORT")
	smtpPort, err := strconv.Atoi(smtpPortStr)
	if err != nil {
		logrus.Error(err)
		smtpPort = 587 // default port jika gagal parsing
	}
	smtpUser := os.Getenv("CONFIG_AUTH_EMAIL")
	smtpPass := os.Getenv("CONFIG_AUTH_PASSWORD")

	// Buat dialer
	dialer := gomail.NewDialer(smtpHost, smtpPort, smtpUser, smtpPass)

	// Kirim email
	return dialer.DialAndSend(mailer)
}
