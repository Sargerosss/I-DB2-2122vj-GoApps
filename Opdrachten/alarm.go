package main

import (
	"fmt"
)

func main() {
	alarm()
}

const number = 5

func alarm() {
	i := 0
	for i <= number {
		fmt.Println("Alarm", i, "!")
		i += 1
	}
}
