package main

// Standaard package die nodig is
import (
	"fmt"
	"os"
)

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

// Soort library voor het ondersteunen van het printen van lines
func main() {
	var filePath string
	fmt.Println("Choose file")
	fmt.Scanln(&filePath)
	file(filePath)
}

func file(path string) {
	data, err := os.ReadFile(path)
	checkErr(err)

	fmt.Print(string(data))
	f, err := os.Open(path)
	checkErr(err)

	b1 := make([]byte, 7)
	n1, err := f.Read(b1)
	checkErr(err)

	gebouwA := string(b1[:n1])
	fmt.Println(gebouwA)

	//hoogteA := make([]byte)
}
