package entity

import (
	"math"

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
	op.GeoM.Translate(o.topLeft().X, o.topLeft().Y)
	screen.DrawImage(o.Sprite, op)
}

func (o GameObject) Collides(o2 GameObject) bool {
	r1 := o.Size.Radius()
	r2 := o2.Size.Radius()

	d := o2.Position.Sub(o.Position).Length()

	var intersects bool
	switch {
	case d > r1+r2, d < math.Abs(r1-r2):
		intersects = false
	default:
		intersects = true
	}
	return intersects
}

func (o *GameObject) BounceOff(o2 GameObject) {
	//elisticity
	e := 0.8
	//normal vector
	n := o.Position.Sub(o2.Position)

	//normalize
	nlength := n.Length()
	nn := n.Normalize()

	overlap := o.Size.Radius() + o2.Size.Radius() - nlength
	if overlap > 0 {
		o.Position = o.Position.Add(nn.Scale(overlap))
	}

	//unit tangent vector
	t := geometry.Vector2{X: -nn.Y, Y: nn.X}

	//puck's speed along normal
	v1n := o.Velocity.Dot(nn)

	//mallet's speed along normal
	v2n := o2.Velocity.Dot(nn)

	//puck's speed along tangent (unchanged by collision)
	v1t := o.Velocity.Dot(t)

	//puck's new normal speed
	v1nPrime := v2n*(1.0+e) - e*v1n

	o.Velocity = nn.Scale(v1nPrime).Add(t.Scale(v1t))
}

func (o GameObject) topLeft() geometry.Vector2 {
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
