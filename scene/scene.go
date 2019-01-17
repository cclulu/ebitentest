package scene

import (
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	ScreenWidth  = 500
	ScreenHeight = 500
)

var (
	imageGameBG   *ebiten.Image
	imageWindows  *ebiten.Image
	imageGameover *ebiten.Image
	lightGray     ebiten.ColorM
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

func init() {

	background, err := os.Open("resources/dungeon.jpg")
	if err != nil {
		panic("on reading the file")
	}
	defer background.Close()

	img, _, err := image.Decode(background)
	if err != nil {
		panic(err)
	}

	imageWindows, err = ebiten.NewImage(ScreenWidth, ScreenHeight, ebiten.FilterDefault)
	if err != nil {
		spew.Dump(err)
	}

	imageGameBG, err = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	if err != nil {
		spew.Dump(err)
	}

	imageWindows.Fill(color.White)

	w, h := imageGameBG.Size()
	scaleW := ScreenWidth / float64(w)
	scaleH := ScreenHeight / float64(h)
	scale := scaleW
	if scale < scaleH {
		scale = scaleH
	}

	ebitenutil.DrawRect(imageWindows, float64(20), float64(20), float64(scaleW), float64(scaleH), color.RGBA{0, 0, 0, 0xc0})

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
	op.GeoM.Scale(scale, scale)
	op.GeoM.Translate(ScreenWidth/2, ScreenHeight/2)
	op.ColorM = lightGray
	op.Filter = ebiten.FilterLinear
	imageWindows.DrawImage(imageGameBG, op)
}
