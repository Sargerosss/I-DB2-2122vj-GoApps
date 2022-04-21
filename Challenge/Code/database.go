package main

import (
	"database/sql"
	"fmt"
	"log"
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

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func readAllUsers() ([]User, error) {
	var users []User

	rows, err := db.Query("SELECT * FROM Users")
	if err != nil {
		return nil, fmt.Errorf("All users: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Username, &user.Permissionlevel); err != nil {
			return nil, fmt.Errorf("All users: %v", err)
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("All users: %v", err)
	}
	return users, nil
}

func login() int {
	var user User
	user.Username = os.Args[1]
	password, err := bcrypt.GenerateFromPassword([]byte(os.Args[2]), bcrypt.MinCost)
	user.Password = string(password)
	rows, err := db.Query("SELECT Permission FROM Users WHERE Username = ? AND Password = ?", user.Username, user.Password)

	for rows.Next() {
		err = rows.Scan(user.Permissionlevel)
		if err != nil {
			panic(err.Error())
		}
		defer rows.Close()
		user.Permissionlevel, err = fmt.Println(user.Permissionlevel)
		return user.Permissionlevel
	}
	return user.Permissionlevel
}

func createUser(name string, pwd string, level int) {
	passwd, errs := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	rows, err := db.Query("INSERT INTO Users (Username, Password, Permission) VALUES (?, ?, ?) ", name, passwd, level)
	defer rows.Close()

	checkError(err)
	checkError(errs)
}
