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
	problem := scanner.Text()
	fmt.Println("How can we contact you? In case of a serious problem, we would like to keep you updated on our progress")
	fmt.Println("You agree by receiving emails/phone calls about your request, and you can always ask us to delete your personal data.")
	fmt.Println("If you don't agree, please let us know it's not a real email address/leave it blank/put in no/not a real phone number")
	fmt.Print("Your contact information: ")
	var contact string
	fmt.Scanln(&contact)
	time.Sleep(2 * time.Second)
	databaseConn(db)
	time.Sleep(2 * time.Second)
	helpdeskRequest(user, db, problem, contact)
	time.Sleep(2 * time.Second)
	fmt.Println("-----------------------")
	fmt.Println("Sent the request")
	fmt.Println("Sending you back now")
	time.Sleep(1 * time.Second)
	extendedToolSelect(user, db)
}
