package player

type PlayerType struct {
	PosX float32
	PosY float32
}

func NewPlayer(posX, posY float32) *PlayerType {
	player := &PlayerType{PosX: posX, PosY: posY}
	return player
}

func (p *PlayerType) Move(x, y float32) {
	p.PosX += x
	p.PosY += y
}

func (p *PlayerType) GetPosition() (float32, float32) {
	return p.PosX, p.PosY
}
