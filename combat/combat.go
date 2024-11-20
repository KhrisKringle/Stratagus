package combat

import (
	"math/rand"
)

func Attack(a, d int) int {
	attackRoll := rand.Intn(a)
	defenseRoll := rand.Intn(d)

	if attackRoll > defenseRoll {
		return attackRoll - defenseRoll
	} else {
		return 0
	}
}



