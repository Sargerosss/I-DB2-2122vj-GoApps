package main

// Standaard package die nodig is
import (
	"fmt"
)

// Soort library voor het ondersteunen van het printen van lines
func main() {
	switchNumbers()
}

func oddEven() {
	var number = 50
	for i := 0; i <= number; i++ {
		if i%2 == 0 {
			fmt.Println("Even ", i)
		} else {
			fmt.Println("Odd ", i)
		}
	}
}

func evenOnly() {
	var number = 50
	for i := 0; i <= number; i++ {
		if i%2 == 0 {
			fmt.Println("Even ", i)
		} else {
			i = i + 2
		}

	}
}

func oddOnly() {
	var number = 50
	for i := 0; i <= number; i++ {
		if i%2 == 0 {
			i = i + 2
		} else {
			fmt.Println("Odd ", i)
		}
	}
}

func switchNumbers() {
	var number = 6
	for i := 0; i < number; i++ {
		switch i {
		case 0:
			fmt.Println("Zero")
		case 1:
			fmt.Println("One")
		case 2:
			fmt.Println("Two")
		case 3:
			fmt.Println("Three")
		case 4:
			fmt.Println("Four")
		case 5:
			fmt.Println("Five")

		}
	}
}
