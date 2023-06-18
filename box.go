package main

type Box struct {
	X      float64
	Y      float64
	With   float64
	Height float64
	Scale  float64
	Speed  float64
}

func (b Box) CollidesTo(o Box) bool {
	collision := b.XScaled() < o.XScaled()+o.WithScaled() &&
		b.XScaled()+b.WithScaled() > o.XScaled() &&
		b.YScaled() < o.YScaled()+o.HeightScaled() &&
		b.YScaled()+b.HeightScaled() > o.YScaled()
	return collision
}

func (b Box) XScaled() float64 {
	return b.X * b.Scale
}

func (b Box) YScaled() float64 {
	return b.Y * b.Scale
}

func (b Box) WithScaled() float64 {
	return b.With * b.Scale
}

func (b Box) HeightScaled() float64 {
	return b.Height * b.Scale
}
