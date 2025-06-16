package engine

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type CutSceneAction interface {
	Update() bool
	Draw(*ebiten.Image)
}

type MoveAction struct {
	Character     Character
	StartPosition *Vector2
	EndPosition   *Vector2
	Speed         float64
	Direction     float64
	done          bool
}

func (ma *MoveAction) Update() bool {
	if ma.done {
		return true
	}

	ma.StartPosition = ma.Character.GetPosition()
	ma.Direction = 1.0
	if ma.StartPosition.X > ma.EndPosition.X {
		ma.Direction = -1.0
	}

	nextX := ma.StartPosition.X + ma.Speed*ma.Direction
	if ma.Direction < 0 && nextX <= ma.EndPosition.X || ma.Direction > 0 && nextX >= ma.EndPosition.X {
		nextX = ma.EndPosition.X
		ma.done = true
	}

	if adventurer, ok := ma.Character.(*Adventurer); ok {
		adventurer.FacingDirection = ma.Direction
	}

	ma.Character.SetPosition(&Vector2{nextX, ma.StartPosition.Y})
	return ma.done
}

func (ma *MoveAction) Draw(dst *ebiten.Image) {
	ma.Character.Draw(dst)
}

type WaitAction struct {
	Character Character
	Frames    int
	WaitTime  int
}

func (wa *WaitAction) Update() bool {
	wa.Frames++

	elapsed := 1.0 / ebiten.ActualTPS() * float64(wa.Frames)
	return elapsed >= float64(wa.WaitTime)
}

func (wa *WaitAction) Draw(dst *ebiten.Image) {
	wa.Character.Draw(dst)
}

type TextAction struct {
	Position       *Vector2
	Text           string
	FontFace       *text.GoTextFace
	FontFaceSource *text.GoTextFaceSource
	WaitTime       int
	Frames         int
}

func (ta *TextAction) Update() bool {
	ta.Frames++
	elapsed := 1.0 / ebiten.ActualTPS() * float64(ta.Frames)
	return elapsed >= float64(ta.WaitTime)
}

func (ta *TextAction) Draw(dst *ebiten.Image) {
	top := &text.DrawOptions{}
	top.LineSpacing = 12
	top.PrimaryAlign = text.AlignCenter

	top.ColorScale.ScaleWithColor(color.Black)
	top.GeoM.Translate(ta.Position.X, ta.Position.Y)
	text.Draw(dst, ta.Text, ta.FontFace, top)
}
