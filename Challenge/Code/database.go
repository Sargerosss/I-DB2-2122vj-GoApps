package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB

// Table names:
// 1) Username
// 2) Password (Hashed)
// 3)

type User struct {
	Username        string
	Password        string
	Permissionlevel int
}

func main() {

}

func dbConn() {

	godotenv.Load()
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "(IP):3306",
		DBName: "users",
	}

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected!")

}

func addUser(user User) (string, error) {
	result, err := db.Exec("INSERT INTO users (username, password, permission) VALUES (?, ?, ?)", user.Username, user.Password, user.Permissionlevel)
	if err != nil {
		return "", fmt.Errorf("addAlbum: %v", err)
	}
	return user.Username, nil
}

func removeUser() {}

func editUser() {}

func readAllUsers() {}

func login(user User) {
	user.Username = os.Args[1]
	user.Password = os.Args[2]

	rows, err := db.Query("SELECT Permission FROM users WHERE username = ?", user.Username)

	for rows.Next() {
		err = rows.Scan(user.Permissionlevel)
		if err != nil {
			panic(err.Error())
		}
		user.Permissionlevel, err = fmt.Println(user.Permissionlevel)
	}
}
