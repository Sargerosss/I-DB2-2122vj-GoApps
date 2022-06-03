package main

import (
	"database/sql"
	"fmt"
	"net"
	"strconv"
)

func portsniffer(user User, db *sql.DB) {
	var ip string
	var portStart int
	var portEnd int
	fmt.Println("-----------------------")
	fmt.Println("Please enter an IP address")
	fmt.Println("This will take a while")
	fmt.Println("Please add : after entering IP Address")
	fmt.Scanln(&ip)
	fmt.Println("Please enter the port start")
	fmt.Scanln(&portStart)
	fmt.Println("Please enter the port end")
	fmt.Scanln(&portEnd)
	fmt.Println("-----------------------")
	for i := portStart; i <= portEnd; i++ {
		port := strconv.FormatInt(int64(i), 10)
		conn, err := net.Dial("tcp", ip+port)
		if err == nil {
			fmt.Println("Port", i, "open")
			fmt.Println("-----------------------")
			conn.Close()
		} else {
			fmt.Println("Port", i, "closed")
			fmt.Println("-----------------------")
		}
	}

	defer continueTool(user, db)
}
