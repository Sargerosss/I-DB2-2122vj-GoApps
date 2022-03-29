package main

import (
	"fmt"
	"net/smtp"
	"os"
)

// Check error
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// Syntax
	fmt.Println("Syntax: go run automailer.go <text file> <email from> <email to>")

	// Call function
	autoMailer()
}

func autoMailer() {

	// Commandline arguments
	args := os.Args
	// Index1
	file := args[1]
	// Index2
	to := args[2]
	//Index3
	from := args[3]

	// Sender data
	fromEmail := from
	password := "xopiamlikdnknukg"

	// Receiver data
	toEmail := []string{to}

	// SMTP data
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Email body
	data, errors := os.ReadFile(file)
	check(errors)

	// Convert to []byte
	message := []byte(data)

	// Auth for SENDER
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Send mail
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, fromEmail, toEmail, message)
	// Check error
	if err != nil {
		check(err)
	}
	// Finish email
	fmt.Println("")
	fmt.Println("Email Sent Successfully!")
}
