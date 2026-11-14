package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/safkhamezadeh/airhockey/assets"
	"github.com/safkhamezadeh/airhockey/config"
	"github.com/safkhamezadeh/airhockey/entity"
)

func main() {
	puckSprite := assets.LoadAsset(assets.Puck)

	malletSprite := assets.LoadAsset(assets.Mallet)

	g := &Game{
		puck: entity.Puck{
			GameObject: entity.NewGameObject(config.Default().Puck.SpawnPosition, puckSprite, config.Default().Puck.Scale),
		},
		pMallet: entity.Mallet{
			GameObject: entity.NewGameObject(config.Default().Mallet.SpawnPosition, malletSprite, config.Default().Mallet.Scale),
			IsDragging: false,
		},
	}
	g.puck.Velocity.X = 2
	ebiten.SetWindowSize(int(config.Default().ScreenSize.Width), int(config.Default().ScreenSize.Height))
	ebiten.SetWindowTitle("airhockey")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
