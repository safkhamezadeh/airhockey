package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/safkhamezadeh/airhockey/assets"
	"github.com/safkhamezadeh/airhockey/config"
	"github.com/safkhamezadeh/airhockey/entity"
)

type Game struct {
	puck    entity.Puck
	pMallet entity.Mallet
}

func (g *Game) Update() error {

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.puck.Draw(screen)
	g.pMallet.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	puckSprite := assets.LoadAsset(assets.Puck)

	malletSprite := assets.LoadAsset(assets.Mallet)
	//malletBounds := malletSprite.Bounds()

	g := &Game{
		puck: entity.Puck{
			GameObject: entity.NewGameObject(config.Default().Puck.SpawnPosition, puckSprite, config.Default().Puck.Scale),
		},
		pMallet: entity.Mallet{
			GameObject: entity.NewGameObject(config.Default().Mallet.SpawnPosition, malletSprite, config.Default().Mallet.Scale),
			IsDragging: false,
		},
	}
	ebiten.SetWindowSize(int(config.Default().ScreenSize.Width), int(config.Default().ScreenSize.Height))
	ebiten.SetWindowTitle("airhockey")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
