package main

import "github.com/hajimehoshi/ebiten/v2"

type Gubby struct {
	x, y        float64
	facingRight bool
	speed       float64
}

func NewGubby(x, y float64) *Gubby {
	return &Gubby{
		x:           x,
		y:           y,
		facingRight: false,
		speed:       120,
	}
}

func (g *Gubby) Update() {
	moveAmount := g.speed / 60.0

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.y -= moveAmount
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.y += moveAmount
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.x -= moveAmount
		g.facingRight = false
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.x += moveAmount
		g.facingRight = true
	}
}

func (g *Gubby) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(0.25, 0.25)

	if g.facingRight {
		op.GeoM.Scale(-1, 1)
		scaledWidth := float64(gubbyImage.Bounds().Dx()) * 0.25
		op.GeoM.Translate(scaledWidth, 0)
	}
	op.GeoM.Translate(g.x, g.y)
	screen.DrawImage(gubbyImage, op)
}
