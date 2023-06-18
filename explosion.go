package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
)

var defaultExplosionImg *ebiten.Image

const initialScale = 0.3

type Explosion struct {
	img *ebiten.Image
	box Box

	currentFrame int
	lastFrame    int
}

func NewExplosion(x, y float64) *Explosion {
	e := &Explosion{
		box: Box{
			X:      x,
			Y:      y,
			With:   64,
			Height: 64,
			Scale:  initialScale,
		},
		img:       defaultExplosionImg,
		lastFrame: 10,
	}
	fmt.Print("creating explosion on the x: %v, translated to %v and later as %v", x, e.box.X, e.box.XScaled())
	return e
}

func (e *Explosion) Update() {
	e.currentFrame++
	e.box.Scale += 0.05
}
func (e *Explosion) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(e.box.X/e.box.Scale, e.box.Y/e.box.Scale)
	op.GeoM.Scale(e.box.Scale, e.box.Scale)
	screen.DrawImage(e.img, op)
}

func (e *Explosion) ToBeCleaned() bool {
	return e.currentFrame > e.lastFrame
}
