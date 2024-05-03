package world

import (
	"fmt"
	"game/item"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const tileSize = 50
const NumTrees = 20

var (
	trees           []item.Tree = make([]item.Tree, 0)
	treeGrowthSpeed float64     = 2
	lastTreeGrowth  float64     = 0
)

type Tile struct {
	Grass bool
}

type World struct {
	Tiles [][]Tile
	Trees []item.Tree
}

func (w *World) GetTrees() []item.Tree {
	return w.Trees
}

func (w *World) SetTrees(trees []item.Tree) {
	fmt.Println(trees)
	w.Trees = trees
}
func (w *World) RemoveTree(treePos int) {
	newTreeArray := make([]item.Tree, 0)
	for i, tree := range w.Trees {
		if i != treePos {
			newTreeArray = append(newTreeArray, tree)
		}
	}
	w.Trees = newTreeArray
}

func (w *World) Update() {
	currentTime := rl.GetTime()
	if currentTime-lastTreeGrowth > treeGrowthSpeed {
		for pos := range w.Trees {
			w.Trees[pos].Grow()
		}
		lastTreeGrowth = currentTime
	}

	// Remove dead trees
	for pos, tree := range w.Trees {
		if tree.Health <= 0 {
			w.RemoveTree(pos)
		}
	}
}

func (w *World) AddTree(x, y float32) {
	tree := item.NewTree(x, y)
	w.Trees = append(w.Trees, tree)
}

func (w *World) HitTree(treeIndex int, damage int) {
	if treeIndex >= 0 && treeIndex < len(w.Trees) {
		w.Trees[treeIndex].Hit(damage)
	}
}

func NewWorld(width, height int) *World {
	world := &World{} // Create a new world by allocating memory for it

	world.Tiles = make([][]Tile, width)
	for x := range world.Tiles {
		world.Tiles[x] = make([]Tile, height)
		for y := range world.Tiles[x] {
			// Here we can do some better logic for the world generation
			world.Tiles[x][y] = Tile{Grass: true}
		}
	}

	for i := 0; i < NumTrees; i++ { // Adjust the number of trees as needed
		margin := 50
		randX := Random(margin, (width*tileSize)-margin)
		randY := Random(margin, (height*tileSize)-margin)
		newTree := item.NewTree(float32(randX), float32(randY))
		newTree.SetFullyGrown()
		world.Trees = append(world.Trees, newTree)
	}
	return world
}

func init() {
	// This function is called when the package is imported
	fmt.Println("World package initialized")
}

func Random(min, max int) int {
	return min + rand.Intn(max-min)
}
