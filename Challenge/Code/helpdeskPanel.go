package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"
)

// 11: Helpdesk Employee
// 12: Sr. Helpdesk Employee
// 13: Lead Helpdesk
// 14: Manager

func startTextHD(user User, db *sql.DB) {
	fmt.Println("-----------------------")
	fmt.Println("Hello", user.Username)
	fmt.Println("Your permission level:", user.Permissionlevel)
	fmt.Println("Your ID:", user.ID)
	fmt.Println("This is the Helpdesk Panel")
	time.Sleep(2 * time.Second)
	fmt.Println("Please choose an option")
}
func endTextHD() {
	fmt.Println("4. Log out")
	fmt.Println("5. Close Application")
	fmt.Println("-----------------------")
}
func helpdeskSelectTool(user User, db *sql.DB) {
	if user.Permissionlevel >= 11 {
		startTextHD(user, db)
		fmt.Println("1. Retrieve all helpdesk requests (11)")
	}

	if user.Permissionlevel >= 12 {
		fmt.Println("2. Edit helpdesk requests (12/13)")

	}

	if user.Permissionlevel >= 13 {
		fmt.Println("3. Remove request (13/14)")

	}

	endTextHD()
	time.Sleep(2 * time.Second)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter an option: ")

	scanner.Scan()

	optionString := scanner.Text()
	option, err := strconv.Atoi(optionString)
	checkError(err)
	helpdeskOptions(user, db, option)
}

func helpdeskOptions(user User, db *sql.DB, option int) {
	if option == 1 && user.Permissionlevel >= 11 {
		databaseConn(db)
		time.Sleep(2 * time.Second)
		retrieveAllHDRequests(db)
		time.Sleep(2 * time.Second)
		helpdeskSelectTool(user, db)
	} else if option == 2 && user.Permissionlevel >= 12 {
		editHelpdeskRequest(user, db)
		time.Sleep(2 * time.Second)
		helpdeskSelectTool(user, db)
	} else if option == 3 && user.Permissionlevel >= 13 {
		removeRequest(user, db)
		time.Sleep(2 * time.Second)
		helpdeskSelectTool(user, db)
	} else if option == 4 {
		cybertool()
	} else if option == 5 {
		fmt.Println("Closing application in 2 seconds")
		time.Sleep(2 * time.Second)
	} else {
		falseOptionFunc(user, db)
	}
}
