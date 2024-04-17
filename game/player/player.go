package player

type PlayerType struct {
	PosX  float32
	PosY  float32
	Speed float32
}

func NewPlayer(posX, posY float32) *PlayerType {
	player := &PlayerType{PosX: posX, PosY: posY, Speed: 2}
	return player
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
