package main

import (
	"database/sql"
	"fmt"
	"time"

	passwordvalidator "github.com/wagslane/go-password-validator"
)

func johnRipper(user User, password string, db *sql.DB) {
	fmt.Println("This is John the Ripper, a password checker")
	fmt.Println("Please wait,", user.Username, ", while I check your password")
	entropy := 60.0
	fmt.Println("The entropy used: ", entropy)
	fmt.Println("Your entropy:", passwordvalidator.GetEntropy(password))
	time.Sleep(1 * time.Second)
	fmt.Println(passwordvalidator.Validate(password, entropy))
	time.Sleep(3 * time.Second)
	defer continueTool(user, db)
}
