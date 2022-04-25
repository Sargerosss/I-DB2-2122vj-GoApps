package main

import "fmt"

func main() {

}

func executeFunc() {
	dbConn()

	var (
		name   string
		passwd string
		level  int
	)
	fmt.Scan(&name)
	fmt.Scan(&passwd)
	fmt.Scan(&level)
	createUser(name, passwd, level)
}
