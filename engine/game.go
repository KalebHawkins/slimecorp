package engine

import (
	"bytes"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/image/font/gofont/goregular"
)

type Config struct {
	ScreenWidth    int
	ScreenHeight   int
	FontFace       *text.GoTextFace
	FontFaceSource *text.GoTextFaceSource
}

type Game struct {
	CurrentGameState GameState
	Config           *Config
	OffScreen        *ebiten.Image
}

func NewGame(scrWidth, scrHeight int) *Game {
	g := &Game{
		CurrentGameState: nil,
		Config: &Config{
			ScreenWidth:  scrWidth,
			ScreenHeight: scrHeight,
		},
		OffScreen: ebiten.NewImage(scrWidth, scrHeight),
	}

	s, err := text.NewGoTextFaceSource(bytes.NewReader(goregular.TTF))
	if err != nil {
		log.Fatal(err)
	}

	g.Config.FontFaceSource = s
	g.Config.FontFace = &text.GoTextFace{
		Source: s,
		Size:   14,
	}

	g.CurrentGameState = NewCombatState(g.Config)

	return g
}

func (g *Game) Update() error {
	g.CurrentGameState.Update()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.OffScreen.Clear()
	g.CurrentGameState.Draw(g.OffScreen)

	screenSize := screen.Bounds().Size()
	op := &ebiten.DrawImageOptions{}
	scaleX := float64(screenSize.X) / float64(g.Config.ScreenWidth)
	scaleY := float64(screenSize.Y) / float64(g.Config.ScreenHeight)
	op.GeoM.Scale(scaleX, scaleY)

	screen.DrawImage(g.OffScreen, op)

}

func (g *Game) Layout(outWidth, outHeight int) (int, int) {
	return g.Config.ScreenWidth, g.Config.ScreenHeight
}
