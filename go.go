package main

import (
	"fmt"
)

func main() {

	var num int
	var num2 int
	var num3 int

	fmt.Println("Enter a Number: ")
	fmt.Scanln(&num)
	numberToWord(num)

	fmt.Println("Enter number 2: ")
	fmt.Scanln(&num2)
	numberToWord(num2)

	fmt.Println("Enter number 3: ")
	fmt.Scanln(&num3)
	numberToWord(num3)
}

func numberToWord(num int) {
	if num <= 10 && num >= 0 {
		if num == 1 {
			fmt.Println("One")
		} else if num == 2 {
			fmt.Println("Two")
		} else if num == 3 {
			fmt.Println("Three")
		} else if num == 4 {
			fmt.Println("Four")
		} else if num == 5 {
			fmt.Println("Five")
		} else if num == 6 {
			fmt.Println("Six")
		} else if num == 7 {
			fmt.Println("Seven")
		} else if num == 8 {
			fmt.Println("Eight")
		} else if num == 9 {
			fmt.Println("Nine")
		} else if num == 10 {
			fmt.Println("Ten")
		} else if num == 0 {
			fmt.Println("Zero")
		}
	} else {
		fmt.Println("Number too big")
	}
}
