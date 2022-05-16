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

func dbConnection() (db *sql.DB) {

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
	return db
}

func login(db *sql.DB) User {
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

func createUser(name string, pwd string, level int, db *sql.DB) {
	dbConnection()
	passwd, errs := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	rows, err := db.Query("SELECT * FROM Users WHERE Username VALUES (?)", name)

	checkError(errs)
	checkError(err)

	defer rows.Close()

	if rows.Next() {
		fmt.Println("User already exists")
	} else {
		db.Query("INSERT INTO Users (Username, Password, Permission) VALUES (?, ?, ?) ", name, passwd, level)

	}

}
