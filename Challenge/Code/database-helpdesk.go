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
	query := "CREATE TABLE IF NOT EXISTS `helpdesk` (`ID` int(11) NOT NULL AUTO_INCREMENT, `Severity` int(11) NOT NULL, `Message` varchar(2000) NOT NULL, `Username` varchar(2000) NOT NULL, `AssignedTo` varchar(100) NOT NULL, `UserID` int(10) NOT NULL, `Contact` varchar(75) NOT NULL, `UserPermission` int(11) NOT NULL, PRIMARY KEY (`ID`)) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4"
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
		fmt.Println("ID:", id, "Severity (Level):", problemValue, "Message:", problemMessage, "Username:", username, "Assigned to:", assignedTo, "User ID:", userID, "Contact:", contactUser, "Permission:", permission)
		time.Sleep(2 * time.Second)
	}
	defer rows.Close()
}

func helpdeskRequest(user User, db *sql.DB, problem string, contact string) {
	problemValue := 0
	contactInfo := contact
	insertExec := "INSERT INTO helpdesk (Severity, Message, Username, AssignedTo, UserID, Contact, UserPermission) VALUES (?, ?, ?, ?, ?, ?, ?)"
	_, erro := db.Exec(insertExec, problemValue, problem, user.Username, "", user.ID, contactInfo, user.Permissionlevel)
	checkError(erro)
}

func editHelpdeskRequest(user User, db *sql.DB) {
	fmt.Println("Please enter an ID")
	var id int
	fmt.Scanln(&id)
	selectStatement := "SELECT * FROM helpdesk WHERE ID = ?"
	_, erro := db.Query(selectStatement, id)
	checkError(erro)

	fmt.Println("What do you want to edit?")
	fmt.Println("Possible answers:")
	fmt.Println("1. Severity Level (Requires lv 13)")
	fmt.Println("2. Assigned To (Requires lv 13)")
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
	} else if edit == 2 && user.Permissionlevel >= 13 {
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
		fmt.Println("Please be sure to have permission, and keep things logged in the file")
		time.Sleep(2 * time.Second)
		remove(user, db)
	} else {
		remove(user, db)
	}
}

func remove(user User, db *sql.DB) {
	fmt.Println("Please enter an ID")
	var id int
	fmt.Scanln(&id)
	exec := "DELETE FROM helpdesk WHERE ID = ?"
	_, erro := db.Exec(exec, id)
	checkError(erro)
	time.Sleep(2 * time.Second)
	fmt.Println("Successfully deleted HD request with ID", id)
	helpdeskSelectTool(user, db)
}

func retrievePersonalHDReq(user User, db *sql.DB) {
	query := "SELECT ID, Message, Contact FROM helpdesk WHERE UserID = ?"
	rows, err := db.Query(query, user.ID)
	checkError(err)
	defer rows.Close()
	for rows.Next() {
		var id int
		var message string
		var contact string
		err := rows.Scan(&id, &message, &contact)
		checkError(err)
		fmt.Println("ID:", id, "Message:", message, "Contact:", contact)
		time.Sleep(1 * time.Second)
	}
	defer continueTool(user, db)
}
