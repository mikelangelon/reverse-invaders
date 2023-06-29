package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/colorm"
	"math"
	"math/rand"
)

type alienType int

const (
	alienDefault alienType = iota
	alienSlow
	alienFast

	alienSpeed = 3
)

type Alien struct {
	img          []*ebiten.Image
	box          Box
	player       bool
	state        state
	alienType    alienType
	shootFrame   int
	initMovDownY float64
}

type Aliens []*Alien

func (a *Alien) Draw(currentFrame int, screen *ebiten.Image) {
	i := (currentFrame / 5) % 2
	//if !a.player {
	//	op := &ebiten.DrawImageOptions{}
	//	op.GeoM.Translate(a.box.X, a.box.Y)
	//	op.GeoM.Scale(a.box.Scale, a.box.Scale)
	//	screen.DrawImage(a.img[i], op)
	//	return
	//}
	op := &colorm.DrawImageOptions{}
	op.GeoM.Translate(a.box.X, a.box.Y)
	op.GeoM.Scale(a.box.Scale, a.box.Scale)
	colorm.DrawImage(screen, a.img[i], colorByAlienType(a.alienType, a.player), op)
}
func (a Aliens) Move(game *game) {
	if len(a) == 0 {
		return
	}
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

func (a Aliens) shuffle() Aliens {
	for i := range a {
		j := rand.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}
	return a
}

func generateAliens(img []*ebiten.Image) []*Alien {
	alienType := []alienType{alienDefault, alienDefault, alienDefault, alienDefault, alienDefault, alienSlow, alienFast}[rand.Intn(7)]
	var aliens []*Alien
	for i := 0; i < 8; i++ {
		for j := 0; j < 4; j++ {
			a := &Alien{
				img:       img,
				alienType: alienType,
				box: Box{
					X:      float64(100 + i*xSeparationByAlienType(alienType)),
					Y:      float64(30 + j*100),
					With:   64,
					Height: 64,
					SpeedY: speedYByAlienType(alienType),
					Scale:  0.5,
				},
			}
			a.setShootFrame()
			aliens = append(aliens, a)
		}
	}
	// Pick random alien to be the player
	n := rand.Int() % len(aliens)
	aliens[n].player = true
	aliens[n].box.SpeedY = alienSpeed

	return aliens
}

func speedYByAlienType(alienType alienType) float64 {
	switch alienType {
	case alienSlow:
		return alienSpeed - 0.1
	case alienFast:
		return alienSpeed + 0.1
	default:
		return alienSpeed
	}
}

func colorByAlienType(alienType alienType, player bool) colorm.ColorM {
	var cm colorm.ColorM
	cm.Scale(0, 0, 0, 1)

	if player {
		cm.Translate(200, 0, 0, 0)
		return cm
	}

	switch alienType {
	case alienSlow:
		cm.Translate(0, 0, 200, 0)
	case alienFast:
		cm.Translate(200, 0, 200, 0)
	default:
		cm.Translate(0, 200, 0, 0)
	}
	return cm
}

func xSeparationByAlienType(alienType alienType) int {
	switch alienType {
	case alienSlow:
		return 110
	case alienFast:
		return 140
	default:
		return 120
	}
}
