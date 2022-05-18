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

	//fmt.Println("Connected!")
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

	//fmt.Println("Your permissionlevel:", user.Permissionlevel)
	currentUser := User{username, string(password), user.Permissionlevel}
	return currentUser
}

func createUser(name string, pwd string, level int, db *sql.DB) {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS`users` (`Username` varchar(50) NOT NULL, `Password` varchar(100) NOT NULL, `Permission` int(10) NOT NULL) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;")

	passwd, errs := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	checkError(errs)
	rows, err := db.Query("SELECT * FROM users WHERE Username VALUES (?)", name)
	defer rows.Close()
	checkError(err)

	if rows.Next() {
		fmt.Println("User already exists")
	} else {
		_, erro := db.Exec("INSERT INTO users (Username, Password, Permission) VALUES (?, ?, ?) ", name, string(passwd), level)
		checkError(erro)
	}

}
