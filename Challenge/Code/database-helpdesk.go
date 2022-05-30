package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"
)

func databaseConn(db *sql.DB) {
	db.Query("CREATE TABLE `helpdesk` ( `ID` int(11) NOT NULL AUTO_INCREMENT, `ProblemValue` int(11) NOT NULL, `Message` varchar(2000) NOT NULL, `Username` varchar(2000) NOT NULL, `AssignedTo` varchar(100) NOT NULL, `UserID` int(10) NOT NULL, `Contact` varchar(75) NOT NULL, PRIMARY KEY (`ID`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4")
}

func retrieveAllHDRequests(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM helpdesk")
	checkError(err)
	var id uint8
	var problemValue int
	var problemMessage string
	var username string
	var assignedTo string
	var userID int
	var contactUser string
	for rows.Next() {
		err = rows.Scan(&id, &problemValue, &problemMessage, &username, &assignedTo, &userID, &contactUser)
		checkError(err)
		fmt.Println("ID:", id, "Problem (Value):", problemValue, "Message:", problemMessage, "Username:", username, "Assigned to:", assignedTo, "User ID:", userID, "Contact:", contactUser)
		time.Sleep(2 * time.Second)
	}
	defer rows.Close()
}

func helpdeskRequest(user User, db *sql.DB, problem string, contact string) {
	problemValue := 0
	_, erro := db.Exec("INSERT INTO helpdesk (ProblemValue, Message, Username, AssignedTo, UserID, Contact) VALUES (?, ?, ?, ?, ?, ?) ", problemValue, problem, user.Username, "", user.ID, contact)
	checkError(erro)
}

func editHelpdeskRequest(user User, db *sql.DB) {
	fmt.Println("Please enter an ID")
	var id int
	fmt.Scanln(&id)

	_, erro := db.Exec("SELECT * FROM helpdesk WHERE ID = ?", id)
	checkError(erro)

	fmt.Println("What do you want to edit?")
	fmt.Println("Possible answers:")
	fmt.Println("1. Problem Value (Requires lv 13)")
	fmt.Println("2. Assigned To (Requires lv 14)")
	fmt.Println("3. Contact (Requires lv 12)")
	var edit int
	fmt.Scanln(&edit)

	if edit == 1 && user.Permissionlevel >= 13 {
		fmt.Println("Please enter the new value")
		var value int
		fmt.Scanln(&value)
		db.Exec("UPDATE helpdesk SET ProblemValue = ? WHERE ID = ?", value, id)
	} else if edit == 2 && user.Permissionlevel >= 14 {
		fmt.Println("Please enter the new value")
		var value string
		fmt.Scanln(&value)
		db.Exec("UPDATE helpdesk SET AssignedTo = ? WHERE ID = ?", value, id)
	} else if edit == 3 && user.Permissionlevel >= 12 {
		fmt.Println("Please enter the new value")
		var value string
		fmt.Scanln(&value)
		db.Exec("UPDATE helpdesk SET Contact = ? WHERE ID = ?", value, id)
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

		if username == os.Getenv("USERNAME") {
			fmt.Println("Please enter an ID")
			var id int
			fmt.Scanln(&id)

			_, erro := db.Exec("DELETE FROM helpdesk WHERE ID = ?", id)
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

		_, erro := db.Exec("DELETE FROM helpdesk WHERE ID = ?", id)
		checkError(erro)
		time.Sleep(2 * time.Second)
		fmt.Println("Success")
		helpdeskSelectTool(user, db)
	}
}
