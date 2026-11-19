package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/safkhamezadeh/airhockey/ai"
	"github.com/safkhamezadeh/airhockey/config"
	"github.com/safkhamezadeh/airhockey/entity"
	"github.com/safkhamezadeh/airhockey/geometry"
)

type Game struct {
	puck     entity.Puck
	pMallet  entity.Mallet
	aiMallet entity.Mallet
}

func (g *Game) Update() error {
	g.updatePlayerMallet()
	g.updateAI()
	g.updatePuck()
	g.collisions()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.puck.Draw(screen)
	g.pMallet.Draw(screen)
	g.aiMallet.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func (g *Game) updatePlayerMallet() {
	x, y := ebiten.CursorPositionF()
	cursor := geometry.Vector2{X: x, Y: y}
	pressed := inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft)
	released := inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft)

	g.pMallet.UpdateDrag(cursor, pressed, released)
}

func (g *Game) updateAI() {
	bounds := config.Default().ScreenSize
	targetPos, targetVel := ai.Predict(g.puck, bounds, 8)
	ai.Move(&g.aiMallet, targetPos, targetVel, bounds)
}

func (g *Game) updatePuck() {
	g.puck.Drag()
	g.puck.Position = g.puck.Position.Add(g.puck.Velocity)
}

func (g *Game) collisions() {
	g.puck.BounceOffWalls(config.Default().ScreenSize)
	if g.puck.Collides(g.pMallet.GameObject) {
		g.puck.BounceOff(g.pMallet.GameObject)
	}
	if g.puck.Collides(g.aiMallet.GameObject) {
		g.puck.BounceOff(g.aiMallet.GameObject)
	}

}
