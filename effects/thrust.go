package effects

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const ENTROPY = 0.1

type ThrustParticle struct {
	RotationalDiff rl.Vector2
	Position       rl.Vector2
	Velocity       rl.Vector2
	Direction      float32
	StoppedX       bool
	StoppedY       bool
}

func Initalize(rotationalDiff rl.Vector2, position rl.Vector2, direction float32, magnitude float32) *ThrustParticle {

	rad := direction * (math.Pi / 180.0) // Convert direction to radians
	initVel := rl.Vector2{
		X: float32(math.Cos(float64(rad))) * float32(magnitude),
		Y: float32(math.Sin(float64(rad))) * float32(magnitude),
	}

	return &ThrustParticle{
		RotationalDiff: rotationalDiff,
		Position:       position,
		Direction:      direction,
		Velocity:       initVel,
	}
}

func (particle *ThrustParticle) Update() {
	particle.applyForce()
	particle.draw()
}

func (p *ThrustParticle) applyForce() {

	// Calculate the acceleration components
	rad := (p.Direction - 180) * (math.Pi / 180.0) // Convert direction to radians
	acceleration := rl.Vector2{
		X: float32(math.Cos(float64(rad))) * ENTROPY,
		Y: float32(math.Sin(float64(rad))) * ENTROPY,
	}

	// fmt.Println("--")
	// fmt.Printf("vel: %f, %f\n", p.Velocity.X, p.Velocity.Y)

	if math.Abs(float64(p.Velocity.X))-math.Abs(float64(acceleration.X)) < 0 {
		p.StoppedX = true
	}

	if !p.StoppedX {
		p.Velocity.X += acceleration.X
		p.Position.X += p.Velocity.X
	}

	if math.Abs(float64(p.Velocity.Y))-math.Abs(float64(acceleration.Y)) < 0 {
		p.StoppedY = true
	}

	if !p.StoppedY {
		p.Velocity.Y += acceleration.Y
		p.Position.Y += p.Velocity.Y
	}

}

func (p *ThrustParticle) draw() {

	// Save the current transformation matrix
	rl.PushMatrix()

	// Translate to the center of the ship
	rl.Translatef(p.Position.X-p.RotationalDiff.X, p.Position.Y-p.RotationalDiff.Y, 0)

	// Apply the rotation
	rl.Rotatef(p.Direction, 0, 0, 1)

	// Translate back to the original position
	rl.Translatef(-p.Position.X+p.RotationalDiff.X, -p.Position.Y+p.RotationalDiff.Y, 0)

	// Draw the ship
	rl.DrawRectangle(int32(p.Position.X)-2, int32(p.Position.Y)-2, 4, 4, rl.Red)
	// rl.DrawCircle(int32(s.Position.X), int32(s.Position.Y), 2, rl.Red)

	// Restore the previous transformation matrix
	rl.PopMatrix()
}
