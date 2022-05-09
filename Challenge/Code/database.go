package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB

// Column names:
// 1) Username
// 2) Password (Hashed)
// 3) Permission

// DATABASE SETUP
// https://www.digitalocean.com/community/tutorials/how-to-install-linux-apache-mysql-php-lamp-stack-on-ubuntu-20-04
// https://www.digitalocean.com/community/tutorials/how-to-install-and-secure-phpmyadmin-on-ubuntu-20-04

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
		Addr:   os.Getenv("DBIP"),
		DBName: os.Getenv("DBTABLE"),
	}

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	checkError(err)

	fmt.Println("Connected!")
}

func checkError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func login() User {
	var user User
	var username string
	var passwrd string

	fmt.Println("Please enter your username")
	fmt.Scanln(&username)
	fmt.Println("Please enter your password")
	fmt.Scanln(&passwrd)

	password, err := bcrypt.GenerateFromPassword([]byte(passwrd), bcrypt.MinCost)
	rows, err := db.Query("SELECT Permission FROM Users WHERE Username = ? AND Password = ?", username, password)

	for rows.Next() {
		err = rows.Scan(&user.Permissionlevel)
		checkError(err)
		defer rows.Close()
	}
	permlevel, err := fmt.Println("Your permissionlevel:", user.Permissionlevel)
	currentUser := User{username, string(password), permlevel}
	return currentUser
}

func createUser(name string, pwd string, level int) {
	var user User
	if user.Permissionlevel == 10 {
		passwd, errs := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
		rows, err := db.Query("SELECT * FROM Users WHERE Username VALUES (?)", name)
		defer rows.Close()

		if rows.Next() {
			fmt.Println("User already exists")
		} else {
			db.Query("INSERT INTO Users (Username, Password, Permission) VALUES (?, ?, ?) ", name, passwd, level)

		}
		checkError(errs)
		checkError(err)
	} else {
		fmt.Println("You can't do this", user.Username)
	}
}
