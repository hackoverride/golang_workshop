package main

import (
	"fmt"
	"github.com/gen2brain/raylib-go/raylib"
)

func main() {
	fmt.Println("Welcome to the Go workshop!")

	rl.InitWindow(800, 450, "Capgemini Go Workshop");
	rl.SetTargetFPS(60);

	for !rl.WindowShouldClose() {
		rl.BeginDrawing();
		rl.ClearBackground(rl.RayWhite);
		rl.DrawText("Welcome to the Go workshop!", 10, 10, 20, rl.DarkGray);
		rl.EndDrawing();
	}

	rl.CloseWindow();
}
