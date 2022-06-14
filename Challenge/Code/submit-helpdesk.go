package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"time"
)

func submitHelpdesk(user User, db *sql.DB) {
	fmt.Print("What is the problem?: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	fmt.Println("How can we contact you? In case of a serious problem, we would like to keep you updated on our progress")
	fmt.Println("You agree by receiving emails about your request, and you can always ask us to delete your data.")
	fmt.Println("If you don't agree, please let us know it's not a real address/leave it blank/put in no")
	fmt.Print("Your contact information: ")
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	contact := scanner.Text()
	time.Sleep(2 * time.Second)
	databaseConn(db)
	time.Sleep(2 * time.Second)
	helpdeskRequest(user, db, line, contact)
	time.Sleep(2 * time.Second)
	fmt.Println("-----------------------")
	fmt.Println("Result")
	fmt.Println("Sent the request")
	time.Sleep(1 * time.Second)
	extendedToolSelect(user, db)
}
