package geometry

import "math"

type Vector2 struct {
	X, Y float64
}

func (v Vector2) Add(o Vector2) Vector2 {
	return Vector2{X: v.X + o.X, Y: v.Y + o.Y}
}

func (v Vector2) Sub(o Vector2) Vector2 {
	return Vector2{X: v.X - o.X, Y: v.Y - o.Y}
}

func (v Vector2) Scale(f float64) Vector2 {
	return Vector2{X: v.X * f, Y: v.Y * f}
}

func (v Vector2) Dot(o Vector2) float64 {
	return v.X*o.X + v.Y*o.Y
}

func (v Vector2) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v Vector2) Normalize() Vector2 {
	return v.Scale(1 / v.Length())
}
