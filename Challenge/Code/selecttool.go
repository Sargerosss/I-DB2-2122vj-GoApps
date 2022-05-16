package main

import (
	"fmt"
)

func selectTool(user User) {
	fmt.Println("Please choose an option:")
	fmt.Println("0. Send mail to Admin (Request higher permission level)")
	fmt.Println("1. Lookup DNS")
	fmt.Println("2. John the Ripper (password check)")
	fmt.Println("3. Extended Lookup DNS")
	fmt.Println("10. Admin Panel")
	var option int
	fmt.Scanln(&option)
	if option == 0 && user.Permissionlevel >= 0 {
		mailer(user)
	} else if option == 1 {
		lookupDNS(user)
	} else if option == 2 {
		var password string
		fmt.Println("Please choose a password")
		fmt.Scanln(&password)
		johnRipper(user, password)
	} else if option == 3 {
		fmt.Println("Coming soon")
	} else if option == 10 && user.Permissionlevel == 10 {
		adminPanel(user)
	} else {
		fmt.Println("Invalid option, please try again", user.Username)
		selectTool(user)
	}
}
