package main

// Standaard package die nodig is
import (
	"fmt"
)

// Soort library voor het ondersteunen van het printen van lines
func main() {
	Gebouwen()
}

func Gebouwen() {

	const gebouwA = 100
	const gebouwB = 150
	const gebouwC = 80

	const verschilAB = gebouwB - gebouwA
	const verschilBC = gebouwB - gebouwC

	if gebouwB > gebouwA && gebouwB > gebouwC {
		fmt.Println("Gebouw B is het hoogst, het is", verschilAB, " meter hoger dan GebouwA en", verschilBC, " meter hoger dan gebouw C")
	}
}
