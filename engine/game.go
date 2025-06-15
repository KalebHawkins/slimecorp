package engine

import "github.com/hajimehoshi/ebiten/v2"

type Game struct {
	ScreenWidth      int
	ScreenHeight     int
	CurrentGameState GameState
}

func NewGame(scrWidth, scrHeight int) *Game {
	g := &Game{
		ScreenWidth:      scrWidth,
		ScreenHeight:     scrHeight,
		CurrentGameState: NewCombatState(scrWidth, scrHeight),
	}

	return g
}

func (g *Game) Update() error {
	g.CurrentGameState.Update()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.CurrentGameState.Draw(screen)
}

func (g *Game) Layout(outWidth, outHeight int) (int, int) {
	return outWidth, outHeight
}
