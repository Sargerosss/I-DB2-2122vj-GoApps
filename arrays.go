package main

// Standaard package die nodig is
import (
	"fmt"
)

// Soort library voor het ondersteunen van het printen van lines
func main() {
	mapsTest()
}

func arrayTest() {
	var x [10]int
	x[5] = 100
	x[2] = 13
	x[7] = 94
	x[3] = 32
	x[8] = 4
	fmt.Println(x)
}

func scores() {
	var x [7]float64
	x[0] = 6.4
	x[1] = 7.2
	x[2] = 8.3
	x[3] = 9
	x[4] = 5.5
	x[5] = 7.1
	x[6] = 6.7

	var total float64
	for i := 0; i < len(x); i++ {
		total += x[i]
	}

	fmt.Println(total / float64(len(x)))
}

func slices() {
	/*var x []float64
	x := make([]float64, 5, 10)

	arr := [5]float64{1,2,3,4,5}
	y := arr[0:5]
	*/
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)
}

func mapsTest() {
	/* var x map[string]int
	x["key"] = 10
	fmt.Println(x)
	GENEREERT RUNTIME ERROR
	*/

	x := make(map[string]int)
	x["key"] = 10
	fmt.Println(x["key"])

	fmt.Println()

	y := make(map[int]int)
	fmt.Println(len(y))
	y[1] = 10
	fmt.Println(y[1])

	fmt.Println()
	fmt.Println(len(y))

	delete(y, 1)
	fmt.Println(len(y))
}
