package main

import (
	"database/sql"
	"fmt"
)

func adminPanel(user User, db *sql.DB) {
	fmt.Println("Admin Panel")
	fmt.Println("Please choose an option")
	fmt.Println("1. Create User")
	fmt.Println("2. Remove User")
	fmt.Println("3. Edit User")
	var option int
	fmt.Scanln(&option)
	if option == 1 {
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
	}
}
