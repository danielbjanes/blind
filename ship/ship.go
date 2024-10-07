package ship

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
	// "fmt"
)

const ACC = 0.01

type Ship struct {
	Position     rl.Vector2
	Mass         float32
	Velocity     rl.Vector2
	Acceleration rl.Vector2
	Direction    float32
}

func Initalize(position rl.Vector2) *Ship {

	return &Ship{
		Position:     position,
		Mass:         1.0,
		Velocity:     rl.Vector2{X: 0, Y: 0},
		Acceleration: rl.Vector2{X: 0, Y: 0},
		Direction:    0,
	}

}

func (s *Ship) Update() {
	s.handleInput()
	s.ApplyForce()
	s.Draw()
}

func (s *Ship) handleInput() {

	verAcc, horAcc := 0.0, 0.0

	if rl.IsKeyDown(rl.KeyW) {
		horAcc = -ACC
	}

	if rl.IsKeyDown(rl.KeyS) {
		horAcc = ACC
	}

	if rl.IsKeyDown(rl.KeyA) {
		verAcc = -ACC
	}

	if rl.IsKeyDown(rl.KeyD) {
		verAcc = ACC
	}

	s.Acceleration = rl.Vector2{X: float32(verAcc), Y: float32(horAcc)}

	fmt.Println(s.Acceleration)

}

func (s *Ship) ApplyForce() {
	s.Velocity = rl.Vector2{X: s.Velocity.X + s.Acceleration.X, Y: s.Velocity.Y + s.Acceleration.Y}
	s.Position = rl.Vector2{X: s.Position.X + s.Velocity.X, Y: s.Position.Y + s.Velocity.Y}
}

func (s *Ship) Draw() {
	rl.DrawRectangle(int32(s.Position.X), int32(s.Position.Y), 50, 50, rl.Red)
}
