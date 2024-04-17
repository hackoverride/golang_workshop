package world

import "fmt"

type Tile struct {
	Grass bool
}

type World struct {
	Tiles [][]Tile
}

func NewWorld(width, height int) *World {
	world := &World{} // Create a new world by allocating memory for it

	world.Tiles = make([][]Tile, width)
	for x := range world.Tiles {
		world.Tiles[x] = make([]Tile, height)
		for y := range world.Tiles[x] {
			// Here we can do some better logic for the world generation.
			world.Tiles[x][y] = Tile{Grass: true}
		}
	}
	return world
}

func init() {
	// This function is called when the package is imported
	fmt.Println("World package initialized")
}
