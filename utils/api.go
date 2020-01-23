package utils

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"net/http"
)

func GetAvatarImg(username string) (image.Image, error) {
	url := fmt.Sprintf("https://github.com/%s.png", username)
	// Just a simple GET request to the image URL
	res, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	// close the res
	defer res.Body.Close()

	// You can register another format here
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
	img, _, err := image.Decode(res.Body)

	if err != nil {
		return nil, err
	}

	return img, err
}
