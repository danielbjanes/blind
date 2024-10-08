package effects

import (
	"fmt"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const ENTROPY = 0.1

type ThrustParticle struct {
	Position  rl.Vector2
	Velocity  rl.Vector2
	Ttl       int32
	Direction float32
	Magnitude float32
}

func Initalize(position rl.Vector2, ttl int32, direction float32, magnitude float32) *ThrustParticle {

	rad := direction * (math.Pi / 180.0) // Convert direction to radians
	initVel := rl.Vector2{
		X: float32(math.Cos(float64(rad))) * float32(magnitude),
		Y: float32(math.Sin(float64(rad))) * float32(magnitude),
	}

	return &ThrustParticle{
		Position:  position,
		Ttl:       ttl,
		Direction: direction,
		Velocity:  initVel,
		Magnitude: magnitude,
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

	fmt.Println("--")
	fmt.Printf("vel: %f, %f\n", p.Velocity.X, p.Velocity.Y)
	// fmt.Printf("acc: %f, %f\n", acceleration.X, acceleration.Y)

	if math.Abs(float64(p.Velocity.X))+math.Abs(float64(acceleration.X)) > 0.1 {
		p.Velocity.X += acceleration.X
		p.Velocity.Y += acceleration.Y
	}
	// p.Velocity.X += acceleration.X
	// p.Velocity.Y += acceleration.Y

	p.Position.X += p.Velocity.X
	p.Position.Y += p.Velocity.Y
}

func (p *ThrustParticle) draw() {
	// Save the current transformation matrix
	rl.PushMatrix()

	// Translate to the center of the ship
	rl.Translatef(p.Position.X, p.Position.Y, 0)

	// Apply the rotation
	rl.Rotatef(p.Direction, 0, 0, 1)

	// Translate back to the original position
	rl.Translatef(-p.Position.X, -p.Position.Y, 0)

	// Draw the ship
	rl.DrawRectangle(int32(p.Position.X)-2, int32(p.Position.Y)-2, 4, 4, rl.White)
	// rl.DrawCircle(int32(s.Position.X), int32(s.Position.Y), 2, rl.Red)

	// Restore the previous transformation matrix
	rl.PopMatrix()
}
