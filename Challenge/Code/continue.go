package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"
)

func continueTool(user User, db *sql.DB) {
	fmt.Println("-----------------------")
	fmt.Println("Would you like to")
	fmt.Println("1. Go back (Option 1 - 8)")
	fmt.Println("2. Go back (Option 9 & other options)")
	fmt.Println("3. Logout (and return to login screen)")
	fmt.Println("4. Close application")
	fmt.Println("-----------------------")

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter an option: ")

	scanner.Scan()

	optionString := scanner.Text()
	option, err := strconv.Atoi(optionString)
	checkError(err)

	if option == 1 {
		time.Sleep(2 * time.Second)
		selectTool(user, db)

	} else if option == 2 {
		time.Sleep(2 * time.Second)
		extendedToolSelect(user, db)
	} else if option == 3 {
		time.Sleep(2 * time.Second)
		cybertool()
	} else if option == 4 {
		fmt.Println("Closing application in 2 seconds")
		time.Sleep(2 * time.Second)
	} else {
		falseOptionFunc(user, db)
	}
}
