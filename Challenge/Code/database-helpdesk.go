package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"
)

func databaseConn(db *sql.DB) {
	query := "CREATE TABLE `helpdesk` (`ID` int(11) NOT NULL AUTO_INCREMENT, `Severity` int(11) NOT NULL, `Message` varchar(2000) NOT NULL, `Username` varchar(2000) NOT NULL, `AssignedTo` varchar(100) NOT NULL, `UserID` int(10) NOT NULL, `Contact` varchar(75) NOT NULL, `UserPermission` int(11) NOT NULL, PRIMARY KEY (`ID`)) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4"
	db.Query(query)
}

func retrieveAllHDRequests(db *sql.DB) {
	query := "SELECT * FROM helpdesk"
	rows, err := db.Query(query)
	checkError(err)
	var id uint8
	var problemValue int
	var problemMessage string
	var username string
	var assignedTo string
	var userID int
	var contactUser string
	var permission int
	for rows.Next() {
		err = rows.Scan(&id, &problemValue, &problemMessage, &username, &assignedTo, &userID, &contactUser, &permission)
		checkError(err)
		fmt.Println("ID:", id, "Severity (Level):", problemValue, "Message:", problemMessage, "Username:", username, "Assigned to:", assignedTo, "User ID:", userID, "Contact:", contactUser)
		fmt.Println("Permission:", permission)
		time.Sleep(2 * time.Second)
	}
	defer rows.Close()
}

func helpdeskRequest(user User, db *sql.DB, problem string, contact string) {
	problemValue := 0
	insertExec := "INSERT INTO helpdesk (Severity, Message, Username, AssignedTo, UserID, Contact, UserPermission) VALUES (?, ?, ?, ?, ?, ? ?)"
	_, erro := db.Exec(insertExec, problemValue, problem, user.Username, "", user.ID, contact, user.Permissionlevel)
	checkError(erro)
}

func editHelpdeskRequest(user User, db *sql.DB) {
	fmt.Println("Please enter an ID")
	var id int
	fmt.Scanln(&id)
	selectStatement := "SELECT * FROM helpdesk WHERE ID = ?"
	_, erro := db.Exec(selectStatement, id)
	checkError(erro)

	fmt.Println("What do you want to edit?")
	fmt.Println("Possible answers:")
	fmt.Println("1. Severity Level (Requires lv 13)")
	fmt.Println("2. Assigned To (Requires lv 14)")
	fmt.Println("3. Contact (Requires lv 12)")
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter an option: ")

	scanner.Scan()

	optionString := scanner.Text()
	edit, err := strconv.Atoi(optionString)
	checkError(err)

	if edit == 1 && user.Permissionlevel >= 13 {
		fmt.Println("Please enter the new level")
		var value int
		fmt.Scanln(&value)
		exec := "UPDATE helpdesk SET Severity = ? WHERE ID = ?"
		db.Exec(exec, value, id)
	} else if edit == 2 && user.Permissionlevel >= 14 {
		fmt.Println("Please enter the new value")
		var value string
		fmt.Scanln(&value)
		exec := "UPDATE helpdesk SET AssignedTo = ? WHERE ID = ?"
		db.Exec(exec, value, id)
	} else if edit == 3 && user.Permissionlevel >= 12 {
		fmt.Println("Please enter the new value")
		var value string
		fmt.Scanln(&value)
		exec := "UPDATE helpdesk SET Contact = ? WHERE ID = ?"
		db.Exec(exec, value, id)
	} else {
		falseOptionFunc(user, db)
	}
}

func removeRequest(user User, db *sql.DB) {
	if user.ID != 14 {
		fmt.Println("Do you have permission, if so who gave you permission?")
		fmt.Println("Please enter a name")
		var username string
		fmt.Scanln(&username)

		if username == "Martijn" {
			fmt.Println("Please enter an ID")
			var id int
			fmt.Scanln(&id)
			exec := "DELETE FROM helpdesk WHERE ID = ?"
			_, erro := db.Exec(exec, id)
			checkError(erro)
			time.Sleep(2 * time.Second)
			fmt.Println("Success")
			helpdeskSelectTool(user, db)
		} else {
			fmt.Println("No permission, closing the application after 20 seconds")
			time.Sleep(20 * time.Second)
		}
	} else {
		fmt.Println("Please enter an ID")
		var id int
		fmt.Scanln(&id)
		exec := "DELETE FROM helpdesk WHERE ID = ?"
		_, erro := db.Exec(exec, id)
		checkError(erro)
		time.Sleep(2 * time.Second)
		fmt.Println("Success")
		helpdeskSelectTool(user, db)
	}
}
