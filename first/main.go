package main

import (
	em "first/enemy"
	p "first/player"
	"first/world"
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = 1200
	screenHeight = 800
)

var (
	pause bool = false

	/* World area that is displayed */
	worldMapX     int32 = 0
	worldMapY     int32 = 0
	worldTileSize int32 = 100

	numberOfTilesWide = screenWidth / worldTileSize
	numberOfTilesHigh = screenHeight / worldTileSize

	/* Sprites */
	playerSprite rl.Texture2D // player sprite
	playerSrc    rl.Rectangle // player source rectangle
	playerDest   rl.Rectangle // player destination rectangle

	grassSprite rl.Texture2D // grass sprite
	grassSrc    rl.Rectangle // grass source rectangle
	grassDest   rl.Rectangle // grass destination rectangle

	biomeSprite rl.Texture2D // tree sprite
	treeSrc     rl.Rectangle // tree source rectangle
	treeDest    rl.Rectangle // tree destination rectangle

	stoneSprite rl.Texture2D // stone sprite
	stoneSrc    rl.Rectangle // stone source rectangle
	stoneDest   rl.Rectangle // stone destination rectangle

	demonSprite, demonSprite2, demonSprite3, demonSprite4 rl.Texture2D // demon sprite
	demonSrc                                              rl.Rectangle // demon source rectangle
	demonDest                                             rl.Rectangle // demon destination rectangle

	/* Sprites end */
)

func Initialize() {
	p.Initialize(screenWidth, screenHeight, worldMapX, worldMapY, int(numberOfTilesWide), int(numberOfTilesHigh))
	em.Initialize(p.GetPlayer())
	rl.InitWindow(screenWidth, screenHeight, "Svada V.1")
	playerSprite = rl.LoadTexture("./sprites/characters/character.png")
	grassSprite = rl.LoadTexture("./sprites/tiles/grassHills.png")
	biomeSprite = rl.LoadTexture("./sprites/objects/biome.png")
	stoneSprite = rl.LoadTexture("./sprites/objects/Rock.png")
	demonSprite = rl.LoadTexture("./sprites/characters/enemies/demon/walk_01.png")
	rl.SetTargetFPS(60)
	world.Initialize()
}

func Render() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.SkyBlue)

	RenderGround()
	RenderPlayer()
	RenderEnemies()

	rl.EndDrawing()
	// update cloud position
}

func RenderGround() {
	world := world.GetWorldSrc()

	var renderPositionX int32 = 0
	var renderPositionY int32 = 0
	for i := worldMapX; i < numberOfTilesWide+worldMapX; i++ {

		if (len(world) - 1) < int(i) {
			continue
		}
		for j := worldMapY; j < numberOfTilesHigh+worldMapY; j++ {
			if (len(world[i]) - 1) < int(j) {
				continue
			}
			tempValue := world[i][j].GetEdgeValue()
			switch tempValue {
			case 0:
				grassSrc = rl.NewRectangle(14, 10, 20, 20)
			case 1: // left
				grassSrc = rl.NewRectangle(0, 14, 24, 18)
			case 2: // right
				grassSrc = rl.NewRectangle(24, 14, 24, 18)
			case 4: // top
				grassSrc = rl.NewRectangle(16, 0, 18, 24)
			case 5:
				// top and left
				grassSrc = rl.NewRectangle(0, 0, 24, 24)
			case 6: // top and right
				grassSrc = rl.NewRectangle(24, 0, 24, 24)
			case 8: // bottom
				grassSrc = rl.NewRectangle(16, 24, 18, 24)
			case 9: // bottom and left
				grassSrc = rl.NewRectangle(0, 24, 24, 24)
			case 10: // bottom and right
				grassSrc = rl.NewRectangle(24, 24, 24, 24)
			default:
				grassSrc = rl.NewRectangle(14, 10, 20, 20)
			}

			grassDest = rl.NewRectangle(float32(renderPositionX*worldTileSize), float32(renderPositionY*worldTileSize), float32(worldTileSize), float32(worldTileSize))
			grassTint := 255

			if j%2 == 0 {
				grassTint -= 10
			}
			if i%2 == 0 {
				grassTint -= 10
			}
			// edges do not have tint
			if i == 0 || j == 0 {
				grassTint = 255
			}
			maxWidth := int32(len(world)) - 1
			maxHeight := int32(len(world[i])) - 1
			if i == maxWidth || j == maxHeight {
				grassTint = 255
			}
			newTint := rl.NewColor(uint8(grassTint), uint8(grassTint), uint8(grassTint), 255)
			rl.DrawTexturePro(grassSprite, grassSrc, grassDest, rl.Vector2{X: 0, Y: 0}, 0, newTint)
			if world[i][j].HasStone {
				// stoneSrc = rl.NewRectangle(81, 67, 15, 13)
				stoneSrc = rl.NewRectangle(0, 0, 128, 128)
				stoneDest = rl.NewRectangle(float32(renderPositionX*worldTileSize), float32(renderPositionY*worldTileSize), float32(worldTileSize), float32(worldTileSize))
				rl.DrawTexturePro(stoneSprite, stoneSrc, stoneDest, rl.Vector2{X: 0, Y: 0}, 0, rl.White)
			}

			if world[i][j].HasTree {
				treeSrc = rl.NewRectangle(20, 0, 26, 31)
				treeDest = rl.NewRectangle(float32(renderPositionX*worldTileSize), float32(renderPositionY*worldTileSize), float32(worldTileSize), float32(worldTileSize))
				rl.DrawTexturePro(biomeSprite, treeSrc, treeDest, rl.Vector2{X: 0, Y: 0}, 0, rl.White)
			}

			renderPositionY++
		}
		renderPositionY = 0
		renderPositionX++
	}
}

func UpdateEnemies() {
	enemies := em.GetEnemyList()
	playerPosX, playerPosY := p.GetPlayerPos()
	for i := 0; i < len(enemies); i++ {
		enemies[i].Update(playerPosX, playerPosY, p.GetWidth(), p.GetHeight())
	}
}

func RenderPlayer() {
	p.IncreaseFrame()
	if p.GetFrame() > 14 {
		p.ResetFrame()
		playerSpeed := p.GetSpeed()

		switch p.GetFacing() {
		case 0: // down
			if playerSpeed != 0 {
				// Moving
				if p.GetSourceFrame() == 1 {
					p.SetSourceFrame(0)
					playerSrc = rl.NewRectangle(113, 16, 14, 16)
				} else {
					p.SetSourceFrame(1)
					playerSrc = rl.NewRectangle(161, 16, 14, 16)
				}
			} else {
				if p.GetSourceFrame() == 1 {
					p.SetSourceFrame(0)
					playerSrc = rl.NewRectangle(16, 16, 14, 16)
				} else {
					p.SetSourceFrame(1)
					playerSrc = rl.NewRectangle(65, 16, 14, 16)
				}
			}
		case 1: // up
			if playerSpeed != 0 {
				// Moving
				if p.GetSourceFrame() == 1 {
					p.SetSourceFrame(0)
					playerSrc = rl.NewRectangle(113, 64, 14, 16)
				} else {
					p.SetSourceFrame(1)
					playerSrc = rl.NewRectangle(161, 64, 14, 16)
				}
			} else {
				if p.GetSourceFrame() == 1 {
					p.SetSourceFrame(0)
					playerSrc = rl.NewRectangle(16, 64, 14, 16)
				} else {
					p.SetSourceFrame(1)
					playerSrc = rl.NewRectangle(65, 64, 14, 16)
				}
			}
		case 2: // left
			if playerSpeed != 0 {
				// Moving
				if p.GetSourceFrame() == 1 {
					p.SetSourceFrame(0)
					playerSrc = rl.NewRectangle(113, 112, 14, 16)
				} else {
					p.SetSourceFrame(1)
					playerSrc = rl.NewRectangle(161, 112, 14, 16)
				}
			} else {
				if p.GetSourceFrame() == 1 {
					p.SetSourceFrame(0)
					playerSrc = rl.NewRectangle(16, 112, 14, 16)
				} else {
					p.SetSourceFrame(1)
					playerSrc = rl.NewRectangle(65, 112, 14, 16)
				}
			}
		case 3: // right
			if playerSpeed != 0 {
				// Moving
				if p.GetSourceFrame() == 1 {
					p.SetSourceFrame(0)
					playerSrc = rl.NewRectangle(113, 160, 14, 16)
				} else {
					p.SetSourceFrame(1)
					playerSrc = rl.NewRectangle(161, 160, 14, 16)
				}
			} else {
				if p.GetSourceFrame() == 1 {
					p.SetSourceFrame(0)
					playerSrc = rl.NewRectangle(16, 160, 14, 16)
				} else {
					p.SetSourceFrame(1)
					playerSrc = rl.NewRectangle(65, 160, 14, 16)
				}
			}
		}
	}

	playerPosX, playerPosY := p.GetPlayerPos()
	playerDest = rl.NewRectangle(float32(playerPosX), float32(playerPosY), float32(p.GetWidth()), float32(p.GetHeight()))
	rl.DrawTexturePro(playerSprite, playerSrc, playerDest, rl.NewVector2(0, 0), 0, rl.White)
}

func RenderEnemies() {
	enemies := em.GetEnemyList()
	for _, enemy := range enemies {
		enemySrc := rl.NewRectangle(260, 200, 420, 558)
		enemyDest := rl.NewRectangle(float32(enemy.GetX()), float32(enemy.GetY()), float32(enemy.GetWidth()), float32(enemy.GetHeight()))
		rl.DrawTexturePro(demonSprite, enemySrc, enemyDest, rl.NewVector2(float32(enemy.GetWidth()/2), float32(enemy.GetHeight()/2)), 0, rl.White)

	}
}

func main() {
	Initialize()
	rl.InitAudioDevice()

	music := rl.LoadMusicStream("stream.ogg")
	// music := rl.LoadMusicStream("./music/sneak.mp3")

	rl.PlayMusicStream(music)
	rl.SetMusicVolume(music, 0.5)

	for !rl.WindowShouldClose() {
		rl.UpdateMusicStream(music) // Update music buffer with new stream data

		/* INPUT */
		if rl.IsKeyDown(rl.KeyRight) || rl.IsKeyDown(rl.KeyD) {
			p.SetFacing(3)
		} else if rl.IsKeyDown(rl.KeyLeft) || rl.IsKeyDown(rl.KeyA) {
			p.SetFacing(2)
		} else if rl.IsKeyDown(rl.KeyUp) || rl.IsKeyDown(rl.KeyW) {
			p.SetFacing(1)
		} else if rl.IsKeyDown(rl.KeyDown) || rl.IsKeyDown(rl.KeyS) {
			p.SetFacing(0)
		}

		if rl.IsKeyUp(rl.KeyRight) && rl.IsKeyUp(rl.KeyLeft) && rl.IsKeyUp(rl.KeyUp) && rl.IsKeyUp(rl.KeyDown) && rl.IsKeyUp(rl.KeyW) && rl.IsKeyUp(rl.KeyA) && rl.IsKeyUp(rl.KeyS) && rl.IsKeyUp(rl.KeyD) {
			p.SetSpeed(0)
		} else if rl.IsKeyDown(rl.KeyLeftShift) || rl.IsKeyDown(rl.KeyRightShift) {
			p.SetSpeed(8)
		} else {
			p.SetSpeed(4)
		}

		// Pause/Resume music playing
		if rl.IsKeyPressed(rl.KeyP) {
			pause = !pause

			if pause {
				rl.PauseMusicStream(music)
			} else {
				rl.ResumeMusicStream(music)
			}
		}

		if rl.IsKeyPressed(rl.KeyQ) {
			em.QuickAddEnemy(p.GetPlayerPos())
		}

		if rl.IsKeyPressed(rl.KeyE) {
			// Use an item or attack
			//p.Action()
			// Find what the player is facing
			facingEnemy, enemy := em.GetFacingEnemy(p.GetPlayer(), 50) // player and margin to hit
			if facingEnemy {
				// Attack the enemy
				fmt.Println("Attacking enemy")
				enemy.TakeDamage(1)

				if enemy.Health <= 0 {
					em.RemoveAllEnemies()
				}

			} else {
				// Use an item
				p.Action()
			}
		}

		if rl.IsKeyPressed(rl.KeyR) {
			// Remove all enemies
			em.RemoveAllEnemies()
		}

		/* UPDATE */
		p.Update()
		worldX, worldY := p.GetWorldMap()
		worldMapX = int32(worldX)
		worldMapY = int32(worldY)
		UpdateEnemies()
		/* END UPDATE */
		/* RENDER */
		Render()

	}

	rl.UnloadMusicStream(music) // Unload music stream buffers from RAM
	rl.CloseAudioDevice()       // Close audio device (music streaming is automatically stopped)

	rl.CloseWindow()
}
