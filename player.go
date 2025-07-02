package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Gubby struct {
	x, y        float64
	facingRight bool
	speed       float64
	health      int
	swordAct    bool
	swordTimer  int
}

func NewGubby(x, y float64) *Gubby {
	return &Gubby{
		x:           x,
		y:           y,
		facingRight: false,
		speed:       240,
		health:      5,
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
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		if g.health > 0 {
			g.health--
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyComma) {
		if g.health < 5 {
			g.health++
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyR) {
		g.swordAct = true
		g.swordTimer = 15
	}

	if g.swordAct {
		g.swordTimer--
		if g.swordTimer <= 0 {
			g.swordAct = false
		}
	}

}

func (g *Gubby) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.Filter = ebiten.FilterLinear
	op.GeoM.Scale(0.40, 0.40)
	if g.facingRight {
		op.GeoM.Scale(-1, 1)
		scaledWidth := float64(gubbyImage.Bounds().Dx()) * 0.40
		op.GeoM.Translate(scaledWidth, 0)
	}

	op.GeoM.Translate(-70, -70)
	op.GeoM.Translate(g.x, g.y)

	currentSprite := gubbyImage
	if g.swordAct {
		currentSprite = swing1Image
		if g.swordTimer < 14 {
			currentSprite = swing2Image
		}
	}
	screen.DrawImage(currentSprite, op)
}
