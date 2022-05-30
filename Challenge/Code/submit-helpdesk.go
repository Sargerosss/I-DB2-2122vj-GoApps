package main

import (
	"database/sql"
	"fmt"
	"time"
)

func submitHelpdesk(user User, db *sql.DB) {
	fmt.Println("What is the problem?")
	fmt.Println("Please add a - as a space, so example: (Function) doesnt work becomes Function-doesnt-work")
	var problem string
	fmt.Scanln(&problem)
	fmt.Println("How can we contact you? In case of a serious problem, we would like to keep you updated on our progress")
	var contact string
	fmt.Scanln(&contact)
	time.Sleep(2 * time.Second)
	databaseConn(db)
	time.Sleep(2 * time.Second)
	helpdeskRequest(user, db, problem, contact)
	time.Sleep(2 * time.Second)
	fmt.Println("Sent the request")
	extendedToolSelect(user, db)
}
