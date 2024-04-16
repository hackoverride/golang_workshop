package trees

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	treeHeight = 220
	treeWidth  = 180
)

var (
	trees              = []Tree{}
	screenWidth  int32 = 1200
	screenHeight int32 = 800
	groundHeight int32 = 100
)

type Tree struct {
	Width       int32
	Height      int32
	PosX        int32
	PosY        int32
	Speed       int32
	SpriteFrame int32
	Frame       int32
}

func (tree *Tree) Update() {
	tree.SpriteFrame++

	if tree.PosX > screenWidth+tree.Width+25 {
		tree.PosX = -tree.Width - 25
	}

	if tree.PosX < -tree.Width-25 {
		tree.PosX = screenWidth + 25
	}

	tree.PosX += tree.Speed
}

func AddTree() []Tree {
	var overlap int32 = 20
	tree := Tree{
		Width:       treeWidth,
		Height:      treeHeight,
		PosX:        -treeWidth - 20,
		PosY:        screenHeight - groundHeight - treeHeight + overlap,
		Speed:       1,
		SpriteFrame: 0,
		Frame:       0,
	}

	if rl.GetRandomValue(0, 100) > 50 {
		tree.Frame = 1
	}

	trees = append(trees, tree)

	return trees
}
