package main

// Standaard package die nodig is
import (
	"fmt"
	"math/rand"
	"strings"
)

// Soort library voor het ondersteunen van het printen van lines
func main() {
	fmt.Println("Password generator & Password strength checker")
	fmt.Println("1. Password generator")
	fmt.Println("2. Password strength checker")
	fmt.Println("Please choose 1 or 2")
	var option int
	fmt.Scanln(&option)
	if option == 1 {
		var number int
		fmt.Println("Please input the length of the password")
		fmt.Scanln(&number)
		fmt.Println(passwordGen(number))
	} else if option == 2 {
		var password string
		fmt.Println("Please input your password")
		fmt.Scanln(&password)
		passwordCheck(password)
	}
}

func passwordGen(number int) string {
	var alphabet = "abcdefghjklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var numbers = "1234567890"
	var specialChar = ",./';[]=-_+{}?><|`~"
	var characterSet = []rune(specialChar + alphabet + numbers)

	s := make([]rune, number)
	for i := range s {
		s[i] = characterSet[rand.Intn(len(characterSet))]
	}
	return string(s)

}

func passwordCheck(password string) {
	var alphabet = "abcdefghjklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var numbers = "1234567890"
	var specialChar = ",./';[]=-_+{}?><|`~"
	var AZ09 = alphabet + numbers
	var AZSC = alphabet + specialChar
	var allChar = alphabet + numbers + specialChar
	len := 0
	for range password {
		len++
	}

	if len > 8 && strings.ContainsAny(password, allChar) {
		fmt.Println("Password strength: Good")
	} else if len > 5 && len < 8 && strings.ContainsAny(password, AZ09) || strings.ContainsAny(password, AZSC) {
		fmt.Println("Password strength: OK")
	} else if len < 8 && strings.ContainsAny(password, alphabet) || strings.Contains(password, numbers) || strings.Contains(password, specialChar) {
		fmt.Println("Password strength: Weak")
	} else {
		fmt.Println("Password couldn't be checked. Try again")
	}
}
