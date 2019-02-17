package initialize

import (
	"image"
	"image/color"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	ScreenWidth  = 640
	ScreenHeight = 480
)

var (
	imageGameBG    *ebiten.Image
	imageWindows   *ebiten.Image
	imageGameover  *ebiten.Image
	imageCharacter *ebiten.Image
	lightGray      ebiten.ColorM
)

func LoadImages() *ebiten.Image {
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
	return imageWindows
	// screen.DrawImage(imageWindows, &ebiten.DrawImageOptions{})
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

	characterSpriteSheet, err := os.Open("resources/MainGuySpriteSheet.png")
	if err != nil {
		panic("on reading the file")
	}
	defer characterSpriteSheet.Close()

	img, _, err = image.Decode(characterSpriteSheet)
	if err != nil {
		panic(err)
	}

	imageCharacter, err = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	if err != nil {
		spew.Dump(err)
	}
}
