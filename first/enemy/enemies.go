package enemy

import (
	p "first/player"
	"log"
	"math/rand"
)

type Enemy struct {
	Health         int
	PositionX      int
	PositionY      int
	Speed          int
	AnimationCycle int
	Width          int
	Height         int
}

var (
	enemyList []*Enemy
	player    p.Player
)

func (e *Enemy) Update(playerPosX, playerPosY, playerWidth, playerHeight int) {
	// update new enemy position

	currentSpeed := e.Speed
	widthOffset := playerWidth
	heightOffset := playerHeight
	stopDistance := 50 // distance at which the enemy should stop

	if e.PositionX > playerPosX+widthOffset+stopDistance {
		e.PositionX = e.PositionX - currentSpeed
	} else if e.PositionX < playerPosX-stopDistance {
		e.PositionX = e.PositionX + currentSpeed
	}

	if e.PositionY > playerPosY+heightOffset+stopDistance {
		e.PositionY = e.PositionY - currentSpeed
	} else if e.PositionY < playerPosY-stopDistance {
		e.PositionY = e.PositionY + currentSpeed
	}
}

func (e *Enemy) GetX() int {
	return e.PositionX
}

func (e *Enemy) GetY() int {
	return e.PositionY
}

func (e *Enemy) GetWidth() int {
	return e.Width
}

func (e *Enemy) GetHeight() int {
	return e.Height
}

type Item struct {
	Name        string
	Description string
	Effect      func(*Enemy)
}

type StatusEffect struct {
	Name        string
	Description string
	Duration    int // How many turns will it last?
	Effect      func(*Enemy)
}

func (e *Enemy) GetHealth() int {
	return e.Health
}

func (e *Enemy) GetAnimationCycle() int {
	return e.AnimationCycle
}

func (e *Enemy) SetAnimationCycle(cycle int) {
	e.AnimationCycle = cycle
}

func (e *Enemy) TakeDamage(dmg int) {
	log.Println("Enemy took damage")
	log.Println(e.Health)
	e.Health -= dmg
}

func Initialize(playerInstance p.Player) {
	enemyList = make([]*Enemy, 0)
	player = playerInstance

}

func RemoveAllEnemies() {
	enemyList = make([]*Enemy, 0)
}

// func GetFacingEnemy(player p.Player) (bool, *Enemy) {
// 	// We check if the position of the enemy is in the same position as the player is and their facing is the same
// 	// If so, then we return true, otherwise false
// 	tempEnemy := Enemy{}
// 	offsetX := 50
// 	offsetY := 50
// 	if player.Facing == 0 {
// 		offsetY = 100
// 	} else if player.Facing == 1 {
// 		offsetY = -100
// 	} else if player.Facing == 2 {
// 		offsetX = -100
// 	} else if player.Facing == 3 {
// 		offsetX = 100
// 	}

// 	for i := 0; i < len(enemyList); i++ {
// 		playerPosX, playerPosY := p.GetPlayerPos()
// 		if enemyList[i].PositionX >= playerPosX-offsetX && enemyList[i].PositionX <= playerPosX+offsetX && enemyList[i].PositionY >= playerPosY-offsetY && enemyList[i].PositionY <= playerPosY+offsetY {
// 			tempEnemy = *enemyList[i]
// 			return true, &tempEnemy
// 		}

// 		if enemyList[i].Health <= 0 {
// 			enemyList = append(enemyList[:i], enemyList[i+1:]...)
// 			p.AddExperience(1)
// 		}
// 	}

//		return false, &tempEnemy
//	}
func GetFacingEnemy(player p.Player, marginToHit int) (bool, *Enemy) {
	offsetX := 90 + marginToHit // Adjusted for margin
	offsetY := 90 + marginToHit // Adjusted for margin

	switch player.Facing {
	case 0: // Facing down
		offsetY += 10
	case 1: // Facing up
		offsetY -= 10
	case 2: // Facing left
		offsetX -= 10
	case 3: // Facing right
		offsetX += 10
	}

	for _, enemy := range enemyList {
		playerPosX, playerPosY := p.GetPlayerPos()
		if enemy.PositionX >= playerPosX-offsetX && enemy.PositionX <= playerPosX+offsetX && enemy.PositionY >= playerPosY-offsetY && enemy.PositionY <= playerPosY+offsetY {
			// Found the facing enemy
			return true, enemy
		}
	}

	// No facing enemy found
	return false, nil
}

func QuickAddEnemy(posX, posY int) {
	AddEnemy(10, 10, posX, posY)
}

func AddEnemy(health, defense, positionX, positionY int) {
	enemyList = append(enemyList, &Enemy{
		Health:         health,
		PositionX:      positionX + 100,
		AnimationCycle: 0,
		Width:          80,
		Height:         100,
		PositionY:      positionY + 100,
		Speed:          rand.Intn(3) + 1,
	})
}

func GetEnemy(index int) Enemy {
	return *enemyList[index]
}

func GetEnemyList() []*Enemy {
	return enemyList
}
