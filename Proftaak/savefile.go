package main

// Standaard package die nodig is
import (
	"fmt"

	"github.com/paij0se/lmmp3"
)

// Soort library voor het ondersteunen van het printen van lines
func main() {
	fmt.Println("Hello world")
	grabFile()
}

func grabFile() {
	lmmp3.DownloadAndConvert("https://www.youtube.com/watch?v=829pvBHyG6I")
}

/*
import "github.com/cavaliercoder/grab"

resp, err := grab.Get(".", "http://www.golang-book.com/public/pdf/gobook.pdf")
if err != nil {
    log.Fatal(err)
}

fmt.Println("Download saved to", resp.Filename)
*/
