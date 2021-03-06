package screenmanager

import (
	"fmt"

	"github.com/cclulu/ebitentest/sprites"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

func DebugScreen(screen *ebiten.Image) {
	x, y := ebiten.CursorPosition()
	ebitenutil.DebugPrint(screen, fmt.Sprintf("X: %d, Y: %d\n", x, y))

	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		ebitenutil.DebugPrint(screen, "\nYou are pressing the 'UP' button")
	}

	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		ebitenutil.DebugPrint(screen, "\nYou are pressing the 'Down' button")
	}

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		ebitenutil.DebugPrint(screen, "\nYou are pressing the 'Left' button")
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		ebitenutil.DebugPrint(screen, "\nYou are pressing the 'Right' button")
	}

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		ebitenutil.DebugPrint(screen, "\nYou are pressing the 'Left' mouse button")
	}

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		ebitenutil.DebugPrint(screen, "\nYou are pressing the 'Right' mouse button")
	}

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonMiddle) {
		ebitenutil.DebugPrint(screen, "\nYou are pressing the 'Middle' mouse button")
	}
}

func DrawBG(screen *ebiten.Image) {
	screen.DrawImage(sprites.LoadImages(), &ebiten.DrawImageOptions{})
}
