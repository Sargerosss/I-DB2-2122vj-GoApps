package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username        string
	Password        string
	Permissionlevel int
	ID              int
}

func createUser(name string, pwd string, level int, db *sql.DB) []byte {
	quered := "CREATE TABLE IF NOT EXISTS `users` (`ID` int(11) NOT NULL AUTO_INCREMENT, `Username` varchar(50) NOT NULL, `Password` varchar(100) NOT NULL, `Permission` int(10) NOT NULL, PRIMARY KEY (`ID`)) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4"
	_, err := db.Exec(quered)

	passwd, errs := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	checkError(errs)
	usernameCheck := usernameCheck(name, db)

	if usernameCheck {
		fmt.Println("Username already exists")
		fmt.Println("Please start again")
		cybertool()
	} else {
		query := "INSERT INTO users (Username, Password, Permission) VALUES (?, ?, ?)"
		_, erro := db.Exec(query, name, string(passwd), level)
		checkError(erro)
		time.Sleep(2 * time.Second)
		fmt.Println("Account created succesfully")
		fmt.Println("-----------------------")
		time.Sleep(2 * time.Second)
	}
	checkError(err)

	return passwd
}

func retrieveAllUsers(user User, db *sql.DB) {
	query := "SELECT ID, Username, Permission FROM users"
	rows, err := db.Query(query)
	checkError(err)
	var username string
	var id uint8
	var permission int
	for rows.Next() {
		err = rows.Scan(&id, &username, &permission)
		checkError(err)
		fmt.Println("ID:", id, "Username:", username, "Permission:", permission)
		time.Sleep(2 * time.Second)
	}
	defer rows.Close()

}
func usernameCheck(username string, db *sql.DB) bool {
	query := "SELECT Username FROM users WHERE Username = ?"
	rows, err := db.Query(query, username)
	checkError(err)
	defer rows.Close()
	return rows.Next()
}
func removeUser(user User, db *sql.DB) {
	fmt.Println("To remove a User, please give their username")
	fmt.Println("Don't know the Username, but do know the ID? Please enter `ID` or `id`")
	var option string
	fmt.Scanln(&option)
	if option == "ID" || option == "id" {
		fmt.Println("Please enter the ID as a number")
		var id int
		fmt.Scanln(&id)
		query := "DELETE FROM users WHERE ID = ?"
		db.Exec(query, id)
		time.Sleep(2 * time.Second)
		fmt.Println("Succesfully deleted")
	} else {
		query := "DELETE FROM users WHERE Username = ?"
		db.Exec(query, option)
		time.Sleep(2 * time.Second)
		fmt.Println("Succesfully deleted")
	}

}

func editUser(user User, db *sql.DB) {
	fmt.Println("These are the options")
	fmt.Println("1. Change ID")
	fmt.Println("2. Change Username")
	fmt.Println("3. Change Permissionlevel")
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter an option: ")

	scanner.Scan()

	optionString := scanner.Text()
	option, err := strconv.Atoi(optionString)
	checkError(err)
	if option == 1 {
		fmt.Println("Please give the current ID")
		var currentId int
		fmt.Scanln(&currentId)
		time.Sleep(2 * time.Second)
		fmt.Println("Please give the new ID")
		var newId int
		fmt.Scanln(&newId)
		time.Sleep(2 * time.Second)
		query := "UPDATE users SET ID = ? WHERE ID = ?"
		db.Exec(query, newId, currentId)
		time.Sleep(2 * time.Second)
		fmt.Println("Succesfully updated")
		time.Sleep(2 * time.Second)
		continueTool(user, db)
	} else if option == 2 {
		fmt.Println("Please give the ID of the User")
		var id int
		fmt.Scanln(&id)
		time.Sleep(2 * time.Second)
		fmt.Println("Please give the new Username")
		var newUsername string
		fmt.Scanln(&newUsername)
		query := "UPDATE users SET Username = ? WHERE ID = ?"
		db.Exec(query, newUsername, id)
		time.Sleep(2 * time.Second)
		fmt.Println("Succesfully updated")
		time.Sleep(2 * time.Second)
		continueTool(user, db)
	} else if option == 3 {
		fmt.Println("Please give the ID of the User")
		var id int
		fmt.Scanln(&id)
		fmt.Println("Please give the new Permissionlevel")
		var newLevel int
		fmt.Scanln(&newLevel)
		query := "UPDATE users SET Permission = ? WHERE ID = ?"
		db.Exec(query, newLevel, id)
		time.Sleep(2 * time.Second)
		fmt.Println("Succesfully updated")
		time.Sleep(2 * time.Second)
		continueTool(user, db)
	} else {
		fmt.Println("Invalid option. Please try again")
		time.Sleep(5 * time.Second)
		falseOptionFunc(user, db)
		editUser(user, db)
	}
}

func leaderboardAdd(user User, db *sql.DB, room string, score int, website string) {
	query := "CREATE TABLE IF NOT EXISTS `leaderboard` (`UserID` int(10) NOT NULL, `Username` varchar(100) NOT NULL, `Website` varchar(50) NOT NULL, `Room` varchar(50) NOT NULL, `Score` int(100) NOT NULL) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4"
	db.Query(query)
	quered := "INSERT INTO leaderboard (UserID, Username, Website, Room, Score) VALUES (?, ?, ?, ?, ?) "
	_, erro := db.Exec(quered, user.ID, user.Username, website, room, score)
	checkError(erro)
}

func retrieveLeaderboard(user User, db *sql.DB) {
	query := "SELECT Username, Website, Score FROM leaderboard"
	rows, err := db.Query(query)
	checkError(err)
	for rows.Next() {
		var username string
		var score int
		var website string
		err = rows.Scan(&username, &website, &score)
		checkError(err)
		fmt.Println("Username:", username, "Website:", website, "Score:", score)
		time.Sleep(2 * time.Second)
	}
}
