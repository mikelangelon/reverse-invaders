package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"math/rand"
)

type Hero struct {
	img      *ebiten.Image
	position Box
	state    state

	shootType      shootType
	shootFrame     int
	movingStrategy movingStrategy
}

type movingStrategy struct {
	wallToWall    int
	adaptToBullet int
	attackLowest  int
	YThreshold    float64
	XRange        float64
}

func (m movingStrategy) all() int {
	return m.wallToWall + m.adaptToBullet + m.attackLowest
}

func (h *Hero) Draw(screen *ebiten.Image) {
	if h.state == stateDead {
		return
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(h.position.X, h.position.Y)
	screen.DrawImage(h.img, op)
}

func (h *Hero) Move(game *game) {
	if len(game.aliens) == 0 {
		return
	}
	if h.state == stateDead {
		return
	}
	h.shootFrame--

	all := h.movingStrategy.all()
	choosenStrategy := rand.Intn(all)
	if choosenStrategy <= h.movingStrategy.wallToWall {
		h.position.X += h.position.SpeedY
		if h.position.X > width {
			h.position.SpeedY = h.position.SpeedY * -1
		} else if h.position.X <= 0 {
			h.position.SpeedY = h.position.SpeedY * -1
		}
	} else if choosenStrategy > h.movingStrategy.adaptToBullet && choosenStrategy <= h.movingStrategy.adaptToBullet {
		futurePosition := h.position.X + h.position.SpeedY
		evenMoreFuturePosition := futurePosition + 2*h.position.SpeedY
		safe := true
		for _, v := range game.shoots {
			if v.byHero {
				continue
			}
			if v.box.YScaled() < h.movingStrategy.YThreshold {
				continue
			}
			if v.box.XScaled() > evenMoreFuturePosition-h.movingStrategy.XRange && v.box.XScaled() < evenMoreFuturePosition+h.movingStrategy.XRange {
				safe = false
			}
		}
		if safe {
			h.position.X = futurePosition
			if h.position.X > width {
				h.position.SpeedY = h.position.SpeedY * -1
			} else if h.position.X <= 0 {
				h.position.SpeedY = h.position.SpeedY * -1
			}
		} else {
			h.position.SpeedY = h.position.SpeedY * -1
			h.position.X += h.position.SpeedY
			if h.position.X > width {
				h.position.SpeedY = h.position.SpeedY * -1
			} else if h.position.X <= 0 {
				h.position.SpeedY = h.position.SpeedY * -1
			}
		}
	} else {
		aliens := game.aliens.shuffle()
		lowestAlien := aliens[0]
		for _, v := range aliens {
			if v.box.YScaled() > lowestAlien.box.YScaled() {
				lowestAlien = v
			}
		}
		if h.position.X > lowestAlien.box.XScaled()-lowestAlien.box.With/2 {
			h.position.X -= h.position.SpeedY
		} else {
			h.position.X += h.position.SpeedY
		}
	}

	if h.shootFrame <= 0 {
		switch h.shootType {
		case defaultShoot:
			h.addShoot(game, h.position.X/0.5+64, h.position.Y/0.5, 0, -5, h.shootType)
		case threeShoots:
			h.addShoot(game, h.position.X/0.5+64-20, h.position.Y/0.5+20, 0, -5, h.shootType)
			h.addShoot(game, h.position.X/0.5+64, h.position.Y/0.5, 0, -5, h.shootType)
			h.addShoot(game, h.position.X/0.5+64+20, h.position.Y/0.5+20, 0, -5, h.shootType)
		case ballShoot:
			h.addShoot(game, h.position.X/0.5+64-20, h.position.Y/0.5+20, 0, -10, h.shootType)
		case diagonalBallShoot:
			h.addShoot(game, h.position.X/0.5+64-10, h.position.Y/0.5+20, -2, -3, h.shootType)
			h.addShoot(game, h.position.X/0.5+64-5, h.position.Y/0.5+20, -1, -3, h.shootType)
			h.addShoot(game, h.position.X/0.5+64, h.position.Y/0.5+20, 0, -3, h.shootType)
			h.addShoot(game, h.position.X/0.5+64+5, h.position.Y/0.5+20, 1, -3, h.shootType)
			h.addShoot(game, h.position.X/0.5+64-20, h.position.Y/0.5+20, 2, -3, h.shootType)
		}
	}

}

func (h *Hero) addShoot(game *game, x, y, speedX, speedY float64, shootType shootType) {
	switch shootType {
	case defaultShoot, threeShoots, ballShoot:
		h.shootFrame = int(ebiten.DefaultTPS / 1.2 * rand.Float64())
	case diagonalBallShoot:
		h.shootFrame = int(ebiten.DefaultTPS / 0.4 * rand.Float64())
	}

	shoot := NewShoot(x, y, shootType)
	shoot.byHero = true
	shoot.box.SpeedX = speedX
	shoot.box.SpeedY = speedY
	game.shoots = append(game.shoots, shoot)
}

func yThreshold() float64 {
	return []float64{75, 150, 200, 300}[rand.Intn(4)]
}

func xThreshold() float64 {
	return []float64{10, 10, 20, 30}[rand.Intn(4)]
}
