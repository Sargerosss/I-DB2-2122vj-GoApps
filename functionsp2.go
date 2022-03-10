package main

// Standaard package die nodig is
import "fmt"

// Soort library voor het ondersteunen van het printen van lines
func main() {
	//panic("HELP")
	//	y := recover()
	//fmt.Println(y)

	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("PANIC")
}
