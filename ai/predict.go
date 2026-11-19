package ai

import (
	"math"

	"github.com/safkhamezadeh/airhockey/entity"
	"github.com/safkhamezadeh/airhockey/geometry"
)

// this function predicts where the puck will be in n frames
func Predict(puck entity.Puck, bounds geometry.Size, frames int) (geometry.Vector2, geometry.Vector2) {
	for range frames {
		puck.Position = puck.Position.Add(puck.Velocity)
		puck.BounceOffWalls(bounds)
		if puck.Position.Y+puck.Size.Radius() <= 5 {
			break
		}
	}
	return puck.Position, puck.Velocity
}

const maxVelocity = 10.0

func Move(mallet *entity.Mallet, targetPos geometry.Vector2, targetVel geometry.Vector2, bounds geometry.Size) {
	zone := ownHalf(bounds)
	isOwnHalf := isOwnHalf(targetPos, zone)
	fast := fast(targetVel)
	towardsGoal := towardsGoal(targetVel)

	switch {
	case fast && towardsGoal:
		defend(mallet, targetPos, zone)
	case isOwnHalf && !fast:
		attack(mallet, targetPos)
	default:
		follow(mallet, targetPos, zone)
	}
}

func defend(mallet *entity.Mallet, target geometry.Vector2, zone geometry.Size) {
	// defend in zone, if out of zone, go in zone towards target with (increasing / decreasing) velocity
	defenseLine := zone.Height / 2.0
	desired := target
	if desired.Y > defenseLine {
		desired.Y = defenseLine
	}
	moveTowards(mallet, desired, maxVelocity)
}

func attack(mallet *entity.Mallet, target geometry.Vector2) {
	//try to hit the puck with increasing velocity until max
	moveTowards(mallet, target, maxVelocity)
}

func follow(mallet *entity.Mallet, target geometry.Vector2, zone geometry.Size) {
	//follow puck on line
	line := zone.Height / 2
	moveTowards(mallet, geometry.Vector2{X: target.X, Y: line}, maxVelocity)
}

// moveTowards sets the mallet's velocity to head straight at target, capped at maxSpeed.
func moveTowards(mallet *entity.Mallet, target geometry.Vector2, maxSpeed float64) {
	direction := target.Sub(mallet.Position)
	distance := direction.Length()
	if distance == 0 {
		mallet.Velocity = geometry.Vector2{}
		return
	}
	mallet.Velocity = direction.Normalize().Scale(math.Min(distance, maxSpeed))
	mallet.Position = mallet.Position.Add(mallet.Velocity)
}

func ownHalf(bounds geometry.Size) geometry.Size {
	return geometry.Size{Width: bounds.Width, Height: bounds.Height / 2}
}

func isOwnHalf(p geometry.Vector2, bounds geometry.Size) bool {
	return p.Y <= bounds.Height
}

func towardsGoal(dir geometry.Vector2) bool {
	return dir.X <= 0
}

func fast(vel geometry.Vector2) bool {
	fast := 20.0
	return vel.Length() >= fast
}
