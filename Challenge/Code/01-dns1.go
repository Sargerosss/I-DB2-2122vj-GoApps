package main

import (
	"database/sql"
	"fmt"
	"net"
	"time"
)

func lookupDNS(user User, db *sql.DB) {
	fmt.Println("-----------------------")
	fmt.Println("Lookup DNS")
	fmt.Println("Enter a host")
	var host string
	fmt.Scan(&host)
	fmt.Println(net.LookupNS(host))
	time.Sleep(2 * time.Second)
	fmt.Println("-----------------------")
	continueTool(user, db)

}
