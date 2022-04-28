package main

import (
	"fmt"
	"net"
)

func main() {

}
func selectTool() {
	var user User
	fmt.Println("Please choose an option:")
	fmt.Println("0. Send mail to Admin (Request higher permission level)")
	fmt.Println("1. Lookup DNS")
	fmt.Println("2. John the Ripper (password check)")
	fmt.Println("3. Extended Lookup DNS")
	var option int
	fmt.Scanln(&option)
	if option == 0 {
		lookupHost(user)
	} else if option == 1 {
		fmt.Println("Coming soon")
	} else if option == 2 {
		fmt.Println("Coming soon")
	} else if option == 3 {
		fmt.Println("Coming soon")
	} else {
		fmt.Println("Invalid option, please try again")
		selectTool()
	}
}
func lookupHost(user User) {
	if user.Permissionlevel >= 1 {
		fmt.Println("Can't do this", user.Username)
	} else {
		fmt.Println("Lookup DNS")
		fmt.Println("Enter a host")
		var host string
		fmt.Scan(&host)
		fmt.Println(net.LookupNS(host))
	}
}
