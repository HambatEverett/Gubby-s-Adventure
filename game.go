package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	gubby *Gubby
}

func NewGame() *Game {
	return &Game{
		gubby: NewGubby(450, 330), // Using the constructor from player.go
	}
}

func (g *Game) Update() error {
	g.gubby.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{50, 50, 100, 255})
	g.gubby.Draw(screen)

	for i := 0; i < 5; i++ {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(0.5, 0.5)

		w := float64(carrotImage.Bounds().Dx()) * 0.5
		h := float64(carrotImage.Bounds().Dy()) * 0.5

		op.GeoM.Translate(-w/2, -h/2)
		op.GeoM.Rotate(-0.2)
		op.GeoM.Translate(w/2, h/2)
		op.GeoM.Translate(float64(20+i*40), 540)

		if i < int(g.gubby.health) {
			screen.DrawImage(carrotImage, op)
		} else {
			screen.DrawImage(eCarrotImage, op)
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
