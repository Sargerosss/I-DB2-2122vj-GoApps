package main

import "fmt"

// https://api.cloudmersive.com/swagger/index.html?urls.primaryName=Validate%20API

func validateText() {
	fmt.Println("This is a validator for data, it can validate things like IP addresses, locations, geocode and decode")
	fmt.Println("Please choose an option")
	fmt.Println("1. Location")
	fmt.Println("2. IP Addresses")
	fmt.Println("3. Input Text")
	fmt.Println("4. Names")
	fmt.Println("5. Domain")
}

func validateOption() int {
	var option int
	fmt.Scanln(&option)
	return option
}

func selectOption(option int) {
	if option == 1 {
		fmt.Println("Location")
		fmt.Println("1. Validate City")
		fmt.Println("2. Validate Street Address")
		fmt.Println("3. Validate Country")
		var options int
		fmt.Scanln(&option)
		if options == 1 {

		} else if options == 2 {

		} else if options == 3 {

		}
	} else if option == 2 {
		fmt.Println("IP Addresses")
	} else if option == 3 {
		fmt.Println("TextInput")
	} else if option == 4 {
		fmt.Println("Names")
	} else if option == 5 {
		fmt.Println("Domain")
	}
}
