package main

// Standaard package die nodig is
import (
	"fmt"
)

func multiply() {
	fmt.Println("---------------------------------")
	fmt.Println("Functions")
	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2
	fmt.Println(output)
}

func f() {
	const ab = "Hello"
	//ab = "Hello World"
	// Error op line hierboven omdat een constante niet gewijzigd kan worden
	fmt.Println("Functions")
	fmt.Println("--------------------------------")
	fmt.Println(ab)

	var (
		a = 100
		b = "hello"
		c = 10.3
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// Verschil Print en Println, Println plakt automatisch \n eraan
	// Print doet het niet (meerdere Print statements blijven dus op 1 line)
}

// Soort library voor het ondersteunen van het printen van lines
func main() {
	fmt.Println("Variables")
	var x string = "Hello World"
	fmt.Println(x)
	var y string = " Martijn"
	fmt.Println("Hello, my name is" + y)
	var z int = 18
	fmt.Println(z + 17)
	var a string
	a = "first"
	fmt.Println(a)
	a = "second"
	fmt.Println(a)
	var ab string = "hello"
	var bc string = "world"
	fmt.Println(ab == bc)

	var aa = "Short way"
	fmt.Println(aa)

	ac := "help"
	fmt.Println(ac)
	// Naam variables:
	// Geef variables een goede en logische naam, wil je een bijvoorbeeld Martijn als waarde geven,
	// dat is een naam dus dan is het logisch als de variable iets met naam als naam heeft

	// Scope
	// Zorg ervoor dat alles goed staat (juiste locatie van variables, global waar nodig, local waar nodig)

	// Constants
	// Deze kunnen geen nieuwe waarde krijgen (zie nieuwe functie hierboven)
	f()
	multiply()
}
