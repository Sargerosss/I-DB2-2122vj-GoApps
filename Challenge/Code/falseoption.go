package main

import (
	"database/sql"
	"fmt"
	"time"
)

var falseOption int

func falseOptionFunc(user User, db *sql.DB) {
	fmt.Println("-----------------------")
	fmt.Println("Invalid option or invalid permission level. Please try again,", user.Username)
	time.Sleep(2 * time.Second)

	if falseOption == 2 {
		fmt.Println("Closing application, too many wrong inputs")
		fmt.Println("-----------------------")
		time.Sleep(30 * time.Second)
	} else {
		falseOption++
		fmt.Println("Please be careful, you have", falseOption, "invalid options given/tried.")
		time.Sleep(10 * time.Second)
		fmt.Println("-----------------------")
		selectTool(user, db)
	}
}
