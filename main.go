package main

import (
	"bytes"
	_ "embed"
	"image"
	"image/color"
	_ "image/png"
	"log"

	ebiten "github.com/hajimehoshi/ebiten/v2"
	ebitenutil "github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

//go:embed images/gubby.png
var gubbyImageData []byte
var gubbyImage *ebiten.Image

type Game struct {
	gubby *Gubby
}
type Gubby struct {
	x, y        float64
	facingRight bool
	speed       float64
}

func (g *Game) Update() error {
	moveAmount := g.gubby.speed / 60.00

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.gubby.y -= moveAmount
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.gubby.y += moveAmount
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.gubby.x -= moveAmount
		g.gubby.facingRight = false
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.gubby.x += moveAmount
		g.gubby.facingRight = true
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{50, 50, 100, 255})
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(0.25, 0.25)

	if g.gubby.facingRight {
		op.GeoM.Scale(-1, 1)
		scaledWidth := float64(gubbyImage.Bounds().Dx()) * 0.25
		op.GeoM.Translate(scaledWidth, 0)
	}
	op.GeoM.Translate(g.gubby.x, g.gubby.y)
	screen.DrawImage(gubbyImage, op)
	ebitenutil.DebugPrint(screen, "Hello, Gubby!")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func NewGame() *Game {
	return &Game{
		gubby: &Gubby{
			x:           160,
			y:           120,
			facingRight: false,
			speed:       120,
		},
	}
}

func init() {
	img, _, err := image.Decode(bytes.NewReader(gubbyImageData))
	if err != nil {
		panic(err)
	}
	gubbyImage = ebiten.NewImageFromImage(img)
}

func main() {
	ebiten.SetWindowSize(900, 660)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
