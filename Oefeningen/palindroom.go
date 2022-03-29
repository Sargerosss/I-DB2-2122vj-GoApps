package main

import (
	"fmt"
)

func main() {
	fmt.Println("Palindrome")
	fmt.Println("Please enter a palindrome")
	var palindrome string
	fmt.Scanln(&palindrome)
	isPalindrome(palindrome)
}

func isPalindrome(str string) bool {
	reversedStr := ""
	for i := len(str) - 1; i >= 0; i-- {
		reversedStr += string(str[i])
	}
	for i := range str {
		if str[i] != reversedStr[i] {
			fmt.Println("False")
			return false
		}
	}
	fmt.Println("Correct")
	return true
}
