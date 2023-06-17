package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type game struct{}

func (g *game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return width, height
}

func (g *game) Update() error {
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {}
