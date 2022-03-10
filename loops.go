package main

import (
	"fmt"
)

func main() {
	countUp()
	splitter()
	countDown()
}

const number = 50

func countUp() {
	i := 0
	for i <= number {
		fmt.Println(i)
		i += 1
	}
}

func countDown() {
	i := number
	for i >= 0 {
		fmt.Println(i)
		i -= 1
	}
}
func splitter() {
	fmt.Println()
	fmt.Println("---------------------------")
	fmt.Println()
}

func forLoops() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
	// Andere manier om een for loop te schrijven
}
