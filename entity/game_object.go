package entity

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/safkhamezadeh/airhockey/geometry"
)

type GameObject struct {
	Position geometry.Vector2
	Size     geometry.Size
	Sprite   *ebiten.Image
}

func (o GameObject) Scale(s float64) geometry.Size {
	return o.Size.Scale(s)
}

func (o GameObject) Mid() geometry.Vector2 {
	return geometry.Vector2{X: o.Size.Width / 2, Y: o.Size.Height / 2}
}
