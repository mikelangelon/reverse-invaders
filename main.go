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
				Y: 500,
			}},
		aliens: generateAliens(),
	}
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

func generateAliens() []Alien {
	img, _, err := ebitenutil.NewImageFromFile("assets/pixel.png")
	if err != nil {
		log.Fatal(err)
	}

	var aliens []Alien
	for i := 0; i < 8; i++ {
		for j := 0; j < 4; j++ {
			aliens = append(aliens, Alien{
				img: img,
				position: Position{
					X: float64(100 + i*100),
					Y: float64(30 + j*100),
				},
			})
		}
	}
	return aliens
}
