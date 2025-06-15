package engine

import "github.com/hajimehoshi/ebiten/v2"

type Character interface {
	Update() error
	Draw(*ebiten.Image)
	SetPosition(*Vector2)
	GetPosition() *Vector2
}
