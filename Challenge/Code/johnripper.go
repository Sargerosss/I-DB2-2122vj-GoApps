package main

import (
	"fmt"

	passwordvalidator "github.com/wagslane/go-password-validator"
)

func johnRipper(user User, password string) {
	fmt.Println("This is John the Ripper, a password checker")
	fmt.Println("Please wait", user.Username, "while I check your password")
	entropy := passwordvalidator.GetEntropy(password)

	err := passwordvalidator.Validate(password, entropy)
	checkError(err)

}
