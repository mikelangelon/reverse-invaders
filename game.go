package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type game struct {
	hero   Hero
	aliens []Alien
}

func (g *game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return width, height
}

func (g *game) Update() error {
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	g.hero.Draw(screen)
	for _, v := range g.aliens {
		v.Draw(screen)
	}
}
