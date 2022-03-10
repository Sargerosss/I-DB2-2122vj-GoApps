package main

// Standaard package die nodig is
import "fmt"

// Soort library voor het ondersteunen van het printen van lines
func main() {
	fmt.Println(value(40))
	fmt.Println("----------------------")
	add := func(x, y int) int {
		return x + y
	}
	fmt.Println("----------------------")
	defer fmt.Println(add(1, 1))
	nextEven := even()
	var count = 5
	for i := 0; i < count; i++ {
		fmt.Println(nextEven())
	}
	fmt.Println("----------------------")
	fmt.Println(factorial(5))
}
func value(x int) (y int) {
	var number = 20*40 + x
	return returnValue(number)
}
func returnValue(x int) (y int) {
	y = x
	return y
}

func even() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

// op deze manier werkt return, de 2e haakjes geven aan wat je terug wilt, dus bijvoorbeeld een int in de vorm van y
