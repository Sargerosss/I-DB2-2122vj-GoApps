package main

import (
	"fmt"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

func mailer(user User) {
	godotenv.Load()
	emailTo := os.Getenv("EMAIL_TO")
	var fromEmail string
	var emailFile string
	var password string
	fmt.Println("Please enter your email address")
	fmt.Scanln(&fromEmail)
	fmt.Println("Please enter your password")
	fmt.Scanln(&password)
	fmt.Println("Please enter a file")
	fmt.Scanln(&emailFile)
	toEmail := []string{emailTo}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	data, errors := os.ReadFile(emailFile)
	checkError(errors)

	message := []byte(data)

	auth := smtp.PlainAuth("", fromEmail, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, fromEmail, toEmail, message)
	checkError(err)
	fmt.Println("Email sent succesfully", user.Username)
}
