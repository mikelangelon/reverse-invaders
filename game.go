package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type game struct {
	hero   Hero
	aliens Aliens
}

func (g *game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return width, height
}

func (g *game) Update() error {
	g.updatePositions()

	g.updateCollisions()

	g.cleanDeads()

	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	g.hero.Draw(screen)
	for _, v := range g.aliens {
		v.Draw(screen)
	}
}

func (g *game) updatePositions() {
	for _, v := range g.aliens {
		if v.player {
			moveAlien(v)
		}
	}
}

func (g *game) updateCollisions() {
	g.playerToAlienCollisions()
	g.alienToShootsCollisions()
	g.heroCollisions()
}
func (g *game) playerToAlienCollisions() {
	player := g.aliens.getPlayers()
	noPlayers := g.aliens.getNoPlayers()
	for _, v := range player {
		for _, z := range noPlayers {
			if AreColliding(v.box, z.box) {
				v.state = stateDead
				z.state = stateDead
			}
		}
	}
}

func (g *game) alienToShootsCollisions() {

}

func (g *game) heroCollisions() {

}

func (g *game) cleanDeads() {
	var alive Aliens
	for _, v := range g.aliens {
		if v.state == stateDead {
			continue
		}
		alive = append(alive, v)
	}
	g.aliens = alive
}
func moveAlien(alien *Alien) {
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		alien.box.Y -= 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		alien.box.Y += 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		alien.box.X -= 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		alien.box.X += 5
	}
}

func AreColliding(b1, b2 Box) bool {
	return b1.CollidesTo(b2)
}
