package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/colorm"
	"math"
	"math/rand"
)

type Alien struct {
	img          []*ebiten.Image
	box          Box
	player       bool
	state        state
	shootFrame   int
	initMovDownY float64
}

type Aliens []*Alien

func (a *Alien) Draw(currentFrame int, screen *ebiten.Image) {
	i := (currentFrame / 5) % 2
	if !a.player {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(a.box.X, a.box.Y)
		op.GeoM.Scale(a.box.Scale, a.box.Scale)
		screen.DrawImage(a.img[i], op)
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
	colorm.DrawImage(screen, a.img[i], cm, op)
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
		switch v.state {
		case stateMovingHoritzontally:
			if max > width-v.box.With && (v.box.SpeedY) > 0 {
				v.state = stateMovingDown
				v.initMovDownY = v.box.Y
			} else if min <= 0 && (v.box.SpeedY) < 0 {
				v.state = stateMovingDown
				v.initMovDownY = v.box.Y
			}
		case stateMovingDown:
			if v.initMovDownY+30 < v.box.Y {
				if max > width-v.box.With && (v.box.SpeedY) > 0 {
					v.state = stateMovingHoritzontally
					v.box.SpeedY = v.box.SpeedY * -1
				} else if min <= 0 && (v.box.SpeedY) < 0 {
					v.state = stateMovingHoritzontally
					v.box.SpeedY = v.box.SpeedY * -1
				}
			}
		}
		v.Move(game)
	}
}
func (a *Alien) Move(game *game) {
	a.shootFrame--
	switch a.state {
	case stateMovingHoritzontally:
		a.box.X += a.box.SpeedY
	case stateMovingDown:
		a.box.Y += math.Abs(a.box.SpeedY)
	}

	if a.shootFrame <= 0 {
		a.setShootFrame()
		shoot := NewShoot(a.box.X, a.box.Y, defaultShoot)
		shoot.box.SpeedY = 5
		game.shoots = append(game.shoots, shoot)
	}
}

func (a *Alien) setShootFrame() {
	a.shootFrame = int(ebiten.DefaultTPS / (rand.Float64() * 0.1))
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
