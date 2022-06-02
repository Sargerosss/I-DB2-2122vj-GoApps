package main

import (
	"fmt"
	"syscall"
	"time"

	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	cybertool()
}

func cybertool() {
	// Print basics
	fmt.Println("-----------------------")
	fmt.Println("Donkey's Cybertool")
	fmt.Println("This tool contains a lot of features")
	fmt.Println("Please login or create an account")
	fmt.Println("Do you want to login (1) or create an account (2)?")
	fmt.Println("Or (3) Close the application")
	fmt.Println("-----------------------")
	fmt.Println("Please enter 1, 2 or 3")
	time.Sleep(2 * time.Second)
	// Enter option
	var option int
	fmt.Scanln(&option)
	time.Sleep(2 * time.Second)
	// Database connection using ENV variables
	database := dbConnection()

	// If option
	if option == 1 {
		login(database)

	} else if option == 2 {
		var name string
		fmt.Println("Please enter a username")
		fmt.Println("Example: Martijn#0001")
		fmt.Println("Second example: Josh#9999")
		fmt.Scanln(&name)
		fmt.Println("Please enter your password")
		passwd, err := terminal.ReadPassword(int(syscall.Stdin))
		checkError(err)
		level := 0
		createUser(name, string(passwd), level, database)
		time.Sleep(2 * time.Second)
		defer cybertool()
	} else if option == 3 {
		fmt.Println("Have a good day, closing application...")
		time.Sleep(2 * time.Second)
	} else {
		fmt.Println("Invalid option, restarting program")
		time.Sleep(2 * time.Second)
		defer cybertool()
	}
}
