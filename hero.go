package main

import "github.com/hajimehoshi/ebiten/v2"

type Hero struct {
	img      *ebiten.Image
	position Box
}

func (h *Hero) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(h.position.X, h.position.Y)
	screen.DrawImage(h.img, op)
}
