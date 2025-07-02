package main

import (
	"bytes"
	_ "embed"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed images/gubby/gubbyidle.png
var gubbyImageData []byte

//go:embed images/hpindic.png
var carrotImageData []byte

//go:embed images/losthpindic.png
var eCarrotImageData []byte

//go:embed images/gubby/sword/topswing.png
var swing1ImageData []byte

//go:embed images/gubby/sword/downswing.png
var swing2ImageData []byte

var (
	gubbyImage   *ebiten.Image
	carrotImage  *ebiten.Image
	eCarrotImage *ebiten.Image
	swing1Image  *ebiten.Image
	swing2Image  *ebiten.Image
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

	// and also also the sword

	img, _, err = image.Decode(bytes.NewReader(swing1ImageData))
	if err != nil {
		panic(err)
	}
	swing1Image = ebiten.NewImageFromImage(img)

	// and also the sword, but down and cool

	img, _, err = image.Decode(bytes.NewReader(swing2ImageData))
	if err != nil {
		panic(err)
	}
	swing2Image = ebiten.NewImageFromImage(img)
}
