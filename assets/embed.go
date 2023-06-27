package assets

import (
	_ "embed"
)

var (
	//go:embed menu.wav
	MenuWav []byte

	//go:embed explosion.wav
	ExplosionWav []byte

	//go:embed menu.png
	MenuPng []byte

	//go:embed pixel.png
	AlienPng []byte

	//go:embed pixel2.png
	AlienAnim2Png []byte

	//go:embed hero.png
	HeroPng []byte

	//go:embed explosion.png
	ExplosionPng []byte

	//go:embed shoot.png
	ShootPng []byte
)
