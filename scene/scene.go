package scene

import (
	_ "image/jpeg"

	"github.com/cclulu/ebitentest/screenmanager"
	"github.com/hajimehoshi/ebiten"
)

func Update(screen *ebiten.Image) error {
	screenmanager.DrawBG(screen)
	screenmanager.DebugScreen(screen)
	return nil
}
