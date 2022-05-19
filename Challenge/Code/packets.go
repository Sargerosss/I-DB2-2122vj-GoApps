package main

import (
	"database/sql"
	"fmt"
)

// https://github.com/ghedo/go.pkt/tree/master/examples
// https://github.com/ghedo/go.pkt

func packetTool(user User, db *sql.DB) {
	option := packetText()
	packetOption(option, user, db)
}

func packetText() int {
	fmt.Println("This is the packet tool, similar to Wireshark, it has a lot of features around packets")
	fmt.Println("These are the options:")
	fmt.Println("1. Capturing")
	fmt.Println("2. Injection")
	fmt.Println("3. Filtering")
	fmt.Println("4. Encoding")
	fmt.Println("5. Decoding")
	var option int
	fmt.Scanln(&option)
	return option
}

func packetOption(option int, user User, db *sql.DB) {
	if option == 1 {

	} else if option == 2 {

	} else if option == 3 {

	} else if option == 4 {

	} else if option == 5 {

	}

}
