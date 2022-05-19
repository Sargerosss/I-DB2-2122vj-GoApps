package main

import (
	"database/sql"
	"fmt"
	"time"
)

func adminPanel(user User, db *sql.DB) {
	fmt.Println("Admin Panel")
	fmt.Println("Please choose an option")
	fmt.Println("1. Show all users")
	fmt.Println("2. Create User")
	fmt.Println("3. Remove User")
	fmt.Println("4. Edit User")
	fmt.Println("5. Return")
	var option int
	fmt.Scanln(&option)
	if option == 1 {
		fmt.Println("Coming soon")
	} else if option == 2 {
		fmt.Println("Please choose a Username")
		var username string
		var password string
		var level int
		fmt.Scanln(&username)
		fmt.Println("Please choose a Password")
		fmt.Scanln(&password)
		fmt.Println("Please choose a permissionlevel")
		fmt.Scan(&level)
		createUser(username, password, level, db)
		time.Sleep(2 * time.Second)
		continueTool(user, db)
	} else if option == 3 {
		fmt.Println("Coming soon")
	} else if option == 4 {
		fmt.Println("Not an option")
		adminPanel(user, db)
	} else if option == 5 {
		fmt.Println("Sending you back")
		time.Sleep(2 * time.Second)
		selectTool(user, db)
	}
}
