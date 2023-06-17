package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

const (
	width  = 800
	height = 600
)

func main() {
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle("Reverse Invaders")

	img, _, err := ebitenutil.NewImageFromFile("assets/pixel.png")
	if err != nil {
		log.Fatal(err)
	}

	g := &game{
		hero: Hero{
			img: img,
			position: Position{
				X: 100,
				Y: 100,
			}},
	}
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
