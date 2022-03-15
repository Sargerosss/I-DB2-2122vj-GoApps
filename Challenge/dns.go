package main

// Standaard package die nodig is
import (
	"fmt"
	"net"
)

// Soort library voor het ondersteunen van het printen van lines
func main() {
	fmt.Println("Domain Name System")
	var ip string
	fmt.Scanln(&ip)
	//lookUpDNS(ip)
	lookUpHost(ip)
}

func lookUpDNS(host string) {
	fmt.Println("URL")
	fmt.Println(net.LookupHost(host))
}

func lookUpHost(host string) {
	fmt.Println("Address")
	fmt.Println(net.LookupAddr(host))
}
