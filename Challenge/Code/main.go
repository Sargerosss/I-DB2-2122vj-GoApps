package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	fmt.Println("-----------------------")
	fmt.Println("This tool contains a lot of features")
	fmt.Println("Please login or create an account")
	fmt.Println("1. Login")
	fmt.Println("2. Create an account")
	fmt.Println("3. Close application")
	fmt.Println("4. Forgot my password")
	fmt.Println("-----------------------")
	// Enter option
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter an option: ")

	scanner.Scan()

	optionString := scanner.Text()
	option, err := strconv.Atoi(optionString)
	checkError(err)
	time.Sleep(2 * time.Second)
	// Database connection using ENV variables
	database := dbConnection()

	// If option
	if option == 1 {
		login(database)

	} else if option == 2 {
		fmt.Println("-----------------------")
		fmt.Println("Create an account - Username")
		fmt.Println("Example: Martijn#5424")
		fmt.Println("Second example: Josh#1236")
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Please enter a username: ")
		scanner.Scan()
		name := scanner.Text()
		time.Sleep(2 * time.Second)
		fmt.Print("Please enter your password: ")
		passwd, err := terminal.ReadPassword(int(syscall.Stdin))
		fmt.Println()
		checkError(err)
		level := 0
		fmt.Println()
		createUser(name, string(passwd), level, database)
		time.Sleep(2 * time.Second)
		defer cybertool()
	} else if option == 3 {
		fmt.Println("Have a good day, closing application...")
		time.Sleep(2 * time.Second)
	} else if option == 4 {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Please enter a username: ")
		scanner.Scan()
		name := scanner.Text()
		fmt.Println()
		fmt.Print("Please enter a new password: ")
		passwd, err := terminal.ReadPassword(int(syscall.Stdin))
		fmt.Println()
		checkError(err)
		db := dbConnection()
		passwordForget(name, string(passwd), db)
	} else {
		fmt.Println("Invalid option, restarting program")
		time.Sleep(2 * time.Second)
		defer cybertool()
	}
}
