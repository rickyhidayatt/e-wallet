package utils

import (
	"log"

	"gopkg.in/gomail.v2"
)

const (
	CONFIG_SMTP_PORT     = 587
	CONFIG_SMTP_HOST     = "smtp.gmail.com"
	CONFIG_AUTH_EMAIL    = "ricky.hidayatyt@gmail.com"
	CONFIG_AUTH_PASSWORD = "pass"
	CONFIG_SENDER_NAME   = "E-Wallet Plus62 <ricky.hidayatyt@gmail.com>"
)

func SendMail() {

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", CONFIG_SENDER_NAME)
	mailer.SetHeader("To", "recipient1@gmail.com")
	mailer.SetHeader("Subject", "Minta Uang")
	mailer.SetBody("text/html", "Hai ada temen kamu lagi butuh bantuan kamu nih !")

	dialer := gomail.NewDialer(
		CONFIG_SMTP_HOST,
		CONFIG_SMTP_PORT,
		CONFIG_AUTH_EMAIL,
		CONFIG_AUTH_PASSWORD,
	)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Mail sent!")

}
