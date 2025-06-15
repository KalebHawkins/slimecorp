package engine

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type TileMap struct {
	Tiles    []int
	Width    int
	Height   int
	TileSize int
	Tileset  *ebiten.Image
}

func NewTileMap(width, height, tileSize int, tiles []int, tileset *ebiten.Image) *TileMap {
	tm := &TileMap{
		Width:    width,
		Height:   height,
		TileSize: tileSize,
		Tiles:    tiles,
		Tileset:  tileset,
	}

	return tm
}

func (tm *TileMap) Draw(dst *ebiten.Image) {
	tilesPerRow := tm.Tileset.Bounds().Dx() / tm.TileSize

	for y := 0; y < tm.Height; y++ {
		for x := 0; x < tm.Width; x++ {
			tileId := tm.Tiles[y*tm.Width+x]
			if tileId < 0 {
				continue
			}

			sx := tileId % tilesPerRow * tm.TileSize
			sy := tileId / tilesPerRow * tm.TileSize
			src := tm.Tileset.SubImage(image.Rect(sx, sy, sx+tm.TileSize, sy+tm.TileSize)).(*ebiten.Image)

			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x*tm.TileSize), float64(y*tm.TileSize))
			dst.DrawImage(src, op)
		}
	}
}
