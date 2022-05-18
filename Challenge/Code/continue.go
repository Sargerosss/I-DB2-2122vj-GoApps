package main

import (
	"database/sql"
	"fmt"
	"time"
)

func continueTool(user User, db *sql.DB) {
	fmt.Println("Would you like to (1) go back")
	fmt.Println("Or (2) logout (and return to login screen)")
	fmt.Println("Or (3) close application")

	var option int
	fmt.Scanln(&option)

	if option == 1 {
		selectTool(user, db)
	} else if option == 2 {
		cybertool()
	} else if option == 3 {
		fmt.Println("Closing application in 2 seconds")
		time.Sleep(2 * time.Second)
	}
}
