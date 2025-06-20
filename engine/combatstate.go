package engine

import (
	"fmt"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type CombatState struct {
	Config                  *Config
	BackGroundImages        []*ebiten.Image
	TileMap                 *TileMap
	Player                  *Player
	Gravity                 float64
	PlayerAcceleration      *Vector2
	PlayerVelocity          *Vector2
	MaxPlayerVelocity       *Vector2
	PlayerJumpImpulse       float64
	PlayerAccelerationSpeed float64
	VelocityDampening       float64
	IsPlayerOnGround        bool
	CutSceneEngine          *CutSceneEngine
	Adventurer              *Adventurer
}

func NewCombatState(cfg *Config) *CombatState {
	cs := &CombatState{
		TileMap:                 NewTileMap(20, 15, 16, JungleMap, LoadImageFile(TileSetAssetsFS, "tilesets/jungle.png")),
		Config:                  cfg,
		Player:                  NewPlayer(&Vector2{float64(cfg.ScreenWidth) / 4, 170}, ebiten.NewImage(16, 16)),
		BackGroundImages:        []*ebiten.Image{},
		Gravity:                 0.918,
		PlayerJumpImpulse:       10.0,
		PlayerAcceleration:      &Vector2{},
		PlayerVelocity:          &Vector2{},
		MaxPlayerVelocity:       &Vector2{8, 10},
		PlayerAccelerationSpeed: 1.0,
		IsPlayerOnGround:        false,
		VelocityDampening:       0.9,
		Adventurer:              NewAdventurer(&Vector2{float64(cfg.ScreenWidth) + 40, 162}, LoadImageFile(StaticImageFS, "static/adventurer_idle.png")),
	}

	for i := 1; i <= 5; i++ {
		imgName := fmt.Sprintf("plx-%d.png", i)
		cs.BackGroundImages = append(cs.BackGroundImages, LoadImageFile(StaticImageFS, "static/"+imgName))
	}

	cs.CutSceneEngine = &CutSceneEngine{
		Actions: []CutSceneAction{
			&MoveAction{
				Character:     cs.Adventurer,
				StartPosition: cs.Adventurer.Position,
				EndPosition:   &Vector2{float64(cfg.ScreenWidth) / 1.5, 162},
				Speed:         1.0,
				done:          false,
			},
			&TextAction{
				Position:       &Vector2{float64(cfg.ScreenWidth) / 1.5, 120},
				Text:           "Oh look another slime...\nHehe, TIME TO DIE!",
				FontFace:       cs.Config.FontFace,
				FontFaceSource: cs.Config.FontFaceSource,
				WaitTime:       3.0,
			},
		},
	}

	return cs
}

func (cs *CombatState) Update() error {
	if cs.CutSceneEngine != nil && !cs.CutSceneEngine.IsCutSceneOver {
		cs.CutSceneEngine.Update()
		return nil
	}

	cs.PlayerAcceleration = &Vector2{}

	// Handle Gravity
	if !cs.IsPlayerOnGround {
		cs.PlayerAcceleration.Y += cs.Gravity
	}

	// Handle Player Input
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		cs.PlayerAcceleration.Add(&Vector2{-cs.PlayerAccelerationSpeed, 0})
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		cs.PlayerAcceleration.Add(&Vector2{cs.PlayerAccelerationSpeed, 0})
	}
	if ebiten.IsKeyPressed(ebiten.KeySpace) && cs.IsPlayerOnGround {
		cs.PlayerAcceleration.Add(&Vector2{0, -cs.PlayerJumpImpulse})
		cs.IsPlayerOnGround = false
	}

	// Add velocity to player.
	cs.PlayerVelocity.Add(cs.PlayerAcceleration)

	// Clamp the X and Y velocity so that the player can't exceed a specific speed.
	cs.PlayerVelocity.X = Clamp(cs.PlayerVelocity.X, -cs.MaxPlayerVelocity.X, cs.MaxPlayerVelocity.X)
	cs.PlayerVelocity.Y = Clamp(cs.PlayerVelocity.Y, -cs.MaxPlayerVelocity.Y, cs.MaxPlayerVelocity.Y)

	// Handle Player Ground Collision.
	cs.CheckCollisions()

	// Update the player's position based on velocity.
	// CheckCollision performs some operations on velocity based on the tiles the player is touching.
	cs.Player.Position.Add(cs.PlayerVelocity)
	cs.Player.Position.X = math.Floor(cs.Player.Position.X)
	cs.Player.Position.Y = math.Floor(cs.Player.Position.Y)

	// Dampening for velocity to slow the character down.
	cs.PlayerVelocity.X *= cs.VelocityDampening
	if math.Abs(cs.PlayerVelocity.X) < 0.1 {
		cs.PlayerVelocity.X = 0
	}

	// Don't let the player leave the window.
	cs.Player.Position.X = Clamp(cs.Player.Position.X, float64(cs.Player.Sprite.Bounds().Dx()/2), float64(cs.Config.ScreenWidth)-float64(cs.Player.Sprite.Bounds().Dy())/2)

	// fmt.Printf("[DEBUG] Player Acceleration: (%.2f, %.2f), Player Velocity: (%.2f, %.2f), Player Position: (%f, %f)\n",
	// 	cs.PlayerAcceleration.X, cs.PlayerAcceleration.Y, cs.PlayerVelocity.X, cs.PlayerVelocity.Y, cs.Player.Position.X, cs.Player.Position.Y)

	return nil
}

func (cs *CombatState) Draw(dst *ebiten.Image) {
	for _, bgImg := range cs.BackGroundImages {
		imgsize := bgImg.Bounds().Size()
		scaleX := float64(cs.Config.ScreenWidth) / float64(imgsize.X)
		scaleY := float64(cs.Config.ScreenHeight) / float64(imgsize.Y)

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(scaleX, scaleY)
		dst.DrawImage(bgImg, op)
	}

	cs.TileMap.Draw(dst)

	cs.Adventurer.Draw(dst)
	cs.Player.Draw(dst)

	if cs.CutSceneEngine != nil && !cs.CutSceneEngine.IsCutSceneOver {
		cs.CutSceneEngine.Draw(dst)
		return
	}
}

func (cs *CombatState) CheckCollisions() {
	currentPlayerPosition := cs.Player.Position
	nextPlayerPosition := &Vector2{
		X: currentPlayerPosition.X + cs.PlayerAcceleration.X,
		Y: currentPlayerPosition.Y + cs.PlayerAcceleration.Y,
	}

	tX := math.Floor(nextPlayerPosition.X + float64(cs.Player.Sprite.Bounds().Dx())/2)
	tY := math.Floor(nextPlayerPosition.Y + float64(cs.Player.Sprite.Bounds().Dy()/2))

	tileId := cs.TileMap.TileAt(int(tX), int(tY))
	// fmt.Printf("[DEBUG] Tile at (%.2f, %.2f) = TILE: %d\n", tX, tY, tileId)

	switch tileId {
	case OUT_OF_MAP:
		cs.Player.Position.X = float64(cs.Config.ScreenWidth) / 2
		cs.Player.Position.Y = 0
	case PLATFORM_TILE:
		cs.PlayerVelocity.Y = Clamp(cs.PlayerVelocity.Y, -cs.MaxPlayerVelocity.Y, 0)
		cs.IsPlayerOnGround = true
	default:
	}
}
