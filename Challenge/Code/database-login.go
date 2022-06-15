package main

import (
	"bufio"
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

func login(db *sql.DB) User {
	fmt.Println("-----------------------")
	fmt.Println("Login - Username")
	fmt.Println("Example: Martijn#0001")
	fmt.Println("Second example: Josh#9999")
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Please enter your username: ")

	scanner.Scan()

	username := scanner.Text()
	fmt.Print("Please enter your password: ")
	passwd, err := terminal.ReadPassword(int(syscall.Stdin))
	fmt.Println()
	query := "SELECT * FROM users WHERE Username = ?"
	rows, err := db.Query(query, username)
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
		//usernameMatch := usernameCheck(name, db)
		if passwordMatch {
			currentUser := User{name, pw, level, id}

			if level < 11 {
				selectTool(currentUser, db)
			} else if level > 11 {
				helpdeskSelectTool(currentUser, db)
			}

			return currentUser
			// || !usernameMatch
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

func usernameCheck(username string, db *sql.DB) bool {
	query := "SELECT Username FROM users WHERE Username = ?"
	rows, err := db.Query(query, username)
	checkError(err)
	defer rows.Close()
	return rows.Next()
}
