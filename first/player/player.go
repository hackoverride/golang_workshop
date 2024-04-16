package player

import (
	"first/world"
	"fmt"
)

var (
	screenHeight      int
	screenWidth       int
	worldMapX         int = 0
	worldMapY         int = 0
	numberOfTilesHigh int
	numberOfTilesWide int
	player            Player
)

type Item struct {
	Name        string
	Description string
	Quantity    int
}

type Player struct {
	PosX       int
	PosY       int
	Width      int
	Height     int
	Health     int // 0 = dead, 100 = full health
	Defense    int
	Facing     int
	Speed      int
	Experience int
	Frame      int
	Source     int
	Inventory  []Item
}

func GetPlayer() Player {
	return player
}

func (p *Player) Update() {
	// 0 = down, 1 = up, 2 = left, 3 = right
	switch p.Facing {
	case 0:
		if p.Speed != 0 {
			p.PosY += p.Speed
		}
	case 1:
		if p.Speed != 0 {
			p.PosY -= p.Speed
		}
	case 2:
		if p.Speed != 0 {
			p.PosX -= p.Speed
		}
	case 3:
		if p.Speed != 0 {
			p.PosX += p.Speed
		}
	}

	// If the player is closing in on the edge, then move the world.
	if player.PosX > screenWidth-(player.Width) {
		if worldMapX < (world.GetWorldWidth())-numberOfTilesWide {
			worldMapX = worldMapX + numberOfTilesWide
			player.PosX = 0 + (player.Width)
		}
	}

	if player.PosX <= 0 {
		if worldMapX-numberOfTilesWide >= 0 {
			worldMapX = worldMapX - numberOfTilesWide
			player.PosX = screenWidth - (player.Width)
		}
	}

	if player.PosY > screenHeight-(player.Height) {
		if worldMapY < world.GetWorldHeight()-numberOfTilesHigh {
			worldMapY = worldMapY + numberOfTilesHigh
			player.PosY = 0
		}
	}
	if player.PosY < 0 {
		if worldMapY-numberOfTilesHigh >= 0 {
			worldMapY = worldMapY - numberOfTilesHigh
			player.PosY = screenHeight - (player.Height)
		}
	}
}

func Action() {
	fmt.Println("Player action")
	switch player.Facing {
	case 0:
		fmt.Println("Player facing down")
	case 1:
		fmt.Println("Player facing up")
	case 2:
		fmt.Println("Player facing left")
	case 3:
		fmt.Println("Player facing right")
	}
}

func AddExperience(exp int) {
	player.Experience += exp
}

func GetWorldMap() (int, int) {
	return worldMapX, worldMapY
}

func (p *Player) AddToInventory() {

}

func Initialize(newScreenWidth, newScreenHeight int, newWorldMapX, newWorldMapY int32, newNumberOfTilesWide, newNumberOfTilesHigh int) {
	player.Facing = 0
	player.Speed = 0
	player.Frame = 0
	player.Source = 0
	player.Width = 70
	player.Height = 70
	screenHeight = newScreenHeight
	screenWidth = newScreenWidth
	player.PosX = int(screenWidth)/2 - player.Width
	player.PosY = int(screenHeight)/2 - player.Height

	player.Experience = 0
	player.Inventory = make([]Item, 0)
	player.Defense = 0
	player.Health = 10
	worldMapX = int(newWorldMapX)
	worldMapY = int(newWorldMapY)
	numberOfTilesWide = newNumberOfTilesWide
	numberOfTilesHigh = newNumberOfTilesHigh
}

func GetSourceFrame() int {
	return player.Source
}

func SetSourceFrame(source int) {
	player.Source = source
}

func Update() {
	player.Update()
}

func GetWidth() int {
	return player.Width
}

func GetHeight() int {
	return player.Height
}

func IncreaseFrame() {
	player.Frame++
}

func ResetFrame() {
	player.Frame = 0
}

func GetFrame() int {
	return player.Frame
}

func SetFacing(facing int) {
	player.Facing = facing
}

func SetSpeed(speed int) {
	player.Speed = speed
}

func GetPlayerPos() (int, int) {
	return player.PosX, player.PosY
}

func GetSpeed() int {
	return player.Speed
}

func GetFacing() int {
	return player.Facing
}
