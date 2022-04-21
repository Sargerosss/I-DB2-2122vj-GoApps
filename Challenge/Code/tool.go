package main

import "fmt"

func main() {

}

func tool1(user User) {
	if user.Permissionlevel <= 1 {
		fmt.Println("Can't do this")
	} else {
		fmt.Println("Lookup Host")
		fmt.Println("Enter IP address")
		var host string
		fmt.Scan(&host)
	}
}
