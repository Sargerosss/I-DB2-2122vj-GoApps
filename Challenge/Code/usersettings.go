package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"syscall"
	"time"

	"golang.org/x/crypto/ssh/terminal"
)

func editSettingsText() {
	fmt.Println("-----------------------")
	fmt.Println("Edit user settings")
	fmt.Println("1. Edit Username")
	fmt.Println("2. Edit Password")
	fmt.Println("3. Return to options")
	fmt.Println("4. Log out")
	fmt.Println("-----------------------")
}

func editSettingsTool(user User, db *sql.DB) {
	editSettingsText()
	time.Sleep(2 * time.Second)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter an option: ")

	scanner.Scan()

	optionString := scanner.Text()
	option, err := strconv.Atoi(optionString)
	checkError(err)

	if option == 1 {
		fmt.Println("Please enter a new username")
		fmt.Println("Example: Martijn#0001 // Josh#9999")
		scan := bufio.NewScanner(os.Stdin)
		scan.Scan()
		username := scanner.Text()
		usernameCheck := usernameCheck(username, db)

		if !usernameCheck {
			editUsername(username, user, db)
		} else {
			fmt.Println("This username already exists")
			fmt.Println("Sending you back to tool select")
			selectTool(user, db)
		}
	} else if option == 2 {
		fmt.Print("Please enter a new password: ")
		passwd, err := terminal.ReadPassword(int(syscall.Stdin))
		fmt.Println()
		checkError(err)
		editPassword(string(passwd), user, db)
	} else if option == 3 {
		selectTool(user, db)
	} else if option == 4 {
		cybertool()
	}
}
