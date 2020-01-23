package main

import (
	utils "./utils"
	"flag"
	"fmt"
	"github.com/nyi2thwin/color"
	"github.com/nyi2thwin/resize"
	"image"
	"os"
)

var namePtr = flag.String("u", "nyi2thwin", "github username")
var widthPtr = flag.Uint("w", 44, "width of the image")

func main() {
	flag.Parse()

	img, getAvatarError := utils.GetAvatarImg(*namePtr)

	processImg(img)

	if getAvatarError != nil {
		fmt.Println("Avatar not found or fail to get Avatar.")
		os.Exit(1)
	}
}

func processImg(img image.Image) {
	// resize the image to fit in command line
	resizedImg := resize.Resize(*widthPtr, 0, img, resize.Lanczos3)

	bounds := resizedImg.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			printPixel(resizedImg.At(x, y).RGBA())
		}
		fmt.Println("")
	}
}

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
