package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/colorm"
)

type Alien struct {
	img    *ebiten.Image
	box    Box
	player bool
	state  alienState
}

type alienState int

const (
	stateAlive alienState = iota
	stateDead
)

type Aliens []*Alien

func (a *Alien) Draw(screen *ebiten.Image) {
	if !a.player {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(a.box.X, a.box.Y)
		op.GeoM.Scale(a.box.Scale, a.box.Scale)
		screen.DrawImage(a.img, op)
		return
	}
	// TODO unify common
	op := &colorm.DrawImageOptions{}
	op.GeoM.Translate(a.box.X, a.box.Y)
	op.GeoM.Scale(a.box.Scale, a.box.Scale)
	// Reset RGB (not Alpha) 0 forcibly
	var cm colorm.ColorM
	cm.Scale(0, 0, 0, 1)

	// Set color
	cm.Translate(200, 0, 0, 0)
	colorm.DrawImage(screen, a.img, cm, op)
}

func (a Aliens) getPlayers() Aliens {
	var aliens Aliens
	for _, v := range a {
		if v.player {
			aliens = append(aliens, v)
		}
	}
	return aliens
}
func (a Aliens) getNoPlayers() Aliens {
	var aliens Aliens
	for _, v := range a {
		if !v.player {
			aliens = append(aliens, v)
		}
	}
	return aliens
}
