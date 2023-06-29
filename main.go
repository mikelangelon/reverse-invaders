package main

import (
	"bytes"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/mikelangelon/reverse-invaders/assets"
	"log"
	"time"
)

const (
	width  = 800
	height = 600
)

func main() {
	const sampleRate = 48000
	audioContext := audio.NewContext(sampleRate)
	{
		decoded, err := wav.DecodeWithSampleRate(sampleRate, bytes.NewReader(assets.MenuWav))
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
		decoded, err := wav.DecodeWithSampleRate(sampleRate, bytes.NewReader(assets.ExplosionWav))
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

	menu, _, err := ebitenutil.NewImageFromReader(bytes.NewReader(assets.MenuPng))
	if err != nil {
		log.Fatal(err)
	}
	img, _, err := ebitenutil.NewImageFromReader(bytes.NewReader(assets.AlienPng))
	if err != nil {
		log.Fatal(err)
	}
	img2, _, err := ebitenutil.NewImageFromReader(bytes.NewReader(assets.AlienAnim2Png))
	if err != nil {
		log.Fatal(err)
	}
	heroImg, _, err := ebitenutil.NewImageFromReader(bytes.NewReader(assets.HeroPng))
	if err != nil {
		log.Fatal(err)
	}
	hero2Img, _, err := ebitenutil.NewImageFromReader(bytes.NewReader(assets.Hero2Png))
	if err != nil {
		log.Fatal(err)
	}
	hero3Img, _, err := ebitenutil.NewImageFromReader(bytes.NewReader(assets.Hero3Png))
	if err != nil {
		log.Fatal(err)
	}
	defaultShootImg, _, err = ebitenutil.NewImageFromReader(bytes.NewReader(assets.ShootPng))
	if err != nil {
		log.Fatal(err)
	}
	ballShootImg, _, err = ebitenutil.NewImageFromReader(bytes.NewReader(assets.Shoot2Png))
	if err != nil {
		log.Fatal(err)
	}
	defaultExplosionImg, _, err = ebitenutil.NewImageFromReader(bytes.NewReader(assets.ExplosionPng))
	if err != nil {
		log.Fatal(err)
	}

	g := &game{
		updateTick: time.Now(),
		images: images{
			menu:  menu,
			alien: []*ebiten.Image{img, img2},
			hero:  []*ebiten.Image{heroImg, hero2Img, hero3Img},
		},
	}
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
