package entity

import "github.com/safkhamezadeh/airhockey/geometry"

type Puck struct {
	GameObject
}

func (p *Puck) BounceOffWalls(bounds geometry.Size) {
	r := p.Size.Width / 2
	if p.Position.X-r <= 0 || p.Position.X+r >= bounds.Width {
		p.Velocity.X *= -1

	}
}
