package main

import (
	"database/sql"
	"fmt"
	"time"
)

func extendedToolSelect(user User, db *sql.DB) {
	fmt.Println("Other options:")
	fmt.Println("A reminder, your permission level is", user.Permissionlevel)
	fmt.Println("1. Tool 9 (Permission 9)")
	fmt.Println("2. Admin Tool (Permission 10)")
	fmt.Println("3. Return to previous options")
	fmt.Println("4. Logout")
	fmt.Println("5. Close app")

	var option int
	fmt.Scanln(&option)

	if option == 1 && user.Permissionlevel >= 9 {
		fmt.Println("Coming soon")
	} else if option == 2 && user.Permissionlevel == 10 {
		adminPanel(user, db)
	} else if option == 3 {
		selectTool(user, db)
	} else if option == 4 {
		cybertool()
	} else if option == 5 {
		fmt.Println("Closing application in 2 seconds")
		time.Sleep(2 * time.Second)
	} else {
		fmt.Println("Invalid option or invalid permission level. Please try again,", user.Username)
		time.Sleep(2 * time.Second)

		if falseOption == 2 {
			time.Sleep(30 * time.Second)
			fmt.Println("Closing application, too many wrong inputs")
		} else {
			falseOption++
			fmt.Println("Please be careful, you have", falseOption, "invalid options given/tried.")
			time.Sleep(10 * time.Second)
			extendedToolSelect(user, db)
		}

	}
}
