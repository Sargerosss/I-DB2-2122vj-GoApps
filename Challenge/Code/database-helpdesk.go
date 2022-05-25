package main

import (
	"database/sql"
	"fmt"
	"time"
)

func databaseConn(db *sql.DB) {
	db.Query("CREATE TABLE IF NOT EXISTS `helpdesk` (`ID` int(11) NOT NULL AUTO_INCREMENT, `Problem (Value)` int(11) NOT NULL, `Message` varchar(2000) NOT NULL, `Username` varchar(2000) NOT NULL, `Assigned to` varchar(100) NOT NULL, PRIMARY KEY (`ID`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4")
}

func retrieveAllHDRequests(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM helpdesk")
	checkError(err)
	var id uint8
	var permission int
	for rows.Next() {
		err = rows.Scan(&id, &permission)
		checkError(err)
		fmt.Println("ID:", id, "Username:", "Permission:", permission)
		time.Sleep(2 * time.Second)
	}
	defer rows.Close()
}

func helpdeskRequest(user User, db *sql.DB, problem string) {
	problemValue := 0
	_, erro := db.Exec("INSERT INTO helpdesk (ProblemValue, Message, Username, AssignedTo, UserID) VALUES (?, ?, ?) ", problemValue, problem, user.Username, "", user.ID)
	checkError(erro)
}
