package main

import "github.com/hajimehoshi/ebiten/v2"

type Alien struct {
	img      *ebiten.Image
	position Position
}

func (a *Alien) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(a.position.X, a.position.Y)
	op.GeoM.Scale(0.5, 0.5)
	screen.DrawImage(a.img, op)
}
