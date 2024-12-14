package combat

import (
	"math/rand"

	"github.com/KhrisKringle/Stratagus/NPC/Enemies"
	"github.com/KhrisKringle/Stratagus/player"
)

func Attack(a, d int, p player.Player, e Enemies.Enemy) int {
	attackRoll := rand.Intn(a + 1)
	defenseRoll := rand.Intn(d + 1)

	if attackRoll+p.Modifier("Strength") > defenseRoll+e.Modifier("Constitution") {
		return attackRoll + p.Modifier("Strength") - defenseRoll + e.Modifier("Constitution")
	} else {
		return 0
	}
}
