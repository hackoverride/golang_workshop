package player

type PlayerType struct {
	PosX      float32
	PosY      float32
	Speed     float32
	RenderPos int
}

func NewPlayer(posX, posY, speed float32) PlayerType {
	player := PlayerType{PosX: posX, PosY: posY, Speed: speed, RenderPos: 0}
	return player
}

func (p *PlayerType) GetRenderPos() int {
	return p.RenderPos
}

func (p *PlayerType) SetRenderPos(newCycle int) {
	p.RenderPos = newCycle
}

func (p *PlayerType) GetSpeed() float32 {
	return p.Speed
}

func (p *PlayerType) Move(x, y float32) {
	p.PosX += x
	p.PosY += y
}

func (p *PlayerType) GetPosition() (float32, float32) {
	return p.PosX, p.PosY
}
