package item

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	FireSpriteSource []rl.Rectangle
)

type FireLight struct {
	RenderPos int
	PosX      float32
	PosY      float32
	Height    float32
	Width     float32
}

func NewFireLight(posX, posY float32) FireLight {
	fireLight := FireLight{PosX: posX, PosY: posY, RenderPos: 0, Height: 50, Width: 50}

	return fireLight
}

func (f *FireLight) GetFireWidth() float32 {
	return f.Width
}

func (f *FireLight) GetFireHeight() float32 {
	return f.Height
}

func (f *FireLight) GetPosition() (float32, float32) {
	return f.PosX, f.PosY
}

func (f *FireLight) GetRenderPos() int {
	return f.RenderPos
}

func (f *FireLight) SetRenderPos(renderPos int) {
	f.RenderPos = renderPos
}

func init() {
	FireSpriteSource = []rl.Rectangle{
		rl.NewRectangle(25, 25, 50, 50),
		rl.NewRectangle(125, 25, 50, 50),
		rl.NewRectangle(225, 25, 50, 50),
		rl.NewRectangle(325, 25, 50, 50),
		rl.NewRectangle(425, 25, 50, 50),
		rl.NewRectangle(525, 25, 50, 50),
		rl.NewRectangle(625, 25, 50, 50),
		rl.NewRectangle(25, 125, 50, 50),
		rl.NewRectangle(125, 125, 50, 50),
		rl.NewRectangle(225, 125, 50, 50),
		rl.NewRectangle(325, 125, 50, 50),
		rl.NewRectangle(425, 125, 50, 50),
		rl.NewRectangle(525, 125, 50, 50),
		rl.NewRectangle(625, 125, 50, 50),
		rl.NewRectangle(25, 225, 50, 50),
		rl.NewRectangle(125, 225, 50, 50),
		rl.NewRectangle(225, 225, 50, 50),
		rl.NewRectangle(325, 225, 50, 50),
		rl.NewRectangle(425, 225, 50, 50),
		rl.NewRectangle(525, 225, 50, 50),
		rl.NewRectangle(625, 225, 50, 50),
		rl.NewRectangle(25, 325, 50, 50),
		rl.NewRectangle(125, 325, 50, 50),
		rl.NewRectangle(225, 325, 50, 50),
		rl.NewRectangle(325, 325, 50, 50),
		rl.NewRectangle(425, 325, 50, 50),
		rl.NewRectangle(525, 325, 50, 50),
		rl.NewRectangle(625, 325, 50, 50),
		rl.NewRectangle(25, 425, 50, 50),
		rl.NewRectangle(125, 425, 50, 50),
		rl.NewRectangle(225, 425, 50, 50),
		rl.NewRectangle(325, 425, 50, 50),
		rl.NewRectangle(425, 425, 50, 50),
		rl.NewRectangle(525, 425, 50, 50),
		rl.NewRectangle(625, 425, 50, 50),
		rl.NewRectangle(25, 525, 50, 50),
		rl.NewRectangle(125, 525, 50, 50),
		rl.NewRectangle(225, 525, 50, 50),
		rl.NewRectangle(325, 525, 50, 50),
		rl.NewRectangle(425, 525, 50, 50),
		rl.NewRectangle(525, 525, 50, 50),
		rl.NewRectangle(625, 525, 50, 50),
		rl.NewRectangle(25, 625, 50, 50),
		rl.NewRectangle(125, 625, 50, 50),
		rl.NewRectangle(225, 625, 50, 50),
		rl.NewRectangle(325, 625, 50, 50),
		rl.NewRectangle(425, 625, 50, 50),
		rl.NewRectangle(525, 625, 50, 50),
		rl.NewRectangle(625, 625, 50, 50),
	}
}
