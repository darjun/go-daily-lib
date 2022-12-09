package main

import (
	"bytes"
	"github.com/darjun/ebiten/11-pack/resources"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	_ "image/png"
	"log"
)

type Alien struct {
	GameObject
	image       *ebiten.Image
	speedFactor float64
}

func NewAlien(cfg *Config) *Alien {
	img, _, err := ebitenutil.NewImageFromReader(bytes.NewReader(resources.AlienPng))
	if err != nil {
		log.Fatal(err)
	}

	width, height := img.Size()
	return &Alien{
		GameObject: GameObject{
			width:  width,
			height: height,
			x:      0,
			y:      0,
		},
		image:       img,
		speedFactor: cfg.AlienSpeedFactor,
	}
}

func (alien *Alien) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(alien.x, alien.y)
	screen.DrawImage(alien.image, op)
}

func (alien *Alien) outOfScreen(cfg *Config) bool {
	return alien.y > float64(cfg.ScreenHeight)
}
