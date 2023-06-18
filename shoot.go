package main

import "github.com/hajimehoshi/ebiten/v2"

var defaultShootImg *ebiten.Image

type Shoot struct {
	box    Box
	img    *ebiten.Image
	byHero bool
	state  state
}

func NewShoot(x, y float64) *Shoot {
	return &Shoot{
		box: Box{
			X:      x,
			Y:      y,
			With:   10,
			Height: 64,
			Speed:  5,
			Scale:  0.5,
		},
		img: defaultShootImg,
	}
}
func (h *Shoot) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(h.box.X, h.box.Y)
	op.GeoM.Scale(0.5, 0.5)
	screen.DrawImage(h.img, op)
}
