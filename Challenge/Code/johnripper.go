package main

import (
	"database/sql"
	"fmt"
	"time"

	passwordvalidator "github.com/wagslane/go-password-validator"
)

func johnRipper(user User, password string, db *sql.DB) {
	fmt.Println("This is John the Ripper, a password checker")
	fmt.Println("Please wait", user.Username, "while I check your password")
	entropy := passwordvalidator.GetEntropy(password)

	err := passwordvalidator.Validate(password, entropy)
	checkError(err)
	time.Sleep(2 * time.Second)
	continueTool(user, db)
}
