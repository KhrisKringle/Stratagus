package Enemies

import (
	"math/rand"
)

type Enemy struct {
	Race            string
	Resistances     []string
	Strength        int
	Dexterity       int
	Constitution    int
	Health          int
	Alive           bool
	Inventory       map[string]float32
	AttackTurnState bool
}

func (e *Enemy) RandomAttributeSetter() {

	attributeHolder := make([]int, 0)

	for i := 0; i <= 2; i++ {
		attributeRandomNumber := rand.Intn(19)
		attributeNumber := attributeRandomNumber + 1 // Have to add one because it starts at 0
		attributeHolder = append(attributeHolder, attributeNumber)
	}
	e.Strength = attributeHolder[0]
	e.Dexterity = attributeHolder[1]
	e.Constitution = attributeHolder[2]
}

func (e Enemy) Modifier(attribute string) int {
	switch attribute {
	case "Strength":
		return (e.Strength - 10) / 2
	case "Dexterity":
		return (e.Dexterity - 10) / 2
	case "Constitution":
		return (e.Constitution - 10) / 2
	default:
		return 0
	}
}

func (e *Enemy) ChangeTurnState(et bool) {

	if et == true {
		e.AttackTurnState = false
	} else {
		e.AttackTurnState = true
	}
}

func (e Enemy) DoDamage() int {
	attrole := rand.Intn(20)

	withmod := attrole + e.Modifier("Strength")
	return withmod
}

func (e *Enemy) TakeDamage(dmg int) {
	if e.Health != 0 {
		e.Health -= dmg
		return
	} else {
		e.Alive = false
		return
	}
}
