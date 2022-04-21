package main

import (
	"database/sql"
)

func main() {
}

func dbConn() bool {
	// after the / is db name
	db, err := sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/test")

	if err != nil {
		panic(err.Error())
		return false
	} else {
		return true
	}
	defer db.Close()

}
