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
	erro := passwordvalidator.Validate(password, entropy)

	if erro == nil {
		fmt.Println("You have a strong password (has to be bigger than 60)")
	} else {
		fmt.Println("You can have a stronger password")
	}
	time.Sleep(3 * time.Second)
	defer continueTool(user, db)
}
