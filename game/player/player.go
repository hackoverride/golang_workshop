package player

type PlayerType struct {
	PosX         float32
	PosY         float32
	Width        float32
	Height       float32
	Speed        float32
	RenderPos    int
	IsMoving     bool
	IsUsingSpell bool
	FaceRight    bool
}

func NewPlayer(posX, posY, speed float32) PlayerType {
	player := PlayerType{
		PosX:         posX,
		PosY:         posY,
		Speed:        speed,
		RenderPos:    0,
		Width:        40,
		Height:       60,
		IsMoving:     false,
		IsUsingSpell: false,
		FaceRight:    true,
	}
	return player
}

func (p *PlayerType) SetPlayerUsingSpell(usingSpell bool) {
	if p.IsUsingSpell == usingSpell {
		return
	}
	p.IsUsingSpell = usingSpell
}

func (p *PlayerType) GetPlayerWidth() float32 {
	return p.Width
}

func (p *PlayerType) GetPlayerHeight() float32 {
	return p.Height
}

func (p *PlayerType) IsPlayerUsingSpell() bool {
	return p.IsUsingSpell
}

func (p *PlayerType) IsPlayerMoving() bool {
	return p.IsMoving
}

func (p *PlayerType) SetPlayerFaceRight(faceRight bool) {
	p.FaceRight = faceRight
}

func (p *PlayerType) IsPlayerFaceRight() bool {
	return p.FaceRight
}

func (p *PlayerType) SetPlayerMoving(moving bool) {
	if p.IsMoving == moving {
		return
	}
	p.IsMoving = moving
	p.RenderPos = 0
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
