package entity

import (
	"math"

	"github.com/safkhamezadeh/airhockey/geometry"
)

type Mallet struct {
	GameObject
	IsDragging bool
}

func (m Mallet) InBounds(point geometry.Vector2) bool {
	x := point.X
	y := point.Y

	r := m.Size.Width / 2
	mMidX := m.Position.X
	mMidY := m.Position.Y

	ds := math.Pow((x-mMidX), 2) + math.Pow((y-mMidY), 2)

	//inside or on boundary
	if ds <= math.Pow(r, 2) {
		return true
	} else {
		return false
	}
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
