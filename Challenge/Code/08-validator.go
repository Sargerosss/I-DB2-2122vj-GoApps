package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func validateTool(user User, db *sql.DB) {
	validateText()
	option := validateOption()
	selectOption(option, user, db)
}
func validateText() {
	fmt.Println("-----------------------")
	fmt.Println("This is a validator for data, it can validate things like IP addresses, locations, geocode and decode")
	fmt.Println("Please choose an option")
	fmt.Println("1. Domain")
	fmt.Println("2. IP Addresses")
	fmt.Println("3. Company Email")
	fmt.Println("4. Names")
	fmt.Println("5. Return to options")
	fmt.Println("-----------------------")
}

func validateOption() int {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter an option: ")

	scanner.Scan()

	optionString := scanner.Text()
	option, err := strconv.Atoi(optionString)
	checkError(err)
	fmt.Println("-----------------------")
	return option
}

func selectOption(option int, user User, db *sql.DB) {
	if option == 1 {
		fmt.Println("Domain")
		fmt.Println("Check Domain Name")
		time.Sleep(2 * time.Second)
		domainNameCheck()
		time.Sleep(2 * time.Second)
		continueTool(user, db)
	} else if option == 2 {
		fmt.Println("IP Addresses")
		fmt.Println("IP Address Intelligence")
		ipAddressIntelligence()
		time.Sleep(2 * time.Second)
		continueTool(user, db)
	} else if option == 3 {
		fmt.Println("Company Email")
		companyEmail()
		time.Sleep(2 * time.Second)
		continueTool(user, db)
	} else if option == 4 {
		fmt.Println("Names")
		fmt.Println("1. Validate a full name")
		validateFullName()
		time.Sleep(2 * time.Second)
		continueTool(user, db)
	} else if option == 5 {
		time.Sleep(2 * time.Second)
		selectTool(user, db)
	} else {
		time.Sleep(2 * time.Second)
		falseOptionFunc(user, db)
	}
}

func domainNameCheck() {
	fmt.Println("Please enter a website (example: google.com)")
	var domain string
	fmt.Scanln(&domain)
	url := "https://api.cloudmersive.com/validate/domain/check"
	method := "POST"

	payload := strings.NewReader(domain)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	checkError(err)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Apikey", os.Getenv("API_KEY"))

	res, err := client.Do(req)
	checkError(err)
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	checkError(err)
	fmt.Println(string(body))
}

func ipAddressIntelligence() {
	var ip string
	fmt.Println("Please enter an IP address")
	time.Sleep(2 * time.Second)
	fmt.Scanln(&ip)

	url := "https://api.cloudmersive.com/validate/ip/intelligence"
	method := "POST"

	payload := strings.NewReader(ip)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	checkError(err)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Apikey", os.Getenv("API_KEY"))

	res, err := client.Do(req)
	checkError(err)
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	checkError(err)
	fmt.Println(string(body))
}

func validateFullName() {
	var fileName string
	fmt.Println("Please enter a JSON file name")
	fmt.Scanln(&fileName)

	url := "https://api.cloudmersive.com/validate/name/full-name"
	method := "POST"
	jsonFile, err := os.Open(fileName)

	defer jsonFile.Close()

	bytevalue, _ := ioutil.ReadAll(jsonFile)
	payload := strings.NewReader(string(bytevalue))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	checkError(err)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Apikey", os.Getenv("API_KEY"))

	res, err := client.Do(req)
	checkError(err)
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	checkError(err)
	fmt.Println(string(body))
}

func companyEmail() {
	fmt.Println("Check a company email and get information")
	fmt.Println("Please enter a JSON file")
	var jsonFile string
	fmt.Scanln(&jsonFile)

	openedFile, err := os.Open(jsonFile)
	checkError(err)
	defer openedFile.Close()
	url := "https://api.cloudmersive.com/validate/lead-enrichment/lead/email/company-information"
	method := "POST"
	bytevalue, _ := ioutil.ReadAll(openedFile)
	payload := strings.NewReader(string(bytevalue))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	checkError(err)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Apikey", os.Getenv("API_KEY"))

	res, err := client.Do(req)
	checkError(err)
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	checkError(err)
	fmt.Println(string(body))
}
