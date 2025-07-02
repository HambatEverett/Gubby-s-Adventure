package main

import (
	"bytes"
	_ "embed"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed images/gubby.png
var gubbyImageData []byte

var (
	gubbyImage *ebiten.Image
)

func init() {
	// Load Gubby
	img, _, err := image.Decode(bytes.NewReader(gubbyImageData))
	if err != nil {
		panic(err)
	}
	gubbyImage = ebiten.NewImageFromImage(img)

}
