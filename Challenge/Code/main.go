package main

import (
	"fmt"
	"time"
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

	// Enter option
	var option int
	fmt.Scanln(&option)

	// Database connection using ENV variables
	database := dbConnection()

	// If option
	if option == 1 {
		login(database)

	} else if option == 2 {
		var name string
		var password string
		fmt.Println("Please enter a username")
		fmt.Scanln(&name)
		fmt.Println("Please enter your password")
		fmt.Scanln(&password)
		level := 0
		createUser(name, password, level, database)
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
