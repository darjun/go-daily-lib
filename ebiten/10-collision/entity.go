package main

type Entity interface {
	Width() int
	Height() int
	X() float64
	Y() float64
}
