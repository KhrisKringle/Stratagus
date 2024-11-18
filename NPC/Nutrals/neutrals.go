package nutrals

import (
	"math/rand"
)

type neutral struct {
	Race         string
	Resistances  []string
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Inventory    map[string]float32
}

func (n *neutral) RandomAttributeSetter(){

	attributeHolder := make([]int, 0)

    for i := 0; i <= 5; i++ {
		attributeRandomNumber := rand.Intn(19)
    	attributeNumber := attributeRandomNumber + 1
		attributeHolder = append(attributeHolder, attributeNumber)
	}
	n.Strength = attributeHolder[0]
	n.Dexterity = attributeHolder[1]
	n.Constitution = attributeHolder[2]
	n.Intelligence = attributeHolder[3]
	n.Wisdom = attributeHolder[4]
	n.Charisma = attributeHolder[5]
}