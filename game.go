package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
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

	ebitenutil.DebugPrint(screen, "Hello, Gubby!")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
