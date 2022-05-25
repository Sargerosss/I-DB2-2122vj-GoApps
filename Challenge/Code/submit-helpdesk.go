package main

import (
	"database/sql"
	"fmt"
)

func submitHelpdesk(user User, db *sql.DB) {
	fmt.Println("What is the problem?")
	var problem string
	fmt.Scanln(&problem)
	fmt.Println("One moment, while I write everything down")
	helpdeskRequest(user, db, problem)
}
