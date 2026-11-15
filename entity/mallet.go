package entity

import (
	"github.com/safkhamezadeh/airhockey/geometry"
)

type Mallet struct {
	GameObject
	IsDragging bool
}

func (m Mallet) InBounds(point geometry.Vector2) bool {
	r := m.Size.Width / 2

	//inside or on boundary
	return point.Sub(m.Position).Length() <= r
}

func (m *Mallet) UpdateDrag(cursor geometry.Vector2, pressed, released bool) {
	switch {
	case pressed && m.InBounds(cursor):
		m.IsDragging = true
	case released:
		m.IsDragging = false
		m.Velocity = geometry.Vector2{}
	case m.IsDragging:
		m.Velocity = geometry.Vector2{X: cursor.X - m.Position.X, Y: cursor.Y - m.Position.Y}
		m.Position = cursor
	}
}
