package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type game struct {
	hero   Hero
	aliens []*Alien
}

func (g *game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return width, height
}

func (g *game) Update() error {
	moveAlien(g.aliens[0])
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	g.hero.Draw(screen)
	for _, v := range g.aliens {
		v.Draw(screen)
	}
}

func moveAlien(alien *Alien) {
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		alien.position.Y -= 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		alien.position.Y += 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		alien.position.X -= 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		alien.position.X += 5
	}
}
