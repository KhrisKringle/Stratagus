package Nutrals

import (
	"math/rand"
)

type Neutral struct {
	Race         string
	Resistances  []string
	Strength     int
	Dexterity    int
	Constitution int
	Charisma     int
	Inventory    map[string]float32
}

func (n *Neutral) RandomAttributeSetter(){

	attributeHolder := make([]int, 0)

    for i := 0; i <= 3; i++ {
		attributeRandomNumber := rand.Intn(19)
    	attributeNumber := attributeRandomNumber + 1
		attributeHolder = append(attributeHolder, attributeNumber)
	}
	n.Strength = attributeHolder[0]
	n.Dexterity = attributeHolder[1]
	n.Constitution = attributeHolder[2]
	n.Charisma = attributeHolder[3]
}