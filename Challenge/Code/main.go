package main

import "fmt"

func main() {

	fmt.Println("Donkey's Cybertool")
	fmt.Println("This tool contains a lot of features")
	fmt.Println("Please login or create an account")
	fmt.Println("Do you want to login (1) or create an account (2)?")
	var option int
	fmt.Scanln(&option)

	if option == 1 {
		dbConn()
		login()
		defer selectTool()
	} else if option == 2 {
		var name string
		var password string
		fmt.Println("Please enter a username")
		fmt.Scanln(&name)
		fmt.Println("Please enter your password")
		fmt.Scanln(&password)
		level := 0
		createUser(name, password, level)
		defer main()
	} else {
		fmt.Println("Have a good day, closing application...")
	}
}
