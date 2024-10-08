package ship

import (
	"math"
	"math/rand"

	p "blind/thrust"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const ACC = 0.001
const TURN_RATE = 1

type Ship struct {
	Position     rl.Vector2
	Mass         float32
	Velocity     rl.Vector2
	Acceleration rl.Vector2
	Direction    float32
	Paricles     []*p.ThrustParticle
}

func Initalize(position rl.Vector2) *Ship {

	return &Ship{
		Position:     position,
		Mass:         1.0,
		Velocity:     rl.Vector2{X: 0, Y: 0},
		Acceleration: rl.Vector2{X: 0, Y: 0},
		Direction:    0,
		Paricles:     []*p.ThrustParticle{},
	}

}

func (s *Ship) Update() {
	s.handleInput()
	s.applyForce()
	s.draw()

	// Iterate over the particles slice
	for i := 0; i < len(s.Paricles); {
		particle := s.Paricles[i]
		particle.Update()

		// Condition to remove the particle (example: if particle is out of bounds)
		if particle.StoppedX && particle.StoppedY {
			// Remove the particle by appending slices before and after the current index
			s.Paricles = append(s.Paricles[:i], s.Paricles[i+1:]...)
		} else {
			// Only increment the index if no removal happened
			i++
		}
	}
}

func (s *Ship) handleInput() {
	verSpin, horAcc, verAcc := 0.0, 0.0, 0.0

	if rl.IsKeyDown(rl.KeyLeftShift) {
		if rl.IsKeyDown(rl.KeyW) {
			verAcc += -ACC / 10
			s.Paricles = append(
				s.Paricles,
				p.Initalize(
					rl.Vector2{X: 6, Y: 1},
					rl.Vector2{X: s.Position.X + 6, Y: s.Position.Y + 1},
					s.Direction+rand.Float32()*12-6,
					rand.Float32()*2))
		}

		if rl.IsKeyDown(rl.KeyS) {
			verAcc += ACC / 10

			s.Paricles = append(
				s.Paricles,
				p.Initalize(
					rl.Vector2{X: 11, Y: 0},
					rl.Vector2{X: s.Position.X + 11, Y: s.Position.Y},
					s.Direction+rand.Float32()*12-2-180,
					(rand.Float32()*2)))
		}

		if rl.IsKeyDown(rl.KeyA) {
			horAcc += -ACC / 10

			s.Paricles = append(
				s.Paricles,
				p.Initalize(
					rl.Vector2{X: 5, Y: -8},
					rl.Vector2{X: s.Position.X + 5, Y: s.Position.Y - 8},
					s.Direction+rand.Float32()*2-90,
					(rand.Float32())))

			s.Paricles = append(
				s.Paricles,
				p.Initalize(
					rl.Vector2{X: 5, Y: 8},
					rl.Vector2{X: s.Position.X + 5, Y: s.Position.Y + 8},
					s.Direction+rand.Float32()*2-90,
					(rand.Float32())))
		}

		if rl.IsKeyDown(rl.KeyD) {
			horAcc += ACC / 10

			s.Paricles = append(
				s.Paricles,
				p.Initalize(
					rl.Vector2{X: 5, Y: 8},
					rl.Vector2{X: s.Position.X + 5, Y: s.Position.Y + 8},
					s.Direction+rand.Float32()*2+90,
					(rand.Float32())))

			s.Paricles = append(
				s.Paricles,
				p.Initalize(
					rl.Vector2{X: 5, Y: -8},
					rl.Vector2{X: s.Position.X + 5, Y: s.Position.Y - 8},
					s.Direction+rand.Float32()*2+90,
					(rand.Float32())))

		}
	} else {
		if rl.IsKeyDown(rl.KeyW) {
			verAcc += -ACC
			s.Paricles = append(
				s.Paricles,
				p.Initalize(
					rl.Vector2{X: 6, Y: 1},
					rl.Vector2{X: s.Position.X + 6, Y: s.Position.Y + 1},
					s.Direction+rand.Float32()*12-6,
					(rand.Float32()+2)*2))

			s.Paricles = append(
				s.Paricles,
				p.Initalize(
					rl.Vector2{X: 6, Y: 1},
					rl.Vector2{X: s.Position.X + 6, Y: s.Position.Y + 1},
					s.Direction+rand.Float32()*12-6,
					(rand.Float32()+2)*2))
		}

		if rl.IsKeyDown(rl.KeyS) {
			verAcc += ACC / 10

			s.Paricles = append(
				s.Paricles,
				p.Initalize(
					rl.Vector2{X: 11, Y: 0},
					rl.Vector2{X: s.Position.X + 11, Y: s.Position.Y},
					s.Direction+rand.Float32()*12-2-180,
					(rand.Float32()*2)))
		}

		if rl.IsKeyDown(rl.KeyA) {
			verSpin += -TURN_RATE

			s.Paricles = append(
				s.Paricles,
				p.Initalize(
					rl.Vector2{X: 5, Y: -8},
					rl.Vector2{X: s.Position.X + 5, Y: s.Position.Y - 8},
					s.Direction+rand.Float32()*2-90,
					(rand.Float32())))

			s.Paricles = append(
				s.Paricles,
				p.Initalize(
					rl.Vector2{X: 5, Y: -8},
					rl.Vector2{X: s.Position.X + 5, Y: s.Position.Y - 8},
					s.Direction+rand.Float32()*2+90,
					(rand.Float32())))
		}

		if rl.IsKeyDown(rl.KeyD) {
			verSpin += TURN_RATE

			s.Paricles = append(
				s.Paricles,
				p.Initalize(
					rl.Vector2{X: 5, Y: 8},
					rl.Vector2{X: s.Position.X + 5, Y: s.Position.Y + 8},
					s.Direction+rand.Float32()*2+90,
					(rand.Float32())))

			s.Paricles = append(
				s.Paricles,
				p.Initalize(
					rl.Vector2{X: 5, Y: 8},
					rl.Vector2{X: s.Position.X + 5, Y: s.Position.Y + 8},
					s.Direction+rand.Float32()*2-90,
					(rand.Float32())))
		}
	}
	// Update the ship's direction
	s.Direction += float32(verSpin)

	// Calculate the acceleration components
	rad := s.Direction * (math.Pi / 180.0) // Convert direction to radians
	s.Acceleration = rl.Vector2{
		X: float32(math.Cos(float64(rad))) * float32(verAcc),
		Y: float32(math.Sin(float64(rad))) * float32(verAcc),
	}

	rad = (s.Direction + 270) * (math.Pi / 180.0) // Convert direction to radians
	s.Acceleration.X += float32(math.Cos(float64(rad))) * float32(horAcc)
	s.Acceleration.Y += float32(math.Sin(float64(rad))) * float32(horAcc)

}

func (s *Ship) applyForce() {
	s.Velocity = rl.Vector2{X: s.Velocity.X + s.Acceleration.X, Y: s.Velocity.Y + s.Acceleration.Y}
	s.Position = rl.Vector2{X: s.Position.X + s.Velocity.X, Y: s.Position.Y + s.Velocity.Y}
}

func (s *Ship) draw() {

	// Save the current transformation matrix
	rl.PushMatrix()

	// Translate to the center of the ship
	rl.Translatef(s.Position.X, s.Position.Y, 0)

	// Apply the rotation
	rl.Rotatef(s.Direction-90, 0, 0, 1)

	// Translate back to the original position
	rl.Translatef(-s.Position.X, -s.Position.Y, 0)

	// Draw the ship
	rl.DrawRectangle(int32(s.Position.X)-5, int32(s.Position.Y)-10, 10, 20, rl.White)
	// rl.DrawCircle(int32(s.Position.X), int32(s.Position.Y), 2, rl.Blue)

	// Restore the previous transformation matrix
	rl.PopMatrix()
}
