package main

// Standaard package die nodig is
import (
	"fmt"
	"log"

	"github.com/cavaliercoder/grab"
)

// Soort library voor het ondersteunen van het printen van lines
func main() {
	fmt.Println("Hello world")
	grabFile()
}

func grabFile() {
	resp, err := grab.Get("C:/Users/HP/Desktop/PROFTAAK S2/FG7M2S2GF5VAQ9Q/BeatWrite/data", "https://github.com/Sargerosss/I-DB2-2122vj-GoApps/blob/main/hw.mp3")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Added file", resp.Filename)
}

/*
import "github.com/cavaliercoder/grab"

resp, err := grab.Get(".", "http://www.golang-book.com/public/pdf/gobook.pdf")
if err != nil {
    log.Fatal(err)
}

fmt.Println("Download saved to", resp.Filename)
*/
