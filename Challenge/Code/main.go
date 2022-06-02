package main

import (
	"bufio"
	"fmt"
	"os"
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
	fmt.Println("1. Login")
	fmt.Println("2. Create an account")
	fmt.Println("3. Close application")
	fmt.Println("-----------------------")
	// Enter option
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter an option: ")

	scanner.Scan()

	option := scanner.Text()
	time.Sleep(2 * time.Second)
	// Database connection using ENV variables
	database := dbConnection()

	// If option
	if option == "1" {
		login(database)

	} else if option == "2" {
		fmt.Println("Example: Martijn#0001")
		fmt.Println("Second example: Josh#9999")
		fmt.Println("Please enter a username: ")
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Enter a username: ")
		scanner.Scan()
		name := scanner.Text()
		time.Sleep(2 * time.Second)
		fmt.Println("Please enter your password: ")
		passwd, err := terminal.ReadPassword(int(syscall.Stdin))
		checkError(err)
		level := 0
		createUser(name, string(passwd), level, database)
		time.Sleep(2 * time.Second)
		defer cybertool()
	} else if option == "3" {
		fmt.Println("Have a good day, closing application...")
		time.Sleep(2 * time.Second)
	} else {
		fmt.Println("Invalid option, restarting program")
		time.Sleep(2 * time.Second)
		defer cybertool()
	}
}
