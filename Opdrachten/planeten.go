package main

import (
	"fmt"
)

const distance = 150000000
const hours = 8766

func main() {
	calculateSpeed(distance, hours)
}

func calculateSpeed(distance int, hours int) {
	var speed = distance / hours
	fmt.Println(speed)

}
