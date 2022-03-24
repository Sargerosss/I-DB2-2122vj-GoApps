package main

import (
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	args := os.Args

	distance := args[1]
	hours := args[2]
	fmt.Println("Planeten")
	fmt.Println("Syntax:")
	fmt.Println("go run planeten.go <distance> <hours>")
	dist, err := strconv.Atoi(distance)
	hour, errs := strconv.Atoi(hours)

	check(err)
	check(errs)

	calculateSpeed(dist, hour)
}

func calculateSpeed(distance int, hours int) {
	var speed = distance / hours
	fmt.Println(speed, "km/u")

}
