package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"math/rand"
)

type Hero struct {
	img      *ebiten.Image
	position Box
	state    state

	shootType  shootType
	shootFrame int
}

func (h *Hero) Draw(screen *ebiten.Image) {
	if h.state == stateDead {
		return
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(h.position.X, h.position.Y)
	screen.DrawImage(h.img, op)
}

func (h *Hero) Move(game *game) {
	if h.state == stateDead {
		return
	}
	h.shootFrame--
	h.position.X += h.position.SpeedY
	if h.position.X > width {
		h.position.SpeedY = h.position.SpeedY * -1
	} else if h.position.X <= 0 {
		h.position.SpeedY = h.position.SpeedY * -1
	}
	if h.shootFrame <= 0 {
		switch h.shootType {
		case defaultShoot:
			h.addShoot(game, h.position.X/0.5+64, h.position.Y/0.5, 0, -5)
		case threeShoots:
			h.addShoot(game, h.position.X/0.5+64-20, h.position.Y/0.5+20, 0, -5)
			h.addShoot(game, h.position.X/0.5+64, h.position.Y/0.5, 0, -5)
			h.addShoot(game, h.position.X/0.5+64+20, h.position.Y/0.5+20, 0, -5)
		case ballShoot:
			h.addShoot(game, h.position.X/0.5+64-20, h.position.Y/0.5+20, 0, -10)
		}
	}

}

func (h *Hero) addShoot(game *game, x, y, speedX, speedY float64) {
	h.shootFrame = int(ebiten.DefaultTPS / 1.2 * rand.Float64())
	shoot := NewShoot(x, y, ballShoot)
	shoot.byHero = true
	shoot.box.SpeedY = speedY
	game.shoots = append(game.shoots, shoot)
}
