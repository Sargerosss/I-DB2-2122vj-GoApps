package main

import (
	"database/sql"
	"fmt"
	"os"
	"syscall"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/ssh/terminal"
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
func usernameCheck(username string, db *sql.DB) bool {
	rows, err := db.Query("SELECT Username FROM users WHERE Username = ?", username)
	checkError(err)
	defer rows.Close()
	for rows.Next() {
		var name string
		rows.Scan(&name)

		if name == username {
			return true
		} else {
			fmt.Println("Invalid username")
			return false
		}
	}
	return false
}
func login(db *sql.DB) User {
	var username string
	var id int
	fmt.Println("Please enter your username")
	fmt.Scanln(&username)
	fmt.Println("Please enter your password")
	passwd, err := terminal.ReadPassword(int(syscall.Stdin))
	fmt.Println("Please enter your ID")
	fmt.Scanln(&id)
	rows, err := db.Query("SELECT * FROM users WHERE Username = ? AND ID = ?", username, id)
	checkError(err)
	defer rows.Close()
	for rows.Next() {
		var name string
		var pw string
		var level int
		var id int
		erro := rows.Scan(&id, &name, &pw, &level)
		checkError(erro)
		passwordMatch := passwordCheck(string(passwd), pw)
		if passwordMatch {
			currentUser := User{name, pw, level, id}

			if level < 11 {
				selectTool(currentUser, db)
			} else if level == 11 {
				helpdeskSelectTool(currentUser, db)
			}

			return currentUser
		} else if !passwordMatch {
			fmt.Println("Password doesn't match")
			time.Sleep(3 * time.Second)
			cybertool()
			return User{}
		} else {
			fmt.Println("Invalid login credentials")
			time.Sleep(2 * time.Second)
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
		fmt.Println("Invalid password")
		return false
	}
}
