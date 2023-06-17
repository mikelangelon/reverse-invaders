package main

type Box struct {
	X      float64
	Y      float64
	With   float64
	Height float64
	Scale  float64
}

func (b Box) CollidesTo(o Box) bool {
	collision := b.X < o.X+o.With && b.X+b.With > o.X && b.Y < o.Y+o.Height && b.Y+b.Height > o.Y
	return collision
}
