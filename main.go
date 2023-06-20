package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
	"math/rand"
	"os"
)

const (
	width  = 800
	height = 600
)

func main() {
	const sampleRate = 48000
	audioContext := audio.NewContext(sampleRate)
	{
		f, err := os.Open("assets/menu.wav")
		if err != nil {
			panic(err)
		}
		defer f.Close()

		decoded, err := wav.DecodeWithSampleRate(sampleRate, f)
		if err != nil {
			panic(err)
		}

		loop := audio.NewInfiniteLoop(decoded, decoded.Length())
		p, err := audioContext.NewPlayer(loop)
		if err != nil {
			panic(err)
		}
		p.SetVolume(0.8)
		p.Play()
	}
	{
		f, err := os.Open("assets/explosion.wav")
		if err != nil {
			panic(err)
		}
		defer f.Close()

		decoded, err := wav.DecodeWithSampleRate(sampleRate, f)
		if err != nil {
			panic(err)
		}

		p, err := audioContext.NewPlayer(decoded)
		p.SetVolume(0.4)
		playExplosion = func() {
			p.Rewind()
			p.Play()
		}
	}

	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle("Reverse Invaders")

	menu, _, err := ebitenutil.NewImageFromFile("assets/menu.png")
	if err != nil {
		log.Fatal(err)
	}
	img, _, err := ebitenutil.NewImageFromFile("assets/pixel.png")
	if err != nil {
		log.Fatal(err)
	}
	defaultShootImg, _, err = ebitenutil.NewImageFromFile("assets/Shoot.png")
	if err != nil {
		log.Fatal(err)
	}

	defaultExplosionImg, _, err = ebitenutil.NewImageFromFile("assets/explosion1.png")
	if err != nil {
		log.Fatal(err)
	}

	g := &game{
		images: images{
			menu:  menu,
			alien: img,
			hero:  img,
		},
	}
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

func generateAliens(img *ebiten.Image) []*Alien {
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
