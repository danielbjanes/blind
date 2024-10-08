package main

import (
	s "blind/ship"
	t "blind/types"
	ui "blind/ui"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	state := initalization()

	for !rl.WindowShouldClose() {
		// update(state)
		draw(state)
	}
}

func draw(state *t.State) {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)

	ui.Draw(state)
	state.Ship.Update()

	rl.EndDrawing()
}

func initalization() *t.State {

	state := &t.State{
		WindowWidth:  800,
		WindowHeight: 800,
		Ship:         s.Initalize(rl.Vector2{X: 400, Y: 400}),
	}

	rl.InitWindow(state.WindowWidth, state.WindowHeight, "")
	rl.SetTargetFPS(120)

	return state
}
