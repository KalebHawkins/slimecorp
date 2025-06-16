package main

import (
	"log"

	"github.com/KalebHawkins/slimecorp/engine"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	// ScreenWidth  = 640 // 40 16x16 Tiles
	// ScreenHeight = 480 // 30 16x16 Tiles
	ScreenWidth  = 320 // 20 16x16 Tiles
	ScreenHeight = 240 // 15 16x16 Tiles
)

func main() {
	ebiten.SetWindowSize(ScreenWidth*2, ScreenHeight*2)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("SlimeCorp: Rise of the Collective")

	if err := ebiten.RunGame(engine.NewGame(ScreenWidth, ScreenHeight)); err != nil {
		log.Fatal(err)
	}
}
