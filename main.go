package main

import (
	"github.com/cclulu/ebitentest/scene"
	"github.com/hajimehoshi/ebiten"
)

func main() {
	err := ebiten.Run(scene.Update, 640, 480, 1, "hello world!")
	if err != nil {
		panic("error running ebiten")
	}
}
