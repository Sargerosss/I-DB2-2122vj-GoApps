package main

import (
	"fmt"
)

func main() {
	var distance int
	var hours int
	fmt.Println("Planeten")
	fmt.Println("Afstand")
	fmt.Scanln(&distance)
	fmt.Println("Tijd")
	fmt.Scanln(&hours)
	calculateSpeed(distance, hours)
}

func calculateSpeed(distance int, hours int) {
	var speed = distance / hours
	fmt.Println(speed)

}
