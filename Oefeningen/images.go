package main

import (
	"fmt"
	"log"

	"github.com/disintegration/imaging"
)

func main() {
	var option int
	var path string
	var width int
	var height int

	fmt.Println("Photo editor")
	fmt.Println("1. Resize image")
	fmt.Println("2. Change image (color, gamma, contrast)")
	fmt.Scanln(&option)

	if option == 1 {
		fmt.Println("Please choose the image")
		fmt.Scanln(&path)
		fmt.Println("Please choose the width (px)")
		fmt.Scanln(&width)
		fmt.Println("Please choose the height")
		fmt.Scanln(&height)
		resizeImage(width, height, path)
	} else if option == 2 {
		fmt.Println("TBC")
	}

}

func resizeImage(width, height int, path string) {
	src, err := imaging.Open(path)
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}
	dstImage128 := imaging.Resize(src, width, height, imaging.Lanczos)
	fmt.Println("File name to save")
	var name string
	fmt.Scanln(&name)
	err = imaging.Save(dstImage128, name)
	if err != nil {
		log.Fatalf("failed to save image: %v", err)
	}
	defer fmt.Println("Image added")
}
