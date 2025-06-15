package engine

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

type CombatState struct {
	scrWidth, scrHeight int
	BackGroundImages    []*ebiten.Image
	TileMap             *TileMap
}

func NewCombatState(scrWidth, scrHeight int) *CombatState {
	cs := &CombatState{
		TileMap:   NewTileMap(40, 30, 16, JungleMap, LoadImageFile(TileSetAssetsFS, "tilesets/jungle.png")),
		scrWidth:  scrWidth,
		scrHeight: scrHeight,
	}

	for i := 1; i <= 5; i++ {
		imgName := fmt.Sprintf("plx-%d.png", i)
		cs.BackGroundImages = append(cs.BackGroundImages, LoadImageFile(StaticImageFS, "backgrounds/"+imgName))
	}

	return cs
}

func (cs *CombatState) Update() error {
	return nil
}

func (cs *CombatState) Draw(dst *ebiten.Image) {
	for _, bgImg := range cs.BackGroundImages {
		imgsize := bgImg.Bounds().Size()
		scaleX := float64(cs.scrWidth) / float64(imgsize.X)
		scaleY := float64(cs.scrHeight) / float64(imgsize.Y)

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(scaleX, scaleY)
		dst.DrawImage(bgImg, op)
	}
	cs.TileMap.Draw(dst)
}
