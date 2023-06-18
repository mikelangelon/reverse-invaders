package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/colorm"
)

type Alien struct {
	img    *ebiten.Image
	box    Box
	player bool
	state  state
}

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
func (a Aliens) Move(game *game) {
	var min = a[0].box.XScaled()
	var max = a[0].box.XScaled()
	for _, v := range a {
		if min > v.box.XScaled() {
			min = v.box.XScaled()
		}
		if max < v.box.XScaled() {
			max = v.box.XScaled()
		}
	}
	for _, v := range a {
		if max > width-v.box.With && (v.box.Speed) > 0 {
			v.box.Speed = v.box.Speed * -1
		} else if min <= 0 && (v.box.Speed) < 0 {
			v.box.Speed = v.box.Speed * -1
		}
		v.Move(game)
	}
}
func (a *Alien) Move(game *game) {
	a.box.X += a.box.Speed
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
