package main

import "github.com/hajimehoshi/ebiten/v2"

var (
	defaultShootImg *ebiten.Image
	ballShootImg    *ebiten.Image
)

type Shoot struct {
	box    Box
	img    *ebiten.Image
	byHero bool
	state  state
}

func NewShoot(x, y float64, t shootType) *Shoot {
	s := &Shoot{
		box: Box{
			X:      x,
			Y:      y,
			With:   10,
			Height: 64,
			SpeedY: 5,
			Scale:  0.5,
		},
		img: defaultShootImg,
	}
	switch t {
	case ballShoot:
		s.img = ballShootImg
		s.box.With = 32
		s.box.Height = 32
	}
	return s
}
func (h *Shoot) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(h.box.X, h.box.Y)
	op.GeoM.Scale(0.5, 0.5)
	screen.DrawImage(h.img, op)
}

type shootType int

const (
	defaultShoot shootType = iota
	threeShoots
	ballShoot
)
