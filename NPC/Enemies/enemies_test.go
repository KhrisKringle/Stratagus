package enemies

import (
	"testing"
)

func TestRandomAttributeSetter(t *testing.T) {
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


}