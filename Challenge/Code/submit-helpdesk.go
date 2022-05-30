package main

import (
	"database/sql"
	"fmt"
	"time"
)

func submitHelpdesk(user User, db *sql.DB) {
	fmt.Println("What is the problem?")
	var problem string
	fmt.Scanln(&problem)
	fmt.Println("How can we contact you?")
	var contact string
	fmt.Scanln(&contact)
	fmt.Println("One moment, while I write everything down")
	databaseConn(db)
	time.Sleep(2 * time.Second)
	helpdeskRequest(user, db, problem, contact)
}
