package main

// Standaard package die nodig is
import (
	"fmt"
)

// Soort library voor het ondersteunen van het printen van lines
func main() {
	fmt.Println("Data types")
	fmt.Println("Strings")
	fmt.Println(len("Hello World"))
	// Hoeveel characters heeft Hello World: 11
	fmt.Println("Hello World"[1])
	// Index start op 0, dus H = 0, e = 1. e wordt neergezet in bytes (101)
	fmt.Println("Hello " + "World")
	// Combineren van 2 strings, (let op de spatie)
	fmt.Println()
	fmt.Println("Booleans")
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
	// && betekent AND
	// || betekent OR
	// ! betekent NOT
}
