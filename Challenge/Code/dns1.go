package main

import (
	"database/sql"
	"fmt"
	"net"
	"time"
)

func lookupDNS(user User, db *sql.DB) {
	if user.Permissionlevel <= 1 {
		fmt.Println("Can't do this", user.Username)
		time.Sleep(2 * time.Second)
		fmt.Println("Sending you back to start")
		time.Sleep(time.Second)
		selectTool(user, db)
	} else {
		fmt.Println("Lookup DNS")
		fmt.Println("Enter a host")
		var host string
		fmt.Scan(&host)
		fmt.Println(net.LookupNS(host))
		time.Sleep(2 * time.Second)
		continueTool(user, db)
	}
}
