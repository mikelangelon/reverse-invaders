package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"image/color"
	"time"
)

const (
	startText = `Press RETURN to start`
	roundText = `Round %d`
)

var mplusNormalFont font.Face

func init() {
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		panic(err)
	}

	const dpi = 72
	mplusNormalFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingVertical,
	})
	if err != nil {
		panic(err)
	}
}

type game struct {
	hero       *Hero
	aliens     Aliens
	shoots     []*Shoot
	explosions []*Explosion
	state      gameState
	images     images
	points     int
	round      int
	updateTick time.Time

	currentFrame int
}

type images struct {
	menu  *ebiten.Image
	alien []*ebiten.Image
	hero  *ebiten.Image
}

func (g *game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return width, height
}

func (g *game) Update() error {
	switch g.state {
	case gameStateMenu:
		if inpututil.IsKeyJustReleased(ebiten.KeyEnter) {
			g.state = gameStatePrePlaying
			g.init()
		}
		return nil
	case gameStatePrePlaying:
		if inpututil.IsKeyJustReleased(ebiten.KeyEnter) {
			g.state = gameStatePlaying
		}
		return nil
	case gameRestarts:
		now := time.Now()
		fmt.Printf("%v\n", now.Sub(g.updateTick))
		if now.Sub(g.updateTick) > 1*time.Second {
			g.init()
			g.state = gameStatePrePlaying
			return nil
		}
	}
	g.entitiesPlay()
	g.updateCollisions()
	g.cleanDeads()
	g.updateState()
	g.currentFrame++

	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	switch g.state {
	case gameStateMenu:
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(0, 0)
		op.GeoM.Scale(0.5, 0.5)
		screen.DrawImage(g.images.menu, op)
	case gameStatePrePlaying:
		c := color.NRGBA{0xee, 0xe4, 0xda, 0x59}
		text.Draw(screen, startText, mplusNormalFont, width/2-160, height/2, c)
		text.Draw(screen, fmt.Sprintf(roundText, g.round+1), mplusNormalFont, width/2-75, height/2-50, c)
		fallthrough
	case gameStatePlaying, gameRestarts:
		text.Draw(screen, fmt.Sprintf("Points: %d", g.points), mplusNormalFont, width-160, 50, color.White)
		g.hero.Draw(screen)
		for _, v := range g.aliens {
			v.Draw(g.currentFrame, screen)
		}
		for _, v := range g.shoots {
			v.Draw(screen)
		}
		for _, v := range g.explosions {
			v.Draw(screen)
		}
	}
}

func (g *game) entitiesPlay() {
	g.moveAliens()
	g.hero.Move(g)
	g.moveShoots()
	g.animateExplosions()
}

func (g *game) moveAliens() {
	for _, v := range g.aliens {
		if v.player {
			g.moveAlien(v)
		}
	}
	g.aliens.getNoPlayers().Move(g)
}
func (g *game) moveShoots() {
	for _, v := range g.shoots {
		v.box.Y += v.box.Speed
	}
}

func (g *game) animateExplosions() {
	for _, v := range g.explosions {
		v.Update()
	}
}

func (g *game) updateCollisions() {
	g.playerToAlienCollisions()
	g.alienToShootsCollisions()
	g.heroCollisions()
}
func (g *game) playerToAlienCollisions() {
	player := g.aliens.getPlayers()
	noPlayers := g.aliens.getNoPlayers()
	for _, v := range player {
		for _, z := range noPlayers {
			if AreColliding(v.box, z.box) {
				v.state = stateDead
				z.state = stateDead
				g.explosions = append(g.explosions, NewExplosion(v.box.XScaled(), v.box.YScaled(), 0.8))
			}
		}
	}
}

func (g *game) alienToShootsCollisions() {
	for _, v := range g.aliens {
		for _, z := range g.shoots {
			if !z.byHero {
				continue
			}
			if AreColliding(v.box, z.box) {
				v.state = stateDead
				z.state = stateDead
				g.explosions = append(g.explosions, NewExplosion(v.box.XScaled(), v.box.YScaled(), 0.3))
			}
		}
	}

}

func (g *game) heroCollisions() {
	if g.hero.state == stateDead {
		return
	}
	for _, v := range g.shoots {
		if v.byHero {
			continue
		}
		if v.box.CollidesTo(g.hero.position) {
			g.hero.state = stateDead
			v.state = stateDead
			g.explosions = append(g.explosions, NewExplosion(v.box.XScaled(), v.box.YScaled(), 0.8))
		}
	}
}

func (g *game) updateState() {
	if g.state == gameRestarts {
		return
	}
	if len(g.aliens.getPlayers()) == 0 {
		g.state = gameRestarts
		g.points = 0
		g.round = 0
		g.updateTick = time.Now()
	}
	if g.hero.state == stateDead {
		g.state = gameRestarts
		g.points += len(g.aliens)
		g.round++
		g.updateTick = time.Now()
	}
}
func (g *game) cleanDeads() {
	var alive Aliens
	for _, v := range g.aliens {
		if v.state == stateDead {
			continue
		}
		alive = append(alive, v)
	}
	g.aliens = alive

	var shoot []*Shoot
	for _, v := range g.shoots {
		if v.state == stateDead {
			continue
		}
		if v.box.YScaled() > height-80 || v.box.YScaled() < 10 {
			continue
		}
		shoot = append(shoot, v)
	}
	g.shoots = shoot

	var explosions []*Explosion
	for _, v := range g.explosions {
		if v.ToBeCleaned() {
			continue
		}
		explosions = append(explosions, v)
	}
	g.explosions = explosions
}
func (g *game) moveAlien(alien *Alien) {
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		alien.box.Y -= alien.box.Speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		alien.box.Y += alien.box.Speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		alien.box.X -= alien.box.Speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		alien.box.X += alien.box.Speed
	}
	if inpututil.IsKeyJustReleased(ebiten.KeySpace) {
		g.shoots = append(g.shoots, NewShoot(alien.box.X, alien.box.Y))
	}
}

func (g *game) init() {
	g.hero = &Hero{
		img: g.images.hero,
		position: Box{
			X:      100,
			Y:      500,
			With:   64,
			Height: 64,
			Scale:  1,
			Speed:  5,
		}}
	g.aliens = generateAliens(g.images.alien)
	g.shoots = []*Shoot{}
}

func AreColliding(b1, b2 Box) bool {
	return b1.CollidesTo(b2)
}
