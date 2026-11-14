package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/safkhamezadeh/airhockey/config"
	"github.com/safkhamezadeh/airhockey/entity"
	"github.com/safkhamezadeh/airhockey/geometry"
)

type Game struct {
	puck    entity.Puck
	pMallet entity.Mallet
}

func (g *Game) Update() error {
	g.updatePlayerMallet()
	g.updatePuck()
	g.collisions()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.puck.Draw(screen)
	g.pMallet.Draw(screen)
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

func (g *Game) updatePuck() {
	p := g.puck
	vel := p.Velocity
	newX := vel.X + p.Position.X
	newY := vel.Y + p.Position.Y
	g.puck.Position.X = newX
	g.puck.Position.Y = newY
}

func (g *Game) collisions() {
	g.puck.BounceOffWalls(config.Default().ScreenSize)
	if g.puck.Collides(g.pMallet.GameObject) {
		g.puck.BounceOff(g.pMallet.GameObject)
	}
	//border collisions
	//if puck hits wall, velx * -1
	//if mallet hits wall, velx = 0, xpos = wall

	//entity collisions
}
