package engine

import "github.com/hajimehoshi/ebiten/v2"

type CutSceneEngine struct {
	Actions        []CutSceneAction
	IsCutSceneOver bool
}

func (c *CutSceneEngine) Update() {
	if len(c.Actions) == 0 {
		return
	}

	done := c.Actions[0].Update()
	if done {
		c.Actions = c.Actions[1:]
	}

	if len(c.Actions) == 0 {
		c.IsCutSceneOver = true
	}
}

func (c *CutSceneEngine) Draw(dst *ebiten.Image) {
	if len(c.Actions) > 0 {
		c.Actions[0].Draw(dst)
	}
}
