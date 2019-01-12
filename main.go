package main

import (
	"github.com/cclulu/ebitentest/movement"
	"github.com/hajimehoshi/ebiten"
)

func main() {

	err := ebiten.Run(movement.Update, 320, 240, 2, "hello world!")
	if err != nil {
		panic("error running ebiten")
	}
}
