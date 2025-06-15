package engine

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	Position *Vector2
	Sprite   *ebiten.Image
}

func NewPlayer(position *Vector2, sprite *ebiten.Image) *Player {
	p := &Player{
		Position: position,
		Sprite:   sprite,
	}

	sprite.Fill(color.RGBA{0, 0, 255, 255})
	return p
}

func (p *Player) Update() error {
	return nil
}

func (p *Player) Draw(dst *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(p.Sprite.Bounds().Dx())/2, -float64(p.Sprite.Bounds().Dy())/2)
	op.GeoM.Translate(p.Position.X, p.Position.Y)

	dst.DrawImage(p.Sprite, op)
}

func (a *Player) GetPosition() *Vector2 {
	return a.Position
}

func (a *Player) SetPosition(v *Vector2) {
	a.Position = v
}
