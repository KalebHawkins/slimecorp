package engine

import "github.com/hajimehoshi/ebiten/v2"

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

	ma.Character.SetPosition(&Vector2{nextX, ma.StartPosition.Y})
	return ma.done
}

func (ma *MoveAction) Draw(dst *ebiten.Image) {
	ma.Character.Draw(dst)
}
