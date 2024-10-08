package main

import (
	"strconv"

	s "blind/ship"

	e "blind/effects"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type State struct {
	WindowHeight int32
	WindowWidth  int32
}

func main() {
	initalization()

	ship := s.Initalize(rl.Vector2{X: 400, Y: 400})

	p := e.Initalize(rl.Vector2{X: 400, Y: 400}, 1, 180, 4)

	for !rl.WindowShouldClose() {
		// update(state)
		draw(ship, p)
	}
}

func draw(ship *s.Ship, p *e.ThrustParticle) {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)

	ship.Update()
	p.Update()

	rl.DrawText("fps: "+strconv.Itoa(int(rl.GetFPS())), 20, 20, 20, rl.Red)
	rl.EndDrawing()
}

func initalization() *State {

	state := &State{
		WindowWidth:  800,
		WindowHeight: 800,
	}

	rl.InitWindow(state.WindowWidth, state.WindowHeight, "")
	rl.SetTargetFPS(120)

	return state
}
