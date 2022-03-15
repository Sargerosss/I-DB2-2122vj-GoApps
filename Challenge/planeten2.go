package main

// Standaard package die nodig is
import (
	"fmt"
)

// Soort library voor het ondersteunen van het printen van lines
func main() {
	fmt.Println("Hello world")
	var planet string
	var dist = 1000
	var hrs = 4
	Saturnus(dist, hrs)
	defer fmt.Scanln(&planet)
	if planet == "Saturnus" {
		var dist = 100
		var hrs = 2
		Saturnus(dist, hrs)
	} else if planet == "Jupiter" {
		fmt.Println("JUPITER")
	}
	Jupiter(float64(dist), float64(hrs))

	if "a" == "a" {

	} else if "b" == "b" {

	} else if planet == "Saturnus" {
		var saturnus = "Saturnus"
		var saturnusD = 1500
		var saturnusH = 5
		var saturnusS = saturnusD / saturnusH
		fmt.Println(saturnus, saturnusS)
	}
}

func Saturnus(distance int, hours int) {
	var speed = distance / hours
	fmt.Println("Saturnus: ")
	fmt.Println("Distance:", distance)
	fmt.Println("Hours:", hours)
	fmt.Println(speed)
}

func Jupiter(distanc float64, hour float64) {
	var Speed = distanc / hour
	fmt.Println("Jupiter: ")
	fmt.Println("Distance:", distanc)
	fmt.Println("Hours:", hour)
	fmt.Println(Speed)
}
