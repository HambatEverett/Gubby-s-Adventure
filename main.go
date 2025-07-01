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
	x, y float64
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.gubby.y -= 2 // Move up
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.gubby.y += 2 // Move down
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.gubby.x -= 2 // Move left
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.gubby.x += 2 // Move right
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{50, 50, 100, 255})

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(0.3, 0.3)
	op.GeoM.Translate(g.gubby.x, g.gubby.y)
	screen.DrawImage(gubbyImage, op)

	ebitenutil.DebugPrint(screen, "Hello, Gubby!")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func NewGame() *Game {
	return &Game{
		gubby: &Gubby{x: 160, y: 120}, // Start in center of screen
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
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
