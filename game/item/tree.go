package item

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	MaxHeight = 100
	MaxWidth  = 60
)

var (
	TreeSpriteSource []rl.Rectangle
)

type Tree struct {
	Name     string
	Quantity int
	Growth   int // 0-100
	PosX     float32
	PosY     float32
	Height   float32
	Width    float32
	Health   int
}

func NewTree(posX, posY float32) Tree {
	fmt.Println("Creating new tree")
	fmt.Println("Tree position: ", posX, posY)
	tree := Tree{Name: "Tree", Quantity: 10, PosX: posX, PosY: posY, Height: 10, Width: 20, Health: 100}
	return tree
}

func (t *Tree) GetTreeSprite() rl.Rectangle {
	if t.Growth < 25 {
		return TreeSpriteSource[0]
	} else if t.Growth < 50 {
		return TreeSpriteSource[1]
	} else if t.Growth < 75 {
		return TreeSpriteSource[2]
	} else {
		return TreeSpriteSource[3]
	}
}

func (t *Tree) GetTreeWidth() float32 {
	return t.Width
}

func (t *Tree) GetTreeHeight() float32 {
	return t.Height
}

func (t *Tree) GetPosition() (float32, float32) {
	return t.PosX, t.PosY
}

func (t *Tree) GetHealth() int {
	return t.Health
}

func (t *Tree) GetTreeGrowth() int {
	return t.Growth
}

func (t *Tree) SetFullyGrown() {
	t.Growth = 100
	t.Height = MaxHeight
	t.Width = MaxWidth
}

func (t *Tree) Grow() {
	if t.Growth >= 100 {
		return
	}
	fmt.Println("Growing tree")
	fmt.Println("Tree growth: ", t.Growth)

	t.Growth += 1
	if t.Growth < 25 {
		t.Height = 10
		t.Width = 20
	} else if t.Growth < 40 {
		t.Height = 15
		t.Width = 20
	} else if t.Growth < 50 {
		t.Height = 25
		t.Width = 20
	} else if t.Growth < 60 {
		t.Height = 35
		t.Width = 30
	} else if t.Growth < 75 {
		t.Height = 50
		t.Width = 40
	} else {
		t.Height = 100
		t.Width = 60
	}
}

func (t *Tree) Hit(dmg int) {
	t.Health -= dmg
}

func init() {
	TreeSpriteSource = []rl.Rectangle{
		rl.NewRectangle(83, 21, 10, 8),
		rl.NewRectangle(16, 48, 16, 15),
		rl.NewRectangle(0, 0, 15, 28),
		rl.NewRectangle(19, 0, 25, 31),
	}
}
