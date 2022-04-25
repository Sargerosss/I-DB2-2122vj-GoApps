package main

// Standaard package die nodig is
import (
	"fmt"
	"net"
)

// Soort library voor het ondersteunen van het printen van lines
func main() {
	fmt.Println("Domain Name System")

	//lookUpDNS(ip)
	fmt.Println("Choose option")
	fmt.Println("1 or 2")
	var option int
	if option == 1 {
		var host string
		fmt.Println("Enter IP")
		fmt.Scanln(&host)
		lookUpHost(host)
	} else if option == 2 {
		var host string
		fmt.Println("Enter DNS")
		fmt.Scanln(&host)
		lookUpDNS(host)
	}
}

func lookUpDNS(host string) {
	fmt.Println("URL")
	fmt.Println(net.LookupHost(host))
}

func lookUpHost(host string) {
	fmt.Println("Address")
	fmt.Println(net.LookupAddr(host))
}
