package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"math/rand"
)

type Star struct {
	x, y, brightness, scale float32
}

func (s *Star) Init() {
	s.scale = 64
	s.x = rand.Float32() * width * s.scale
	s.y = rand.Float32() * height * s.scale
	s.brightness = rand.Float32() * 0xff
}

func (s *Star) Draw(screen *ebiten.Image) {
	c := color.RGBA{
		R: uint8(0xbb * s.brightness / 0xff),
		G: uint8(0xdd * s.brightness / 0xff),
		B: uint8(0xff * s.brightness / 0xff),
		A: 0xff}

	vector.DrawFilledRect(screen, s.x/s.scale, s.y/s.scale, 1, 1, c, true)
}
