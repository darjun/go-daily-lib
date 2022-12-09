package main

import (
	"bytes"
	"encoding/json"
	"github.com/darjun/ebiten/11-pack/resources"
	"image/color"
	"log"
)

type Config struct {
	ScreenWidth       int        `json:"screenWidth"`
	ScreenHeight      int        `json:"screenHeight"`
	Title             string     `json:"title"`
	BgColor           color.RGBA `json:"bgColor"`
	ShipSpeedFactor   float64    `json:"shipSpeedFactor"`
	BulletWidth       int        `json:"bulletWidth"`
	BulletHeight      int        `json:"bulletHeight"`
	BulletSpeedFactor float64    `json:"bulletSpeedFactor"`
	BulletColor       color.RGBA `json:"bulletColor"`
	MaxBulletNum      int        `json:"maxBulletNum"`
	BulletInterval    int64      `json:"bulletInterval"`
	AlienSpeedFactor  float64    `json:"alienSpeedFactor"`
	FontSize          int        `json:"fontSize"`
	TitleFontSize     int        `json:"titleFontSize"`
	SmallFontSize     int        `json:"smallFontSize"`
}

func loadConfig() *Config {
	var cfg Config
	if err := json.NewDecoder(bytes.NewReader(resources.ConfigJson)).Decode(&cfg); err != nil {
		log.Fatalf("json.Decode failed: %v\n", err)
	}

	return &cfg
}
