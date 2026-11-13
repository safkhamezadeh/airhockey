package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/safkhamezadeh/airhockey/assets"
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

	g.puck.Position.X += 1
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	scale := config.Default().Puck.Scale

	p := g.puck.Scale(scale)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(scale, scale)
	op.GeoM.Translate(g.puck.Position.X-p.Width/2, g.puck.Position.Y-p.Height/2)
	screen.DrawImage(g.puck.Sprite, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	puckSprite := assets.LoadAsset(assets.Puck)
	puckBounds := puckSprite.Bounds()

	malletSprite := assets.LoadAsset(assets.Mallet)
	malletBounds := malletSprite.Bounds()

	g := &Game{
		puck: entity.Puck{
			GameObject: entity.GameObject{
				Position: config.Default().Puck.SpawnPosition,
				Size:     geometry.Size{Width: float64(puckBounds.Dx()), Height: float64(puckBounds.Dy())},
				Sprite:   puckSprite,
			},
		},
	}
	ebiten.SetWindowSize(int(config.Default().ScreenSize.Width), int(config.Default().ScreenSize.Height))
	ebiten.SetWindowTitle("airhockey")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
