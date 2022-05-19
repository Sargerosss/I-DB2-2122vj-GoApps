package main

import (
	"database/sql"
	"fmt"
	"time"
)

var falseOption int

func selectTool(user User, db *sql.DB) {
	fmt.Println("Hello", user.Username)
	fmt.Println("Your permissionlevel is: ", user.Permissionlevel)
	fmt.Println("Your ID is", user.ID)
	fmt.Println("This is your permission level, you can only choose tools that match your level and lower.")
	fmt.Println("Example: permission level is 5, you can only do 1-5 and not 6")
	fmt.Println("Please choose an option:")
	fmt.Println("0. Send mail to Admin (Request higher permission level) (Sends text files as email)")
	fmt.Println("1. Lookup DNS")
	fmt.Println("2. John the Ripper (password check)")
	fmt.Println("3. Extended Lookup DNS")
	fmt.Println("4. Portsniffer")
	fmt.Println("5. Packets (Capture, Inject, Encode & Decode)")
	fmt.Println("6. File Encryption & Decryption")
	fmt.Println("9. More options (Like log out, close app, and more tools)")
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
		dnsTool(user, db)
	} else if option == 4 && user.Permissionlevel >= 4 {
		portsniffer(user, db)
	} else if option == 5 && user.Permissionlevel >= 5 {
		packetTool(user, db)
	} else if option == 6 && user.Permissionlevel >= 6 {
		encryptTool(user, db)
	} else if option == 7 && user.Permissionlevel >= 7 {
		fmt.Println("Coming soon")
	} else if option == 8 && user.Permissionlevel >= 8 {
		fmt.Println("Coming soon")
	} else if option == 9 {
		extendedToolSelect(user, db)
	} else {
		fmt.Println("Invalid option or invalid permission level. Please try again,", user.Username)
		time.Sleep(2 * time.Second)

		if falseOption == 2 {
			fmt.Println("Closing application, too many wrong inputs")
		} else {
			falseOption++
			fmt.Println("Please be careful, you have", falseOption, "invalid options given/tried.")
			time.Sleep(10 * time.Second)
			selectTool(user, db)
		}
	}
}
