package main

import (
	"fmt"
	"github.com/nyi2thwin/color"
	"image"
	"image/png"
	"io"
	"os"
)

func main() {
	// You can register another format here
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)

	file, err := os.Open("./test.png")

	if err != nil {
		fmt.Println("Error: File could not be opened")
		os.Exit(1)
	}

	defer file.Close()

	decodeErr := decodeAndProcess(file)

	if decodeErr != nil {
		fmt.Println("Error: Image could not be decoded")
		os.Exit(1)
	}
}

func decodeAndProcess(file io.Reader) error {
	img, _, err := image.Decode(file)

	grayImg := image.NewGray(img.Bounds())

	if err != nil {
		return err
	}

	bounds := grayImg.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			printPixel(img.At(x, y).RGBA())
		}
		fmt.Println("")
	}

	return nil
}

// img.At(x, y).RGBA() returns four uint32 values; we want a Pixel
func printPixel(r uint32, g uint32, b uint32, a uint32) {
	red := uint8(r / 257)
	green := uint8(g / 257)
	blue := uint8(b / 257)
	alpha := uint8(a / 257)
	coloredChar := color.RGB(red, green, blue)
	backgroundChar := color.RGB(255, 255, 255)
	if alpha != 0 {
		coloredChar.Print("01")
	} else {
		backgroundChar.Print("  ")
	}
}
