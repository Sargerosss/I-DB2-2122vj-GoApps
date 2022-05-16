package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

// Column names:
// 1) Username
// 2) Password (Hashed)
// 3) Permission

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
	defer db.Close()
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
	rows, err := db.Query("SELECT Permission FROM users WHERE Username = ? AND Password = ?", username, string(password))

	for rows.Next() {
		err = rows.Scan(&user.Permissionlevel)
		checkError(err)

	}

	defer rows.Close()

	permlevel, err := fmt.Println("Your permissionlevel:", user.Permissionlevel)
	currentUser := User{username, string(password), permlevel}
	return currentUser
}

func createUser(name string, pwd string, level int, db *sql.DB) {
	//dbConnection()
	passwd, errs := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	fmt.Println(passwd)
	rows, err := db.Query("SELECT * FROM users WHERE Username VALUES (?)", name)

	checkError(errs)
	checkError(err)

	defer rows.Close()

	if rows.Next() {
		fmt.Println("User already exists")
	} else {
		db.Query("INSERT INTO users (Username, Password, Permission) VALUES (?, ?, ?) ", name, string(passwd), level)

	}

}
