package engine

import "github.com/hajimehoshi/ebiten/v2"

type Adventurer struct {
	Position        *Vector2
	Sprite          *ebiten.Image
	FacingDirection float64
}

func NewAdventurer(position *Vector2, sprite *ebiten.Image) *Adventurer {
	a := &Adventurer{
		Position:        position,
		Sprite:          sprite,
		FacingDirection: 1.0,
	}

	return a
}

func (a *Adventurer) Update() error {
	return nil
}

func (a *Adventurer) Draw(dst *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(float64(a.FacingDirection), 1.0)
	op.GeoM.Translate(-float64(a.Sprite.Bounds().Dx())/2, -float64(a.Sprite.Bounds().Dy())/2)
	op.GeoM.Translate(a.Position.X, a.Position.Y)
	dst.DrawImage(a.Sprite, op)
}

func (a *Adventurer) GetPosition() *Vector2 {
	return a.Position
}

func (a *Adventurer) SetPosition(v *Vector2) {
	a.Position = v
}
