package enemies

import (
	"math/rand"
	"fmt"
)

type enemy struct {
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

func (e *enemy) RandomAttributeSetter() int {

	attributeHolder := make([]int, 0)

    for i := 0; i <= 5; i++ {
		attributeRandomNumber := rand.Intn(19)
    	attributeNumber := attributeRandomNumber + 1
		attributeHolder = append(attributeHolder, attributeNumber)
	}
	e.Strength = attributeHolder[0]
	e.Dexterity = attributeHolder[1]
	e.Constitution = attributeHolder[2]
	e.Intelligence = attributeHolder[3]
	e.Wisdom = attributeHolder[4]
	e.Charisma = attributeHolder[5]


		
}