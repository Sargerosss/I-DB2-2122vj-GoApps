package main

import (
	"fmt"
)

func main() {
	var number int
	fmt.Println("Please enter a number")
	fmt.Scanln(&number)
	alarm(number)
}

func alarm(number int) {
	i := 0
	if number > 0 {
		for i <= number {
			fmt.Println("Alarm", i, "!")
			i += 1
		}
	} else {
		fmt.Println("Try again, invalid input")
		main()
	}
}
