package main

import (
	"github.com/cclulu/ebitentest/movement"
	"github.com/hajimehoshi/ebiten"
)

func main() {
	err := ebiten.Run(movement.Update, 640, 480, 1, "hello world!")
	if err != nil {
		panic("error running ebiten")
	}
}
