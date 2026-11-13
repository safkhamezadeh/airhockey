package entity

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/safkhamezadeh/airhockey/geometry"
)

type GameObject struct {
	Position geometry.Vector2
	Size     geometry.Size
	Sprite   *ebiten.Image
	Velocity geometry.Vector2
}

func (o GameObject) Draw(screen *ebiten.Image) {
	bounds := o.Sprite.Bounds()
	scaleX := o.Size.Width / float64(bounds.Dx())
	scaleY := o.Size.Height / float64(bounds.Dy())

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(scaleX, scaleY)
	op.GeoM.Translate(o.middle().X, o.middle().Y)
	screen.DrawImage(o.Sprite, op)
}

func (o GameObject) middle() geometry.Vector2 {
	return geometry.Vector2{X: o.Position.X - o.Size.Width/2, Y: o.Position.Y - o.Size.Height/2}
}

func NewGameObject(spawnPosition geometry.Vector2, sprite *ebiten.Image, scale float64) GameObject {
	bounds := sprite.Bounds()
	return GameObject{
		Position: spawnPosition,
		Size:     geometry.Size{Width: float64(bounds.Dx()) * scale, Height: float64(bounds.Dy()) * scale},
		Sprite:   sprite,
	}
}
