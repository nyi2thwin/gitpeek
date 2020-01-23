package main

import (
	utils "./utils"
	"flag"
	"fmt"
	"os"
)

var namePtr = flag.String("u", "nyi2thwin", "github username")
var widthPtr = flag.Uint("w", 44, "width of the image")

func main() {
	flag.Parse()

	img, getAvatarError := utils.GetAvatarImg(*namePtr)

	if getAvatarError != nil {
		fmt.Println("Avatar not found or fail to get Avatar.")
		os.Exit(1)
	}

	utils.ProcessImg(img, *widthPtr)
}
