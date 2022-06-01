package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"time"
)

func submitHelpdesk(user User, db *sql.DB) {
	fmt.Println("What is the problem?")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	fmt.Println("How can we contact you? In case of a serious problem, we would like to keep you updated on our progress")
	var contact string
	fmt.Scanln(&contact)
	time.Sleep(2 * time.Second)
	databaseConn(db)
	time.Sleep(2 * time.Second)
	helpdeskRequest(user, db, line, contact)
	time.Sleep(2 * time.Second)
	fmt.Println("Sent the request")
	extendedToolSelect(user, db)
}
