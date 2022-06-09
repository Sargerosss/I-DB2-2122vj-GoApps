package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

func dnsTool(user User, db *sql.DB) {
	option := dnsText()
	dnsOption(option, user, db)
}

func dnsText() int {
	fmt.Println("-----------------------")
	fmt.Println("This is the extended version")
	fmt.Println("These are all the options")
	fmt.Println("1. LookupNS")
	fmt.Println("2. LookupAddr")
	fmt.Println("3. LookupPort")
	fmt.Println("4. Dial")
	fmt.Println("5. Return to previous options")
	fmt.Println("-----------------------")
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter an option: ")
	scanner.Scan()
	optionString := scanner.Text()
	option, err := strconv.Atoi(optionString)
	checkError(err)

	return option
}

func dnsOption(option int, user User, db *sql.DB) {
	if option == 1 {
		var domain string
		fmt.Println("Please enter a domain name")
		fmt.Scanln(&domain)
		fmt.Println(net.LookupNS(domain))
		time.Sleep(2 * time.Second)
		continueTool(user, db)
	} else if option == 2 {
		var ip string
		fmt.Println("Please enter an IP address")
		fmt.Scanln(&ip)
		fmt.Println(net.LookupAddr(ip))
		time.Sleep(2 * time.Second)
		continueTool(user, db)
	} else if option == 3 {
		var network string
		var service string
		fmt.Println("Please enter a network")
		fmt.Println("Example: TCP, UDP")
		fmt.Scanln(&network)
		fmt.Println("Please enter the service")
		fmt.Println("Example: SSH, HTTPS, HTTP")
		fmt.Scanln(&service)
		fmt.Println(net.LookupPort(network, service))
		time.Sleep(2 * time.Second)
		continueTool(user, db)
	} else if option == 4 {
		var ip string
		var network string
		fmt.Println("Please enter an IP address or website")
		fmt.Println("Example: golang.org:http")
		fmt.Scanln(&ip)
		fmt.Println("Please enter a network")
		fmt.Println("Example: TCP, UDP")
		fmt.Scanln(&network)

		fmt.Println(net.Dial(network, ip))
		time.Sleep(2 * time.Second)
		continueTool(user, db)
	} else if option == 5 {
		time.Sleep(2 * time.Second)
		selectTool(user, db)
	} else {
		falseOptionFunc(user, db)
	}
}
