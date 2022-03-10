package main

// Standaard package die nodig is
import "fmt"

// Soort library voor het ondersteunen van het printen van lines
func main() {
	var F float64
	fmt.Println("Temperatuur in Fahrenheit: ")
	fmt.Scanf("%f", &F)
	var C = (F - 32) * 5 / 9
	fmt.Println(C)

}
