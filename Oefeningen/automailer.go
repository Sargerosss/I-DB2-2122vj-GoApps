package main

// Standaard package die nodig is
import (
	"fmt"
	"net/smtp"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Soort library voor het ondersteunen van het printen van lines
func main() {
	fmt.Println("Automailer")
	fmt.Println("Syntax:")
	fmt.Println("go run automailer.go <text file> <email to> <email from>")
	args := os.Args

	file := args[1]
	mail := args[2]
	from := args[3]

	// Sender data
	fromEmail := from
	password := "xopiamlikdnknukg"

	to := []string{mail}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	data, errors := os.ReadFile(file)

	check(errors)
	message := []byte(data)

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, fromEmail, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent Successfully!")
}
