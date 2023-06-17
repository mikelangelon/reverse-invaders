package main

import (
	"github.com/hajimehoshi/ebiten/v2"
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

	if err := ebiten.RunGame(&game{}); err != nil {
		log.Fatal(err)
	}
}
