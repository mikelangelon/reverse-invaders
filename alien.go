package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/colorm"
)

type Alien struct {
	img      *ebiten.Image
	position Position
	player   bool
}

func (a *Alien) Draw(screen *ebiten.Image) {
	if !a.player {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(a.position.X, a.position.Y)
		op.GeoM.Scale(0.5, 0.5)
		screen.DrawImage(a.img, op)
		return
	}
	// TODO unify common
	op := &colorm.DrawImageOptions{}
	op.GeoM.Translate(a.position.X, a.position.Y)
	op.GeoM.Scale(0.5, 0.5)
	// Reset RGB (not Alpha) 0 forcibly
	var cm colorm.ColorM
	cm.Scale(0, 0, 0, 1)

	// Set color
	cm.Translate(200, 0, 0, 0)
	colorm.DrawImage(screen, a.img, cm, op)
}
