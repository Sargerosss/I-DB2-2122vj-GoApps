package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func dbConnection() (db *sql.DB) {

	godotenv.Load()
	cfg := mysql.Config{
		User:                 os.Getenv("DBUSER"),
		Passwd:               os.Getenv("DBPASS"),
		Net:                  "tcp",
		Addr:                 os.Getenv("DBIP"),
		DBName:               os.Getenv("DBNAME"),
		AllowNativePasswords: true,
	}

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	checkError(err)

	return db
}

func login(db *sql.DB) User {
	var username string
	var password string
	fmt.Println("Please enter your username")
	fmt.Scanln(&username)
	fmt.Println("Please enter your password")
	fmt.Scanln(&password)
	rows, err := db.Query("SELECT * FROM users WHERE Username = ?", username)
	checkError(err)
	defer rows.Close()
	for rows.Next() {
		var name string
		var pw string
		var level int
		var id int
		erro := rows.Scan(&id, &name, &pw, &level)
		checkError(erro)
		passwordMatch := passwordCheck(password, pw)

		if passwordMatch {
			currentUser := User{name, pw, level, id}
			selectTool(currentUser, db)
			return currentUser
		} else {
			fmt.Println("Login credentials are incorrect.")
			time.Sleep(3 * time.Second)
			cybertool()
			return User{}
		}
	}
	return User{}
}

func passwordCheck(password string, hashedPassword string) bool {
	encryptionErr := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	if encryptionErr == nil {
		return true
	} else {
		return false
	}
}

// Every password different hash (also double ones are receiving new hash)
func logged(db *sql.DB) User {
	var username string
	var passwrd string

	fmt.Println("Please enter your username")
	fmt.Scanln(&username)
	fmt.Println("Please enter your password")
	fmt.Scanln(&passwrd)

	password, err := bcrypt.GenerateFromPassword([]byte(passwrd), bcrypt.MinCost)
	checkError(err)

	rows, err := db.Query("SELECT Permission, ID, Password FROM users WHERE Username = ? AND Password = ?", username, password)
	checkError(err)
	var permLevel int
	var id int

	defer rows.Close()
	for rows.Next() {
		erro := rows.Scan(&permLevel, &id)
		checkError(erro)

	}

	currentUser := User{username, string(password), permLevel, id}
	fmt.Println(currentUser)
	return currentUser
}
