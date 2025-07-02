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

//go:embed images/hpindic.png
var carrotImageData []byte

//go:embed images/losthpindic.png
var eCarrotImageData []byte

var (
	gubbyImage   *ebiten.Image
	carrotImage  *ebiten.Image
	eCarrotImage *ebiten.Image
)

func init() {
	// Load Gubby
	img, _, err := image.Decode(bytes.NewReader(gubbyImageData))
	if err != nil {
		panic(err)
	}
	gubbyImage = ebiten.NewImageFromImage(img)

	// And the carrot
	img, _, err = image.Decode(bytes.NewReader(carrotImageData))
	if err != nil {
		panic(err)
	}
	carrotImage = ebiten.NewImageFromImage(img)

	// And also the eaten carrot

	img, _, err = image.Decode(bytes.NewReader(eCarrotImageData))
	if err != nil {
		panic(err)
	}
	eCarrotImage = ebiten.NewImageFromImage(img)
}
