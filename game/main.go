package main

import (
	"game/player"
	"game/world"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	tileSize     int   = 50
	screenHeight int32 = 800
	screenWidth  int32 = 1200

	worldWidth  int = 20
	worldHeight int = 12
	worldMargin int = 2
)

var (
	grassSprite  rl.Texture2D // grass sprite
	grassSrc     rl.Rectangle // grass source rectangle
	playerSprite rl.Texture2D // player sprite
	playerSrc    rl.Rectangle // player source rectangle

	TheWorld = world.NewWorld(worldWidth, worldHeight)
	pl       *player.PlayerType
)

func init() {
	rl.InitWindow(screenWidth, screenHeight, "Capgemini - Go Workshop")
	rl.SetTargetFPS(60)

	// Find the middle of the screen
	playerStartPosX := float32(screenWidth / 2)
	playerStartPosY := float32(screenHeight / 2)
	pl = player.NewPlayer(playerStartPosX, playerStartPosY)
	grassSprite = rl.LoadTexture("assets/tilesets/Grass.png")
	playerSprite = rl.LoadTexture("assets/characters/lund_idle_walk.png")
}
func renderWorld() {
	// Here we only render the world

	for x := range TheWorld.Tiles {
		for y := range TheWorld.Tiles[x] {
			// We only want to render the tiles that are visible on the screen

			if TheWorld.Tiles[x][y].Grass {
				// destination rectangle
				currentPlacementX := (x + worldMargin) * tileSize
				currentPlacementY := (y + worldMargin) * tileSize
				destination := rl.NewRectangle(float32(currentPlacementX), float32(currentPlacementY), float32(tileSize), float32(tileSize))
				// source rectangle
				grassSrc = rl.NewRectangle(16, 16, 16, 16) // This is the grass tile
				if x == 0 {
					// This is the left edge of the world
					grassSrc = rl.NewRectangle(0, 16, 16, 16)
				} else if x == worldWidth-1 {
					grassSrc = rl.NewRectangle(32, 16, 16, 16)
				} else if y == 0 {
					grassSrc = rl.NewRectangle(16, 0, 16, 16)
				} else if y == worldHeight-1 {
					grassSrc = rl.NewRectangle(16, 32, 16, 16)
				}

				switch {
				case x == 0 && y == 0:
					grassSrc = rl.NewRectangle(0, 0, 16, 16)
				case x == 0 && y == worldHeight-1:
					grassSrc = rl.NewRectangle(0, 32, 16, 16)
				case x == worldWidth-1 && y == 0:
					grassSrc = rl.NewRectangle(32, 0, 16, 16)
				case x == worldWidth-1 && y == worldHeight-1:
					grassSrc = rl.NewRectangle(32, 32, 16, 16)
				}

				grassTint := rl.NewColor(uint8(255), uint8(255), uint8(255), 255)
				if (x%2) == 0 && (y%2) == 0 {
					grassTint = rl.NewColor(uint8(240), uint8(240), uint8(255), 255)
				}
				rl.DrawTexturePro(grassSprite, grassSrc, destination, rl.Vector2{X: 0, Y: 0}, 0, grassTint)
			}
		}
	}
}

func renderPlayer() {
	// Here we only render the player
	playerPosX, playerPosY := pl.GetPosition()
	playerSrc = rl.NewRectangle(20, 78, 52, 52)
	rl.DrawTexturePro(playerSprite, playerSrc, rl.NewRectangle(playerPosX, playerPosY, 62, 62), rl.NewVector2(0, 0), 0, rl.White)
	// rl.DrawCircle(int32(playerPosX), int32(playerPosY), 10, rl.Red)
}

func render() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.SkyBlue)
	// The order it is rendered is important
	renderWorld()
	renderPlayer()
	rl.EndDrawing()
}

func update() {
	// This is where we do our game logic
}

func input() {
	// This is where we handle user input
	speed := pl.GetSpeed()
	if rl.IsKeyDown(rl.KeyRight) || rl.IsKeyDown(rl.KeyD) {
		pl.Move(speed, 0)
	} else if rl.IsKeyDown(rl.KeyLeft) || rl.IsKeyDown(rl.KeyA) {
		pl.Move(-speed, 0)
	} else if rl.IsKeyDown(rl.KeyUp) || rl.IsKeyDown(rl.KeyW) {
		pl.Move(0, -speed)
	} else if rl.IsKeyDown(rl.KeyDown) || rl.IsKeyDown(rl.KeyS) {
		pl.Move(0, speed)
	}
}

func main() {
	for !rl.WindowShouldClose() {
		input()
		update()
		render()
	}

	rl.CloseWindow()
}
