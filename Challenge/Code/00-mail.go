package main

import (
	"database/sql"
	"fmt"
	"net/smtp"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func mailer(user User, db *sql.DB) {
	godotenv.Load()
	emailTo := os.Getenv("EMAIL_TO")
	var fromEmail string
	var emailFile string
	var password string
	fmt.Println("Please enter your email address")
	fmt.Scanln(&fromEmail)
	fmt.Println("Please enter your password")
	fmt.Scanln(&password)
	fmt.Println("Please enter a file (this will read the file and send it as plaintext in an email)")
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
	time.Sleep(1 * time.Second)
	fmt.Println("Email sent succesfully,", user.Username)
	time.Sleep(2 * time.Second)
	defer continueTool(user, db)
}
