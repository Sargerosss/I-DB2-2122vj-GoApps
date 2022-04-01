package main

import (
	"fmt"
	"os"
)

func main() {
	name, password := authenticate()

	bol := checkUser(name, password)

	if bol {
		fmt.Println("Succesfully logged in")
	} else {
		fmt.Println("Failed to login, please try again.")
	}
}

func authenticate() (string, string) {
	args := os.Args
	name := args[1]
	password := args[2]
	return name, password
}

func checkUser(name string, password string) bool {
	var names [4]string
	names[1] = "Martijn"
	names[2] = "Admin"
	names[3] = "User"

	var passwords [4]string
	passwords[1] = "Martijn"
	passwords[2] = "Admin"
	passwords[3] = "User"

	if (name == names[1] && password == passwords[1]) ||
		(name == names[2] && password == passwords[2]) ||
		(name == names[3] && password == passwords[3]) {
		fmt.Println("Welcome", name)
		return true
	} else {
		fmt.Println("Wrong credentials")
		return false
	}

}
