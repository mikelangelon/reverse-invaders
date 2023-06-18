package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
	"math/rand"
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
	defaultShootImg, _, err = ebitenutil.NewImageFromFile("assets/Shoot.png")
	if err != nil {
		log.Fatal(err)
	}

	g := &game{
		hero: &Hero{
			img: img,
			position: Box{
				X:      100,
				Y:      500,
				With:   64,
				Height: 64,
				Scale:  1,
				Speed:  5,
			}},
		aliens: generateAliens(),
	}
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

func generateAliens() []*Alien {
	img, _, err := ebitenutil.NewImageFromFile("assets/pixel.png")
	if err != nil {
		log.Fatal(err)
	}

	var aliens []*Alien
	for i := 0; i < 8; i++ {
		for j := 0; j < 4; j++ {
			a := &Alien{
				img: img,
				box: Box{
					X:      float64(100 + i*120),
					Y:      float64(30 + j*100),
					With:   64,
					Height: 64,
					Speed:  3,
					Scale:  0.5,
				},
			}
			a.setShootFrame()
			aliens = append(aliens, a)
		}
	}
	// Pick random alien to be the player
	n := rand.Int() % len(aliens)
	aliens[n].player = true

	return aliens
}
