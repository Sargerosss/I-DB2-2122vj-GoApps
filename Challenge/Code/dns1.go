package main

import (
	"fmt"
	"net"
)

func lookupDNS(user User) {
	if user.Permissionlevel <= 1 {
		fmt.Println("Can't do this", user.Username)
	} else {
		fmt.Println("Lookup DNS")
		fmt.Println("Enter a host")
		var host string
		fmt.Scan(&host)
		fmt.Println(net.LookupNS(host))
	}
}
