package main

import (
	"fmt"
	"time"
)

func main() {
	cybertool()
}

func cybertool() {
	fmt.Println("Donkey's Cybertool")
	fmt.Println("This tool contains a lot of features")
	fmt.Println("Please login or create an account")
	fmt.Println("Do you want to login (1) or create an account (2)?")
	fmt.Println("Or (3) Close the application")
	var option int
	fmt.Scanln(&option)
	db := dbConnection()
	if option == 1 {
		user := login(db)
		defer selectTool(user, db)
	} else if option == 2 {
		var name string
		var password string
		fmt.Println("Please enter a username")
		fmt.Scanln(&name)
		fmt.Println("Please enter your password")
		fmt.Scanln(&password)
		level := 0
		createUser(name, password, level, db)
		cybertool()
	} else if option == 3 {
		fmt.Println("Have a good day, closing application...")
		time.Sleep(2 * time.Second)
	} else {
		fmt.Println("Invalid option, restarting program")
		time.Sleep(2 * time.Second)
		cybertool()
	}
}
