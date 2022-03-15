package main

// Standaard package die nodig is
import (
	"fmt"
)

// Soort library voor het ondersteunen van het printen van lines
func main() {

	defer fmt.Println("Hello world")
	fmt.Println("Defer")
}
