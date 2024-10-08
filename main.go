package main

import (
	"strconv"

	s "blind/ship"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type State struct {
	WindowHeight int32
	WindowWidth  int32
}

func main() {
	initalization()

	ship := s.Initalize(rl.Vector2{X: 400, Y: 400})

	for !rl.WindowShouldClose() {
		// update(state)
		draw(ship)
	}
}

func draw(ship *s.Ship) {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)

	ship.Update()

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
