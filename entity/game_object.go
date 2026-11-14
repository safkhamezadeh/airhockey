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

	d := math.Sqrt(math.Pow(o2.Position.X-o.Position.X, 2) + math.Pow(o2.Position.Y-o.Position.Y, 2))

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
	n := geometry.Vector2{X: o.Position.X - o2.Position.X, Y: o.Position.Y - o2.Position.Y}

	//normalize
	nlength := math.Sqrt(math.Pow(n.X, 2) + math.Pow(n.Y, 2))
	nn := geometry.Vector2{X: n.X / nlength, Y: n.Y / nlength}

	overlap := o.Size.Radius() + o2.Size.Radius() - nlength
	if overlap > 0 {
		o.Position.X += nn.X * overlap
		o.Position.Y += nn.Y * overlap
	}

	//unit tangent vector
	t := geometry.Vector2{X: -nn.Y, Y: nn.X}

	//puck's speed along normal
	v1n := o.Velocity.X*nn.X + o.Velocity.Y*nn.Y

	//mallet's speed along normal
	v2n := o2.Velocity.X*nn.X + o2.Velocity.Y*nn.Y

	//puck's speed along tangent (unchanged by collision)
	v1t := o.Velocity.X*t.X + o.Velocity.Y*t.Y

	//puck's new normal speed
	v1nPrime := v2n*(1.0+e) - e*v1n

	o.Velocity.X = v1nPrime*nn.X + v1t*t.X
	o.Velocity.Y = v1nPrime*nn.Y + v1t*t.Y
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
