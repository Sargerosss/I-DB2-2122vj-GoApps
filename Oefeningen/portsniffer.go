package main

import (
	"fmt"
	"net"
	"strconv"
)

func main() {
	var ip string
	fmt.Println("Please enter an IP address")
	fmt.Println("Please add : after entering IP Address")
	fmt.Scanln(&ip)
	for i := 1; i < 100; i++ {
		port := strconv.FormatInt(int64(i), 10)
		conn, err := net.Dial("tcp", ip+port)
		if err == nil {
			fmt.Println("Port", i, "open")
			conn.Close()
		} else {
			fmt.Println("Port", i, "closed")
		}
	}
}
