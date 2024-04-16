package world

import "math/rand"

const (
	worldWidth  = 100
	worldHeight = 100
)

type WorldTile struct {
	posX        int32
	posY        int32
	SpriteFrame int32 //
	HasTree     bool
	HasCow      bool
	HasStone    bool
	Void        bool
}

var worldSrc = [][]WorldTile{} // world source rectangle

func (tile *WorldTile) GetEdgeValue() int {
	// the edge value is the sum of the values of the 8 tiles surrounding the tile (bitwise)
	// left edge value == 1
	// right edge value == 2
	// top edge value == 4
	// bottom edge value == 8

	edgeValue := 0 // 0 = no edge
	// get all edge values
	if tile.posX == 0 {
		// left edge
		edgeValue += 1
	}
	if tile.posX == worldWidth-1 {
		// right edge
		edgeValue += 2
	}
	if tile.posY == 0 {
		// top edge
		edgeValue += 4
	}
	if tile.posY == worldHeight-1 {
		// bottom edge
		edgeValue += 8
	}

	return edgeValue
}

func Initialize() {
	for i := 0; i < worldWidth; i++ {
		worldSrc = append(worldSrc, []WorldTile{})
		for j := 0; j < worldHeight; j++ {
			newTile := WorldTile{
				posX:        int32(i),
				posY:        int32(j),
				SpriteFrame: 0,
				HasTree:     false,
				HasCow:      false,
				HasStone:    false,
				Void:        false,
			}

			randomNumber := rand.Intn(60)
			// should not generate trees on the edges
			if (i == 0 || i == worldWidth-1) || (j == 0 || j == worldHeight-1) {
				randomNumber = -1
			}
			if randomNumber >= 55 {
				newTile.HasTree = true
			}
			if randomNumber == 2 {
				newTile.HasCow = true
			}
			if randomNumber == 3 {
				newTile.HasStone = true
			}
			worldSrc[i] = append(worldSrc[i], newTile)
		}
	}

}

func GetWorldSrc() [][]WorldTile {
	return worldSrc
}

func GetWorldWidth() int {
	return worldWidth
}

func GetWorldHeight() int {
	return worldHeight
}
