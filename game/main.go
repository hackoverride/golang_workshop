package main

import (
	"fmt"
	"game/item"
	"game/player"
	"game/world"
	"log"
	"sort"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	tileSize     int   = 50
	screenHeight int32 = 800
	screenWidth  int32 = 1200

	worldWidth  int = 21
	worldHeight int = 15

	playerAnimationSpeed     float64 = 0.1
	playerIdleAnimationSpeed float64 = 0.2
	playerSpellSpeed         float64 = 0.02
	flameAnimationSpeed      float64 = 0.04
)

type RenderableItem struct {
	Texture     rl.Texture2D
	Source      rl.Rectangle
	Width       float32
	Height      float32
	Destination rl.Rectangle
	Color       rl.Color
}

var (
	debugMode bool = true

	camera rl.Camera2D

	grassSprite       rl.Texture2D
	playerSprite      rl.Texture2D
	playerSpriteLeft  rl.Texture2D
	itemSprite        rl.Texture2D
	FireSpriteTexture rl.Texture2D
	TreeSpriteTexture rl.Texture2D
	grassSrc          rl.Rectangle
	playerSrc         rl.Rectangle

	// Walking animation frames
	playerSourcePositions = []rl.Rectangle{
		rl.NewRectangle(22, 78, 20, 36), // Starting position
		rl.NewRectangle(86, 78, 20, 36),
		rl.NewRectangle(150, 78, 20, 36),
		rl.NewRectangle(215, 78, 20, 36),
		rl.NewRectangle(278, 78, 20, 36),
		rl.NewRectangle(342, 78, 20, 36),
		rl.NewRectangle(407, 78, 20, 36),
		rl.NewRectangle(471, 78, 20, 36),
	}
	playerIdlePositions = []rl.Rectangle{
		rl.NewRectangle(22, 14, 20, 36), // Starting position
		rl.NewRectangle(86, 14, 20, 36),
		rl.NewRectangle(150, 14, 20, 36),
		rl.NewRectangle(215, 14, 20, 36),
		rl.NewRectangle(278, 14, 20, 36),
		rl.NewRectangle(342, 14, 20, 36),
		rl.NewRectangle(407, 14, 20, 36),
		rl.NewRectangle(471, 14, 20, 36),
	}
	lastAnimationTime  float64 = 0
	lastSpellAnimation float64 = 0
	lastFlameAnimation float64 = 0

	TheWorld = world.NewWorld(worldWidth, worldHeight)
	pl       player.PlayerType
	items    []item.ItemType  = make([]item.ItemType, 0)
	flames   []item.FireLight = make([]item.FireLight, 0)

	lockedChestSource = rl.NewRectangle(92, 31, 40, 35)
	openChestSource   = rl.NewRectangle(93, 74, 40, 55)
)

func init() {
	rl.InitWindow(screenWidth, screenHeight, "Game")
	rl.SetTargetFPS(60)

	pl = player.NewPlayer(float32(screenWidth/2), float32(screenHeight/2), 2)
	camera = rl.NewCamera2D(rl.NewVector2(float32(screenWidth/2), float32(screenHeight/2)),
		rl.NewVector2(float32(pl.PosX-(pl.Width/2)), float32(pl.PosY-(pl.Height/2))), 0, 1)

	grassSprite = loadTexture("assets/tilesets/Grass.png")
	playerSprite = loadTexture("assets/characters/lund_right.png")
	playerSpriteLeft = loadTexture("assets/characters/lund_left.png")
	itemSprite = loadTexture("assets/objects/tx_props.png")
	TreeSpriteTexture = loadTexture("assets/objects/grass_biom.png")
	FireSpriteTexture = loadTexture("assets/effects/fire.png")
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
	destX, destY := (x)*tileSize, (y)*tileSize
	destination := rl.NewRectangle(float32(destX), float32(destY-tileSize), float32(tileSize), float32(tileSize))
	grassSrc = selectTileSource(x, y)
	playerPosX, playerPosY := pl.GetPosition()
	playerWidth := pl.GetPlayerWidth()
	playerHeight := pl.GetPlayerHeight()

	playerPosY -= playerHeight

	includesPlayerX := playerPosY+playerHeight >= float32(destY-tileSize) && playerPosY <= float32(destY)
	includesPlayerY := playerPosX+playerWidth >= float32(destX) && playerPosX <= float32(destX+tileSize)

	if tile.Grass {
		grassTint := rl.NewColor(255, 255, 255, 255)
		if (x+y)%2 == 0 {
			grassTint = rl.NewColor(240, 240, 255, 255)
		}
		if includesPlayerX && debugMode {
			grassTint = rl.NewColor(200, 200, 200, 255)
		}
		if includesPlayerY && debugMode {
			grassTint = rl.NewColor(200, 200, 200, 255)
		}
		if debugMode && includesPlayerX && includesPlayerY {
			grassTint = rl.NewColor(150, 150, 150, 255)
		}
		rl.DrawTexturePro(grassSprite, grassSrc, destination, rl.Vector2{}, 0, grassTint)
	}

	if debugMode {
		rl.DrawRectangleLines(int32(destX), int32(destY-tileSize), int32(tileSize), int32(tileSize), rl.White)
	}
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

func renderLayers() {

	renderQueue := make([]RenderableItem, 0)

	renderQueue = renderItems(renderQueue)
	renderQueue = renderPlayer(renderQueue)
	renderQueue = renderSpells(renderQueue)

	// Sort the render queue by Y position
	sort.Slice(renderQueue, func(i, j int) bool {
		return renderQueue[i].Destination.Y < renderQueue[j].Destination.Y
	})

	for _, renderItem := range renderQueue {
		renderItem.Destination.Y -= renderItem.Height
		rl.DrawTexturePro(renderItem.Texture, renderItem.Source, renderItem.Destination, rl.Vector2{}, 0, renderItem.Color)

		if debugMode {
			rl.DrawRectangleLines(int32(renderItem.Destination.X), int32(renderItem.Destination.Y), int32(renderItem.Width), int32(renderItem.Height), rl.Black)
			/* mark item */
			rl.DrawCircle(int32(renderItem.Destination.X), int32(renderItem.Destination.Y), 2, rl.Green)
			rl.DrawCircle(int32(renderItem.Destination.X+renderItem.Width), int32(renderItem.Destination.Y), 2, rl.Green)
			rl.DrawCircle(int32(renderItem.Destination.X), int32(renderItem.Destination.Y+renderItem.Height), 2, rl.Green)
			rl.DrawCircle(int32(renderItem.Destination.X+renderItem.Width), int32(renderItem.Destination.Y+renderItem.Height), 2, rl.Green)
		}
	}
}

func renderPlayer(renderQueue []RenderableItem) []RenderableItem {
	currentTime := rl.GetTime()
	playerPosX, playerPosY := pl.GetPosition()
	playerWidth := pl.GetPlayerWidth()
	playerHeight := pl.GetPlayerHeight()
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
	renderQueue = append(renderQueue, RenderableItem{
		Texture:     spriteTarget,
		Source:      playerSrc,
		Width:       float32(playerWidth),
		Height:      float32(playerHeight),
		Destination: rl.NewRectangle(playerPosX, playerPosY, float32(playerWidth), float32(playerHeight)),
		Color:       rl.White,
	})

	return renderQueue
}

func renderSpells(renderQueue []RenderableItem) []RenderableItem {
	currentTime := rl.GetTime()
	if currentTime-lastSpellAnimation > playerSpellSpeed {
		lastSpellAnimation = currentTime
		// render spell
	}

	return renderQueue
}

func renderItems(renderQueue []RenderableItem) []RenderableItem {

	for _, item := range items {
		positionX, positionY := item.GetPosition()
		source := item.GetSpriteSource()
		width := item.GetWidth()
		height := item.GetHeight()

		renderQueue = append(renderQueue, RenderableItem{
			Texture:     itemSprite,
			Source:      source,
			Width:       float32(width),
			Height:      float32(height),
			Destination: rl.NewRectangle(positionX, positionY, float32(width), float32(height)),
			Color:       rl.White,
		})
	}

	for _, flame := range flames {
		positionX, positionY := flame.GetPosition()
		texture := FireSpriteTexture
		renderPos := flame.GetRenderPos()

		source := item.FireSpriteSource[renderPos]
		width := flame.GetFireWidth()
		height := flame.GetFireHeight()

		renderQueue = append(renderQueue, RenderableItem{
			Texture:     texture,
			Source:      source,
			Height:      height,
			Width:       width,
			Destination: rl.NewRectangle(positionX, positionY, float32(width), float32(height)),
			Color:       rl.White,
		})
	}
	trees := TheWorld.GetTrees()

	for _, tree := range trees {
		positionX, positionY := tree.GetPosition()
		source := tree.GetTreeSprite()
		width := tree.GetTreeWidth()
		positionX -= width / 2
		height := tree.GetTreeHeight()
		color := rl.White
		opacity := 255
		health := tree.GetHealth()
		switch {
		case health == 100:
			opacity = 255
		case health < 100:
			opacity = 200
		case health < 75:
			opacity = 150
		case health < 25:
			opacity = 100
		}

		color = rl.NewColor(255, 255, 255, uint8(opacity))

		renderQueue = append(renderQueue, RenderableItem{
			Texture:     TreeSpriteTexture,
			Source:      source,
			Height:      height,
			Width:       width,
			Destination: rl.NewRectangle(positionX, positionY, width, height),
			Color:       color,
		})
	}

	return renderQueue
}

func render() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.SkyBlue)
	// The order it is rendered is important
	rl.BeginMode2D(camera)
	renderWorld()
	renderLayers()
	rl.EndMode2D()
	rl.EndDrawing()
}

func update() {
	currentTime := rl.GetTime()
	// This is where we do our game logic
	if currentTime-lastFlameAnimation > flameAnimationSpeed {
		for i := range flames {
			flames[i].SetRenderPos((flames[i].GetRenderPos() + 1) % len(item.FireSpriteSource))
		}
		lastFlameAnimation = currentTime
	}

	TheWorld.Update()
	// Update the camera
	camera.Target = rl.NewVector2(float32(pl.PosX), float32(pl.PosY))
	camera.Offset = rl.NewVector2(float32(screenWidth/2), float32(screenHeight/2))

}

func input() {
	// This is where we handle user input
	speed := pl.GetSpeed()
	if rl.IsKeyPressed(rl.KeyP) {
		debugMode = !debugMode
	}
	if rl.IsKeyDown(rl.KeyRight) || rl.IsKeyDown(rl.KeyD) {
		pl.Move(speed, 0)
		pl.SetPlayerMoving(true)
		pl.SetPlayerFaceRight(true)
	}
	if rl.IsKeyDown(rl.KeyLeft) || rl.IsKeyDown(rl.KeyA) {
		pl.Move(-speed, 0)
		pl.SetPlayerMoving(true)
		pl.SetPlayerFaceRight(false)
	}
	if rl.IsKeyDown(rl.KeyUp) || rl.IsKeyDown(rl.KeyW) {
		pl.Move(0, -speed)
		pl.SetPlayerMoving(true)
	}
	if rl.IsKeyDown(rl.KeyDown) || rl.IsKeyDown(rl.KeyS) {
		pl.Move(0, speed)
		pl.SetPlayerMoving(true)
	}

	if rl.IsKeyReleased(rl.KeyRight) || rl.IsKeyReleased(rl.KeyD) ||
		rl.IsKeyReleased(rl.KeyLeft) || rl.IsKeyReleased(rl.KeyA) ||
		rl.IsKeyReleased(rl.KeyUp) || rl.IsKeyReleased(rl.KeyW) ||
		rl.IsKeyReleased(rl.KeyDown) || rl.IsKeyReleased(rl.KeyS) {
		pl.SetPlayerMoving(false)
	}

	/* Use items */
	if rl.IsKeyPressed(rl.KeyE) {
		positionX, positionY := pl.GetPosition()
		playerHeight := pl.GetPlayerHeight()
		playerWidth := pl.GetPlayerWidth()

		alreadyTriggeredOne := false
		for i, it := range items {
			itemX, itemY := it.GetPosition()
			itenWidth := it.GetWidth()
			itemHeight := it.GetHeight()

			// Check if the player posX and posX + playerWidth is greater than the itemX and less than the itemX + itemWidth
			overlappingX := positionX+playerWidth >= itemX && positionX <= itemX+itenWidth
			// Check if the player posY and posY - playerHeight is less than the itemY and greater than the itemY - itemHeight
			overlappingY := positionY >= itemY-itemHeight && positionY-playerHeight <= itemY
			if overlappingX && overlappingY && !alreadyTriggeredOne {
				fmt.Printf("Player is on top of item %d\n", i)
				// if the item is a chest_locked, remove it and add a chest_opened
				if it.Name == "chest_locked" {
					items = append(items[:i], items[i+1:]...)
					items = append(items, item.NewItem("chest_opened", rl.NewRectangle(93, 74, 40, 55), itemX, itemY, 40, 50))
				} else {
					items = append(items[:i], items[i+1:]...)
				}
				alreadyTriggeredOne = true

				// Remove the item from the world
			}
		}
		allTrees := TheWorld.GetTrees()
		for i, tree := range allTrees {
			treeX, treeY := tree.GetPosition()
			treeWidth := tree.GetTreeWidth()
			treeHeight := tree.GetTreeHeight()

			// Check if the player posX and posX + playerWidth is greater than the itemX and less than the itemX + itemWidth
			overlappingX := positionX+playerWidth >= treeX && positionX <= treeX+treeWidth
			// Check if the player posY and posY - playerHeight is less than the itemY and greater than the itemY - itemHeight
			overlappingY := positionY >= treeY-treeHeight && positionY-playerHeight <= treeY
			if overlappingX && overlappingY && !alreadyTriggeredOne {
				fmt.Printf("Player is on top of tree %d\n", i)
				TheWorld.HitTree(i, 10)
			}
		}
	}

	/* Create Fire */
	if rl.IsKeyPressed(rl.KeyF) {
		positionX, positionY := pl.GetPosition()
		flames = append(flames, item.NewFireLight(positionX, positionY))
	}

	/* Create Tree */
	if rl.IsKeyPressed(rl.KeyT) {
		positionX, positionY := pl.GetPosition()
		TheWorld.AddTree(positionX, positionY)
	}

	/* Use Spell */
	if rl.IsKeyPressed(rl.KeyQ) {
		fmt.Println("Player is using spell")
		pl.SetPlayerUsingSpell(true)
	}

	if rl.IsKeyReleased(rl.KeyQ) {
		pl.SetPlayerUsingSpell(false)
	}

	/* Place items */

	if rl.IsKeyPressed(rl.KeySpace) {
		// Add an item to the world
		positionX, positionY := pl.GetPosition()
		items = append(items, item.NewItem("chest_locked", lockedChestSource, positionX, positionY, 40, 30))
	}
}

func main() {

	rl.InitAudioDevice()

	music := rl.LoadMusicStream("intro.mp3")

	rl.PlayMusicStream(music)
	rl.SetMusicVolume(music, 0.05)

	for !rl.WindowShouldClose() {
		rl.UpdateMusicStream(music) // Update music buffer with new stream data
		input()
		update()
		render()
	}

	rl.CloseWindow()
}
