package main

// Standaard package die nodig is
import (
	"fmt"
	"image"
	"image/jpeg"
	"os"

	"github.com/nfnt/resize"
)

// Soort library voor het ondersteunen van het printen van lines
func main() {
	var path string
	fmt.Println("Choose your image")
	fmt.Scanln(&path)
	imageCompressor(path)
}

func imageCompressor(path string) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file: ", err)
		os.Exit(1)
	}
	defer f.Close()
	original, err := jpeg.Decode(f)

	resizeWidth := float64(200)
	scale := resizeWidth / float64(original.Bounds().Dx())
	width, height := Scale(original.Bounds(), scale)

	img := resize.Resize(uint(width), uint(height), original, resize.Lanczos3)

	resized, err := os.Create("resized.jpg")
	if err != nil {
		fmt.Println("Error creating file: ", err)
		os.Exit(1)
	}
	defer resized.Close()

	// write new image to file
	jpeg.Encode(resized, img, nil)
	fmt.Println("Succesfully added resized.jpg")
}

func Scale(image image.Rectangle, scale float64) (width, height uint) {
	return uint(float64(image.Dx()) * scale), uint(float64(image.Dy()) * scale)
}
