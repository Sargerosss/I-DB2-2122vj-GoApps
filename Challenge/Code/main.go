package main

import (
	"bufio"
	"fmt"
	"os"
	"syscall"
	"time"

	"github.com/ttacon/chalk"
	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	cybertool()
}

func cybertool() {
	// Print basics
	fmt.Println("-----------------------")
	fmt.Println(chalk.Bold.TextStyle("Donkey's Cybertool"))
	fmt.Println("This tool contains a lot of features")
	fmt.Println("Please login or create an account")
	fmt.Println(chalk.Bold.TextStyle("1. Login"))
	fmt.Println(chalk.Bold.TextStyle("2. Create an account"))
	fmt.Println(chalk.Bold.TextStyle("3. Close application"))
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
		fmt.Println("-----------------------")
		fmt.Println(chalk.Underline.TextStyle("Create an account - Username"))
		fmt.Println(chalk.Italic.TextStyle("Example: Martijn#0001"))
		fmt.Println(chalk.Italic.TextStyle("Second example: Josh#9999"))
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print(chalk.Bold.TextStyle("Please enter a username: "))
		scanner.Scan()
		name := scanner.Text()
		time.Sleep(2 * time.Second)
		fmt.Print(chalk.Bold.TextStyle("Please enter your password: "))
		passwd, err := terminal.ReadPassword(int(syscall.Stdin))
		fmt.Println()
		checkError(err)
		level := 0
		fmt.Println()
		createUser(name, string(passwd), level, database)
		time.Sleep(2 * time.Second)
		defer cybertool()
	} else if option == "3" {
		fmt.Println(chalk.Italic.TextStyle("Have a good day, closing application..."))
		time.Sleep(2 * time.Second)
	} else {
		fmt.Println(chalk.Red.Color("Invalid option, restarting program"))
		time.Sleep(2 * time.Second)
		defer cybertool()
	}
}
