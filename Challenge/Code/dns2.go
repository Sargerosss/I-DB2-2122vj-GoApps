package main

import (
	"database/sql"
	"fmt"
)

func dnsTool(user User, db *sql.DB) {
	option := dnsText()
	dnsOption(option)
}

func dnsText() int {
	fmt.Println("This is the extended version")
	fmt.Println("These are all the options")
	fmt.Println("1. LookupNS")
	fmt.Println("2. LookupAddr")
	fmt.Println("3. LookupPort")
	fmt.Println("4. Dial")
	var option int
	fmt.Scanln(&option)

	return option
}

func dnsOption(option int) {
	if option == 1 {

	} else if option == 2 {

	} else if option == 3 {

	} else if option == 4 {

	} else {
		fmt.Println("Invalid option")
	}
}
