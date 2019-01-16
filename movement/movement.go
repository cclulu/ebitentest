package movement

import (
	"github.com/cclulu/ebitentest/scene"
	"github.com/hajimehoshi/ebiten"
)

func Update(screen *ebiten.Image) error {
	scene.DebugScreen(screen)
	return nil
}
