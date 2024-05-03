package item

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func init() {
	fmt.Println("Initializing item package")
}

type ItemType struct {
	Name         string
	SpriteSource rl.Rectangle
	PosX         float32
	PosY         float32
	Height       float32
	Width        float32
}

func NewItem(name string, spriteSource rl.Rectangle, posX, posY float32, width, height float32) ItemType {
	fmt.Println("Creating new item")
	item := ItemType{Name: name, SpriteSource: spriteSource, PosX: posX, PosY: posY, Height: height, Width: width}
	return item
}

func (i *ItemType) GetHeight() float32 {
	return i.Height
}

func (i *ItemType) GetWidth() float32 {
	return i.Width
}

func (i *ItemType) GetPosition() (float32, float32) {
	return i.PosX, i.PosY
}

func (i *ItemType) SetPosition(x, y float32) {
	i.PosX = x
	i.PosY = y
}

func (i *ItemType) GetSpriteSource() rl.Rectangle {
	return i.SpriteSource
}
