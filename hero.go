package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"math/rand"
)

type Hero struct {
	img      *ebiten.Image
	position Box
	state    state

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
	h.shootFrame--
	h.position.X += h.position.Speed
	if h.position.X > width {
		h.position.Speed = h.position.Speed * -1
	} else if h.position.X <= 0 {
		h.position.Speed = h.position.Speed * -1
	}
	if h.shootFrame <= 0 {
		h.shootFrame = int(ebiten.DefaultTPS / rand.Float64())
		shoot := NewShoot(h.position.X/0.5, h.position.Y/0.5)
		shoot.byHero = true
		shoot.box.Speed = -5
		game.shoots = append(game.shoots, shoot)
	}

}
