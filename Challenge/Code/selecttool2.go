package main

import (
	"database/sql"
	"fmt"
	"time"
)

func extendedToolSelect(user User, db *sql.DB) {
	fmt.Println("Other options:")
	fmt.Println("A reminder, your permission level is", user.Permissionlevel)
	fmt.Println("1. Challenge (Permission 9)")
	fmt.Println("2. Admin Tool (Permission 10)")
	fmt.Println("3. Submit a Helpdesk Request")
	fmt.Println("4. Return to previous options")
	fmt.Println("5. Logout")
	fmt.Println("6. Close app")

	var option int
	fmt.Scanln(&option)

	if option == 1 && user.Permissionlevel >= 9 {
		challengesTool(user, db)
	} else if option == 2 && user.Permissionlevel == 10 {
		adminPanel(user, db)
	} else if option == 3 {
		submitHelpdesk(user, db)
	} else if option == 4 {
		selectTool(user, db)
	} else if option == 5 {
		cybertool()
	} else if option == 6 {
		fmt.Println("Closing application in 2 seconds")
		time.Sleep(2 * time.Second)

	} else {
		falseOptionFunc(user, db)
	}
}
