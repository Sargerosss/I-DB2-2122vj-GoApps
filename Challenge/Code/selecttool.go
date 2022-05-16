package main

import (
	"database/sql"
	"fmt"
	"time"
)

func selectTool(user User, db *sql.DB) {
	fmt.Println("Hello", user.Username)
	fmt.Println("This is your permission level, you can only choose tools that match your level and lower.")
	fmt.Println("Example: permission level is 5, you can only do 1-5 and not 6")
	fmt.Println("Please choose an option:")
	fmt.Println("0. Send mail to Admin (Request higher permission level)")
	fmt.Println("1. Lookup DNS")
	fmt.Println("2. John the Ripper (password check)")
	fmt.Println("3. Extended Lookup DNS")
	fmt.Println("10. Admin Panel")
	var option int
	fmt.Scanln(&option)
	if option == 0 && user.Permissionlevel >= 0 {
		mailer(user, db)
	} else if option == 1 && user.Permissionlevel >= 1 {
		lookupDNS(user, db)
	} else if option == 2 && user.Permissionlevel >= 2 {
		var password string
		fmt.Println("Please choose a password")
		fmt.Scanln(&password)
		johnRipper(user, password, db)
	} else if option == 3 && user.Permissionlevel >= 3 {
		fmt.Println("Coming soon")
	} else if option == 4 && user.Permissionlevel >= 4 {
		portsniffer(user, db)
	} else if option == 10 && user.Permissionlevel == 10 {
		adminPanel(user, db)
	} else {
		fmt.Println("Invalid option, please try again", user.Username)
		time.Sleep(2 * time.Second)
		selectTool(user, db)
	}
}
