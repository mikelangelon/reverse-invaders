package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type game struct {
	hero   *Hero
	aliens Aliens
	shoots []*Shoot
}

func (g *game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return width, height
}

func (g *game) Update() error {
	g.updatePositions()
	g.updateActions()
	g.updateCollisions()
	g.cleanDeads()

	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	g.hero.Draw(screen)
	for _, v := range g.aliens {
		v.Draw(screen)
	}
	for _, v := range g.shoots {
		v.Draw(screen)
	}
}

func (g *game) updatePositions() {
	g.moveAliens()
	g.moveShoots()
}

func (g *game) moveAliens() {
	for _, v := range g.aliens {
		if v.player {
			g.moveAlien(v)
		}
	}
}
func (g *game) moveShoots() {
	for _, v := range g.shoots {
		v.box.Y += v.box.Speed
	}
}

func (g *game) updateActions() {
	g.hero.Move(g)
	g.aliens.getNoPlayers().Move(g)
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
	for _, v := range g.aliens {
		for _, z := range g.shoots {
			if !z.byHero {
				continue
			}
			if AreColliding(v.box, z.box) {
				v.state = stateDead
				z.state = stateDead
			}
		}
	}

}

func (g *game) heroCollisions() {
	for _, v := range g.shoots {
		if v.byHero {
			continue
		}
		if v.box.CollidesTo(g.hero.position) {
			g.hero.state = stateDead
			v.state = stateDead
		}
	}
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

	var shoot []*Shoot
	for _, v := range g.shoots {
		if v.state == stateDead {
			continue
		}
		//if v.box.ConvertY() > height || v.box.ConvertY() < 0 {
		//	continue
		//}
		shoot = append(shoot, v)
	}
	g.shoots = shoot
}
func (g *game) moveAlien(alien *Alien) {
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
	if ebiten.IsKeyPressed(ebiten.KeyEnter) {
		g.shoots = append(g.shoots, NewShoot(alien.box.X, alien.box.Y))
	}
}

func AreColliding(b1, b2 Box) bool {
	return b1.CollidesTo(b2)
}
