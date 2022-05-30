package main

import (
	"database/sql"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username        string
	Password        string
	Permissionlevel int
	ID              int
}

func createUser(name string, pwd string, level int, db *sql.DB) {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS `users` (`ID` int(11) NOT NULL AUTO_INCREMENT, `Username` varchar(50) NOT NULL, `Password` varchar(100) NOT NULL, `Permission` int(10) NOT NULL, PRIMARY KEY (`ID`)) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4")

	passwd, errs := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	checkError(errs)

	checkError(err)

	_, erro := db.Exec("INSERT INTO users (Username, Password, Permission) VALUES (?, ?, ?) ", name, string(passwd), level)
	checkError(erro)

}

func retrieveAllUsers(user User, db *sql.DB) {
	rows, err := db.Query("SELECT ID, Username, Permission FROM users")
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

func removeUser(user User, db *sql.DB) {
	fmt.Println("To remove a User, please give their username")
	fmt.Println("Don't know the Username, but do know the ID? Please enter `ID` or `id`")
	var option string
	fmt.Scanln(&option)
	if option == "ID" || option == "id" {
		fmt.Println("Please enter the ID as a number")
		var id int
		fmt.Scanln(&id)
		db.Exec("DELETE FROM users WHERE ID = ?", id)
		time.Sleep(2 * time.Second)
		fmt.Println("Succesfully deleted")
	} else {
		db.Exec("DELETE FROM users WHERE Username = ?", option)
		time.Sleep(2 * time.Second)
		fmt.Println("Succesfully deleted")
	}

}

func editUser(user User, db *sql.DB) {
	fmt.Println("These are the options")
	fmt.Println("1. Change ID")
	fmt.Println("2. Change Username")
	fmt.Println("3. Change Permissionlevel")
	fmt.Println("Please choose an option")
	var option int
	fmt.Scanln(&option)
	if option == 1 {
		fmt.Println("Please give the current ID")
		var currentId int
		fmt.Scanln(&currentId)
		time.Sleep(2 * time.Second)
		fmt.Println("Please give the new ID")
		var newId int
		fmt.Scanln(&newId)
		time.Sleep(2 * time.Second)
		db.Exec("UPDATE users SET ID = ? WHERE ID = ?", newId, currentId)
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
		db.Exec("UPDATE users SET Username = ? WHERE ID = ?", newUsername, id)
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
		db.Exec("UPDATE users SET Permission = ? WHERE ID = ?", newLevel, id)
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

func retrieveUserID(name string, password string, db *sql.DB) uint8 {
	rows, err := db.Query("SELECT ID, Password FROM users WHERE Username = ?", name)
	checkError(err)
	var id uint8
	var pwd string
	for rows.Next() {
		err = rows.Scan(&id, &pwd)
		checkError(err)
		encryptionErr := bcrypt.CompareHashAndPassword([]byte(pwd), []byte(password))

		if encryptionErr == nil {
			fmt.Println("Your ID is:", id)
			fmt.Println("Keep this ID safe, you will need this to login")
			time.Sleep(2 * time.Second)
		} else {
			fmt.Println("Invalid password")
		}
	}
	defer rows.Close()

	return id
}

func leaderboardAdd(user User, db *sql.DB, room string, score int, website string) {
	db.Query("CREATE TABLE IF NOT EXISTS `leaderboard` (`UserID` int(10) NOT NULL, `Username` varchar(100) NOT NULL, `Website` varchar(50) NOT NULL, `Room` varchar(50) NOT NULL, `Score` int(100) NOT NULL) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4")

	_, erro := db.Exec("INSERT INTO leaderboard (UserID, Username, Website, Room, Score) VALUES (?, ?, ?, ?, ?) ", user.ID, user.Username, website, room, score)
	checkError(erro)
}

func retrieveLeaderboard(user User, db *sql.DB) {
	rows, err := db.Query("SELECT Username, Website, Score FROM leaderboard")
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
