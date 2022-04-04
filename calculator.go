package main

// Standaard package die nodig is
import (
	"fmt"
)

// Soort library voor het ondersteunen van het printen van lines
func main() {
	fmt.Println("The calculator")
	fmt.Println("This calculator currently has these options")
	fmt.Println("1. Add (1)")
	fmt.Println("2. Subtract (2)")
	fmt.Println("3. Multiply (3)")
	fmt.Println("4. Divide (4)")
	fmt.Println("5. Modulus (5)")
	defer calculate()
}
func calculate() {
	var number1 int
	var number2 int
	var operator int
	var answer int
	fmt.Println("Please choose your first number")
	fmt.Scanln(&number1)
	fmt.Println("Please choose an operator")
	fmt.Scanln(&operator)
	fmt.Println("Please choose your second number")
	fmt.Scanln(&number2)

	if number1 > 0 && number2 > 0 {
		switch operator {
		case 1:
			answer = number1 + number2
			fmt.Println("The result is:", answer)
			playAgain()
		case 2:
			answer = number1 - number2
			fmt.Println("The result is:", answer)
			playAgain()
		case 3:
			answer = number1 * number2
			fmt.Println("The result is:", answer)
			playAgain()
		case 4:
			answer = number1 / number2
			fmt.Println("The result is:", answer)
			playAgain()
		case 5:
			answer = number1 % number2
			fmt.Println("The result is:", answer)
			playAgain()
		}
	}
}

func playAgain() {
	var option string
	fmt.Println("Would you like to play again? (Y/N)")
	fmt.Scanln(&option)

	if option == "N" || option == "n" {
		fmt.Println("Very well, please restart the program if you want to use this calculator again")
	} else if option == "Y" || option == "y" {
		fmt.Println("Very well. Restarting program")
		defer main()
	}
}
