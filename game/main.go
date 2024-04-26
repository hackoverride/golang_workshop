package main

import (
	"fmt"
	"game/player"
	"game/world"
	"log"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	tileSize     int   = 50
	screenHeight int32 = 800
	screenWidth  int32 = 1200

	worldWidth  int = 20
	worldHeight int = 12
	worldMargin int = 2

	playerAnimationSpeed     float64 = 0.1
	playerIdleAnimationSpeed float64 = 0.2
)

var (
	grassSprite      rl.Texture2D // grass sprite
	grassSrc         rl.Rectangle // grass source rectangle
	playerSprite     rl.Texture2D // player sprite
	playerSpriteLeft rl.Texture2D
	playerSrc        rl.Rectangle // player source rectangle

	// Walking animation frames
	playerSourcePositions = []rl.Rectangle{
		rl.NewRectangle(22, 78, 24, 50), // Starting position
		rl.NewRectangle(86, 78, 24, 50),
		rl.NewRectangle(150, 78, 24, 50),
		rl.NewRectangle(215, 78, 24, 50),
		rl.NewRectangle(278, 78, 24, 50),
		rl.NewRectangle(342, 78, 24, 50),
		rl.NewRectangle(407, 78, 24, 50),
		rl.NewRectangle(471, 78, 24, 50),
	}
	playerIdlePositions = []rl.Rectangle{
		rl.NewRectangle(22, 14, 24, 50), // Starting position
		rl.NewRectangle(86, 14, 24, 50),
		rl.NewRectangle(150, 14, 24, 50),
		rl.NewRectangle(215, 14, 24, 50),
		rl.NewRectangle(278, 14, 24, 50),
		rl.NewRectangle(342, 14, 24, 50),
		rl.NewRectangle(407, 14, 24, 50),
		rl.NewRectangle(471, 14, 24, 50),
	}
	lastAnimationTime float64 = 0

	TheWorld = world.NewWorld(worldWidth, worldHeight)
	pl       player.PlayerType
)

func init() {
	rl.InitWindow(screenWidth, screenHeight, "Capgemini - Go Workshop")
	rl.SetTargetFPS(60)

	pl = player.NewPlayer(float32(screenWidth/2), float32(screenHeight/2), 2)
	pl2 := player.NewPlayer(float32(screenWidth/2), float32(screenHeight/2), 1)
	fmt.Println(pl.GetSpeed(), pl2.GetSpeed())

	grassSprite = loadTexture("assets/tilesets/Grass.png")
	playerSprite = loadTexture("assets/characters/lund_right.png")
	playerSpriteLeft = loadTexture("assets/characters/lund_left.png")
}

func loadTexture(filePath string) rl.Texture2D {
	texture := rl.LoadTexture(filePath)
	if texture.ID == 0 {
		log.Fatalf("Failed to load texture: %s", filePath)
	}
	return texture
}

func renderWorld() {
	for x, column := range TheWorld.Tiles {
		for y, tile := range column {
			if tile.Grass {
				renderTile(x, y, &tile)
			}
		}
	}
}

func renderTile(x, y int, tile *world.Tile) {
	destX, destY := (x+worldMargin)*tileSize, (y+worldMargin)*tileSize
	destination := rl.NewRectangle(float32(destX), float32(destY), float32(tileSize), float32(tileSize))
	grassSrc = selectTileSource(x, y)

	grassTint := rl.NewColor(255, 255, 255, 255)
	if (x+y)%2 == 0 {
		grassTint = rl.NewColor(240, 240, 255, 255)
	}

	rl.DrawTexturePro(grassSprite, grassSrc, destination, rl.Vector2{}, 0, grassTint)
}

func selectTileSource(x, y int) rl.Rectangle {
	switch {
	case x == 0 && y == 0:
		return rl.NewRectangle(0, 0, 16, 16)
	case x == worldWidth-1 && y == worldHeight-1:
		return rl.NewRectangle(32, 32, 16, 16)
	case x == 0 && y == worldHeight-1:
		return rl.NewRectangle(0, 32, 16, 16)
	case x == worldWidth-1 && y == 0:
		return rl.NewRectangle(32, 0, 16, 16)
	case x == 0:
		return rl.NewRectangle(0, 16, 16, 16)
	case x == worldWidth-1:
		return rl.NewRectangle(32, 16, 16, 16)
	case y == 0:
		return rl.NewRectangle(16, 0, 16, 16)
	case y == worldHeight-1:
		return rl.NewRectangle(16, 32, 16, 16)
	default:
		return rl.NewRectangle(16, 16, 16, 16)
	}
}

func renderPlayer() {
	currentTime := rl.GetTime()
	playerPosX, playerPosY := pl.GetPosition()
	playerRenderCycle := pl.GetRenderPos()
	triggerSpeed := playerAnimationSpeed
	if !pl.IsMoving {
		triggerSpeed = playerIdleAnimationSpeed
	}

	// Update animation based on time, not every frame
	if currentTime-lastAnimationTime > triggerSpeed {
		lastAnimationTime = currentTime
		playerRenderCycle = (playerRenderCycle + 1) % len(playerSourcePositions)
		pl.SetRenderPos(playerRenderCycle) // Assuming there's a method to set the render cycle
	}

	// Player sprite source is based upon which render cycle we are in
	var isMoving = pl.IsMoving
	if !isMoving {
		// we are missing one position for the idle animation, when facing left we are missing the first position
		// when facing right we are missing the last position

		if pl.IsPlayerFaceRight() && playerRenderCycle == len(playerIdlePositions)-1 {
			playerRenderCycle = 0
		}
		if !pl.IsPlayerFaceRight() && playerRenderCycle == 0 {
			playerRenderCycle = 1
		}

		playerSrc = playerIdlePositions[playerRenderCycle]
	} else {
		playerSrc = playerSourcePositions[playerRenderCycle]
	}

	// Render the player
	spriteTarget := playerSprite
	if !pl.IsPlayerFaceRight() {
		spriteTarget = playerSpriteLeft
	}
	rl.DrawTexturePro(spriteTarget, playerSrc, rl.NewRectangle(playerPosX, playerPosY, 50, 72), rl.NewVector2(0, 0), 0, rl.White)
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
		pl.SetPlayerMoving(true)
		pl.SetPlayerFaceRight(true)
	} else if rl.IsKeyDown(rl.KeyLeft) || rl.IsKeyDown(rl.KeyA) {
		pl.Move(-speed, 0)
		pl.SetPlayerMoving(true)
		pl.SetPlayerFaceRight(false)
	} else if rl.IsKeyDown(rl.KeyUp) || rl.IsKeyDown(rl.KeyW) {
		pl.Move(0, -speed)
		pl.SetPlayerMoving(true)
	} else if rl.IsKeyDown(rl.KeyDown) || rl.IsKeyDown(rl.KeyS) {
		pl.Move(0, speed)
		pl.SetPlayerMoving(true)
	}

	if rl.IsKeyReleased(rl.KeyRight) || rl.IsKeyReleased(rl.KeyD) ||
		rl.IsKeyReleased(rl.KeyLeft) || rl.IsKeyReleased(rl.KeyA) ||
		rl.IsKeyReleased(rl.KeyUp) || rl.IsKeyReleased(rl.KeyW) ||
		rl.IsKeyReleased(rl.KeyDown) || rl.IsKeyReleased(rl.KeyS) {
		pl.SetPlayerMoving(false)
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
