package combat

import (
	"math/rand"
	"github.com/KhrisKringle/Stratagus/player"
	"github.com/KhrisKringle/Stratagus/NPC/Enemies"
)

func Attack(a, d int, p player.Player, e Enemies.Enemy) int {
	attackRoll := rand.Intn(a)
	defenseRoll := rand.Intn(d)

	if attackRoll + p.Strength > defenseRoll + e.Constitution {
		return attackRoll + p.Strength - defenseRoll + e.Constitution
	} else {
		return 0
	}
}
