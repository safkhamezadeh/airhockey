package entity

import "github.com/safkhamezadeh/airhockey/geometry"

type Puck struct {
	GameObject
}

func (p *Puck) BounceOffWalls(bounds geometry.Size) {
	r := p.Size.Width / 2
	const wallRestitution = 0.95

	switch {
	case p.Position.X-r <= 0:
		p.Velocity.X *= -wallRestitution
		p.Position.X = 0 + r
	case p.Position.X+r >= bounds.Width:
		p.Velocity.X *= -wallRestitution
		p.Position.X = bounds.Width - r
	}

	//temp
	if p.Position.Y-r <= 0 || p.Position.Y+r >= bounds.Height {
		p.Velocity.Y *= -1
	}
}
