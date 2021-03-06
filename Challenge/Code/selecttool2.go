package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"
)

func text() {
	fmt.Println("3. Submit a Helpdesk Request")
	fmt.Println("4. Check your Helpdesk requests")
	fmt.Println("5. Return to previous options")
	time.Sleep(2 * time.Second)
	fmt.Println("6. User settings")
	fmt.Println("7. Logout")
	fmt.Println("8. Close app")
	fmt.Println("-----------------------")
}
func extendedToolSelect(user User, db *sql.DB) {
	time.Sleep(2 * time.Second)
	fmt.Println("-----------------------")
	fmt.Println("A reminder, your permission level is", user.Permissionlevel)
	fmt.Println("Other options:")
	if user.Permissionlevel < 9 {
		text()
	}
	if user.Permissionlevel == 9 {
		fmt.Println("1. Challenge (Permission 9)")
		text()
	}
	time.Sleep(2 * time.Second)
	if user.Permissionlevel == 10 {
		fmt.Println("1. Challenge (Permission 9)")
		fmt.Println("2. Admin Tool (Permission 10)")
		time.Sleep(2 * time.Second)
		text()

	}
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter an option: ")

	scanner.Scan()

	optionString := scanner.Text()
	option, err := strconv.Atoi(optionString)
	checkError(err)

	if option == 1 && user.Permissionlevel >= 9 {
		challengesTool(user, db)
	} else if option == 2 && user.Permissionlevel == 10 {
		adminPanel(user, db)
	} else if option == 3 {
		submitHelpdesk(user, db)
	} else if option == 4 {
		retrievePersonalHDReq(user, db)
	} else if option == 5 {
		selectTool(user, db)
	} else if option == 6 {
		editSettingsTool(user, db)
	} else if option == 7 {
		cybertool()
	} else if option == 8 {
		fmt.Println("-----------------------")
		fmt.Println("Closing application in 2 seconds")
		time.Sleep(2 * time.Second)
		fmt.Println("-----------------------")

	} else {
		falseOptionFunc(user, db)
	}
}
