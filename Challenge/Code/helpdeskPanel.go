package main

import (
	"database/sql"
	"fmt"
	"time"
)

// 11: Helpdesk Employee
// 12: Sr. Helpdesk Employee
// 13: Lead Helpdesk
// 14: Manager

func helpdeskSelectTool(user User, db *sql.DB) {
	fmt.Println("Hello", user.Username)
	fmt.Println("Your permission level:", user.Permissionlevel)
	fmt.Println("Your ID:", user.ID)
	fmt.Println("This is the Helpdesk Panel")
	time.Sleep(2 * time.Second)
	fmt.Println("Please choose an option")
	fmt.Println("1. Retrieve all helpdesk requests (11)")
	fmt.Println("2. Edit helpdesk requests (12/13)")
	fmt.Println("3. Remove request (13/14)")
	fmt.Println("4. Log out")
	fmt.Println("5. Close Application")
}

func helpdeskOptions(user User, db *sql.DB, option int) {
	if option == 1 {
		retrieveAllHDRequests(db)
		time.Sleep(2 * time.Second)
		helpdeskSelectTool(user, db)
	} else if option == 2 {
		fmt.Println("")
	}
}
