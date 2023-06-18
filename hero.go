package main

import "github.com/hajimehoshi/ebiten/v2"

type Hero struct {
	img      *ebiten.Image
	position Box
	state    state
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
	h.position.X += h.position.Speed
	if h.position.X > width {
		h.position.Speed = h.position.Speed * -1
	} else if h.position.X <= 0 {
		h.position.Speed = h.position.Speed * -1
	}
}
