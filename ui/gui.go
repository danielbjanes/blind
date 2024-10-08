package ui

import (
	s "blind/ship"
	t "blind/types"
	"math"
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func Draw(state *t.State) {

	drawShipStats(state.Ship)

	// rl.DrawText("fps: "+strconv.Itoa(int(rl.GetFPS())), 30, 20, 20, rl.Red)
	// rl.DrawLine(20, 55, 200, 55, rl.White)
}

func drawShipStats(ship *s.Ship) {

	centerX, centerY := int32(90), int32(90)

	rl.DrawCircle(centerX, centerY, 52, rl.White)
	rl.DrawCircle(centerX, centerY, 51, rl.Black)

	rl.DrawCircle(centerX, centerY, 40, rl.White)
	rl.DrawCircle(centerX, centerY, 39, rl.Black)

	rl.DrawCircle(centerX, centerY, 12, rl.White)
	rl.DrawCircle(centerX, centerY, 11, rl.Black)

	rl.DrawRectangle(centerX-1, centerY-52, 2, 52*2, rl.White)
	rl.DrawRectangle(centerX-52, centerY-1, 52*2, 2, rl.White)

	// rl.DrawLine(centerX-20, centerY-20, centerX+20, centerY+20, rl.White)
	// rl.DrawLine(centerX-20, centerY+20, centerX+20, centerY-20, rl.White)

	drawVelocity(ship)

	// Coordinates
	rl.DrawText("{ "+strconv.Itoa(int(ship.Position.X))+
		", "+strconv.Itoa(int(ship.Position.Y))+"}",
		centerX-52-10, centerY+52+10, 8, rl.White)

	if rl.IsKeyDown(rl.KeyLeftShift) {
		rl.DrawText("[position]", centerX+52-28, centerY+52+10, 8, rl.White)
	} else {
		rl.DrawText("[thrust]", centerX+52-28, centerY+52+10, 8, rl.White)
	}

	drawDirection(ship)
}

func drawVelocity(ship *s.Ship) {

	centerX, centerY := int32(90), int32(90)

	angleRad := math.Atan2(float64(ship.Velocity.Y), float64(ship.Velocity.X))
	angleDeg := angleRad*(180.0/math.Pi) - 90

	rl.PushMatrix()
	rl.Translatef(float32(centerX), float32(centerY), 0)
	rl.Rotatef(float32(angleDeg), 0, 0, 1)
	rl.Translatef(-float32(centerX), -float32(centerY), 0)

	velocityMagnitude := math.Sqrt(float64(ship.Velocity.X*ship.Velocity.X+ship.Velocity.Y*ship.Velocity.Y)) * 100

	rl.DrawRectangle(centerX-1, centerY-1, 3, int32(velocityMagnitude), rl.Red)

	rl.PopMatrix()
}

func drawDirection(ship *s.Ship) {

	centerX, centerY := int32(90), int32(90)

	offsetX, offsetY := centerX+52, centerY-52

	rl.PushMatrix()
	rl.Translatef(float32(offsetX), float32(offsetY), 0)
	rl.Rotatef(float32(ship.Direction-90), 0, 0, 1)
	rl.Translatef(-float32(offsetX), -float32(offsetY), 0)

	rl.DrawRectangle(offsetX-4, offsetY-8, 8, 16, rl.White)

	rl.PopMatrix()
}
