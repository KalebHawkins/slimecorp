package engine

import (
	"bytes"
	"embed"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed assets/tilesets/*.png
var TileSetAssetsFS embed.FS

//go:embed assets/static/*.png
var StaticImageFS embed.FS

func LoadImageFile(fs embed.FS, file string) *ebiten.Image {
	fullPath := "assets/" + file
	f, err := fs.ReadFile(fullPath)
	if err != nil {
		log.Fatalf("failed to load tileset %s: %v", fullPath, err)
	}

	img, _, err := image.Decode(bytes.NewReader(f))
	if err != nil {
		log.Fatalf("failed to decode tileset %s: %v", fullPath, err)
	}

	return ebiten.NewImageFromImage(img)
}
