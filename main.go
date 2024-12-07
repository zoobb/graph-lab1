package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"lab1/entity"
)

func main() {
	var screenWidth int32 = 1200
	var screenHeight int32 = 800
	screenCenter := rl.Vector2{
		X: float32(screenWidth / 2),
		Y: float32(screenHeight / 2),
	}

	e := entity.New(
		50,
		200,
		100,
		10,
		0.1,
		rl.White,
		screenCenter,
		entity.Info{
			FontSize:  20,
			FontColor: rl.Blue,
			Gap:       2,
			Pos: rl.Vector2{
				X: 10,
				Y: 10,
			},
		},
	)

	rl.InitWindow(screenWidth, screenHeight, "boop")
	defer rl.CloseWindow()

	var lastFrameTime = float32(rl.GetTime())

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		var currentFrameTime = float32(rl.GetTime())
		deltaTime := currentFrameTime - lastFrameTime
		lastFrameTime = currentFrameTime
		e.DeltaTime = deltaTime

		e.Control()

		rl.BeginDrawing()
		rl.ClearBackground(rl.Blank)

		e.Draw()
		e.DrawInfo()

		rl.EndDrawing()
	}
}
