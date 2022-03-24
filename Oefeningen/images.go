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

	var option2 int

	var options int

	fmt.Println("Photo editor")
	fmt.Println("1. Resize image")
	fmt.Println("2. Change image (brightness, contrast, rotate, sharpen)")
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
		fmt.Println("Please choose an option:")
		fmt.Println("1. Change contrast")
		fmt.Println("2. Change brightness")
		fmt.Println("3. Rotate (multiple options)")
		fmt.Println("4. Sharpen image")
		fmt.Scanln(&option2)

		if option == 1 {
			fmt.Println("Please choose an image")
			var image string
			fmt.Scanln(&image)
			fmt.Println("Please choose a percentage")
			var percent float64
			fmt.Scanln(&percent)
			ChangeContrast(image, percent)

		} else if option == 2 {
			fmt.Println("Please choose an image")
			var image string
			fmt.Scanln(&image)
			fmt.Println("Please choose a percentage")
			var percent float64
			fmt.Scanln(&percent)
			ChangeBrightness(image, percent)
		} else if option == 3 {
			fmt.Println("Please choose an option:")
			fmt.Println("1. Rotate image (90)")
			fmt.Println("2. Rotate image (180)")
			fmt.Println("3. Rotate image (270)")

			fmt.Scanln(&options)

			if options == 1 {
				fmt.Println("Rotate image (90)")
				fmt.Println("Please choose an image to rotate")
				var img string
				fmt.Scanln(&img)
				IMGRotate90(img)

			} else if options == 2 {
				fmt.Println("Rotate image (180)")
				fmt.Println("Please choose an image to rotate")
				var img string
				fmt.Scanln(&img)
				IMGRotate180(img)

			} else if options == 3 {
				fmt.Println("Rotate image (270)")
				fmt.Println("Please choose an image to rotate")
				var img string
				fmt.Scanln(&img)
				IMGRotate270(img)
			} else if options == 4 {
				fmt.Println("Sharpen image")
				fmt.Println("Please choose an image to sharpen")
				var image string
				fmt.Scanln(&image)
				fmt.Println("Please choose how much you want to sharpen it")
				var sharpen float64
				fmt.Scanln(&sharpen)
				SharpenImage(image, sharpen)
			}
		}
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

func ChangeContrast(image string, percentage float64) {
	src, err := imaging.Open(image)
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}
	dstImage := imaging.AdjustContrast(src, percentage)
	fmt.Println("File name to save")
	var fn string
	fmt.Scanln(&fn)
	err = imaging.Save(dstImage, fn)
	if err != nil {
		log.Fatalf("failed to save image: %v", err)
	}
	defer fmt.Println("Image added")
}

func ChangeBrightness(image string, percentage float64) {
	src, err := imaging.Open(image)
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}
	dstImage := imaging.AdjustBrightness(src, percentage)
	fmt.Println("File name to save")
	var fn string
	fmt.Scanln(&fn)
	err = imaging.Save(dstImage, fn)
	if err != nil {
		log.Fatalf("failed to save image: %v", err)
	}
	defer fmt.Println("Image added")
}

func IMGRotate90(img string) {
	src, err := imaging.Open(img)
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}
	dstImage := imaging.Rotate90(src)
	fmt.Println("File name to save")
	var fn string
	fmt.Scanln(&fn)
	err = imaging.Save(dstImage, fn)
	if err != nil {
		log.Fatalf("failed to save image: %v", err)
	}
	defer fmt.Println("Image added")
}

func IMGRotate180(img string) {
	src, err := imaging.Open(img)
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}
	dstImage := imaging.Rotate180(src)
	fmt.Println("File name to save")
	var fn string
	fmt.Scanln(&fn)
	err = imaging.Save(dstImage, fn)
	if err != nil {
		log.Fatalf("failed to save image: %v", err)
	}
	defer fmt.Println("Image added")
}

func IMGRotate270(img string) {
	src, err := imaging.Open(img)
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}
	dstImage := imaging.Rotate270(src)
	fmt.Println("File name to save")
	var fn string
	fmt.Scanln(&fn)
	err = imaging.Save(dstImage, fn)
	if err != nil {
		log.Fatalf("failed to save image: %v", err)
	}
	defer fmt.Println("Image added")
}

func SharpenImage(img string, percentage float64) {
	src, err := imaging.Open(img)
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}
	dstImage := imaging.Sharpen(src, percentage)
	fmt.Println("File name to save")
	var fn string
	fmt.Scanln(&fn)
	err = imaging.Save(dstImage, fn)
	if err != nil {
		log.Fatalf("failed to save image: %v", err)
	}
	defer fmt.Println("Image added")
}
