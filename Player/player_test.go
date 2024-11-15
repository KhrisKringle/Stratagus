package Player

import (
	"testing"
)

func TestModifier(t *testing.T) {
	p := Player{
		Strength:     20,
		Dexterity:    18,
		Constitution: 14,
		Intelligence: 10,
		Wisdom:       24,
		Charisma:     8,
		Inventory: map[string]float32{
			"sword":  15.2,
			"potion": 0.5,
		},
	}
	got := p.Modifier("Strength")
	var want int = 5

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestWeightChecker(t *testing.T) {
	p := Player{
		Strength:     20,
		Dexterity:    18,
		Constitution: 14,
		Intelligence: 10,
		Wisdom:       24,
		Charisma:     8,
		Inventory: map[string]float32{
			"sword":  15.2,
			"potion": 0.5,
		},
	}
	got := p.WeightChecker(p.Inventory)
	want := true

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestDeckSetter(t *testing.T) {
	p := Player{
		Race:         "Human",
		Resistances:  make([]string, 5),
		Journal:      make([]string, 0),
		Strength:     20,
		Dexterity:    18,
		Constitution: 14,
		Intelligence: 10,
		Wisdom:       24,
		Charisma:     8,
		Inventory: map[string]float32{
			"sword":  15.2,
			"potion": 0.5,
		},
	}

	got := p.DeckSetter("Human")
	want := []map[string]int{
		{"Holy": 3},
		{"Holy": 4},
		{"Force": 8},
		{"Force": 7},
		{"Force": 9},
	}

	if !compareSlices(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

// Need this function to compare the slices
func compareSlices(slice1, slice2 []map[string]int) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	for i := range slice1 {
		if len(slice1[i]) != len(slice2[i]) {
			return false
		}

		for key, value1 := range slice1[i] {
			value2, ok := slice2[i][key]
			if !ok || value1 != value2 {
				return false
			}
		}
	}

	return true
}
