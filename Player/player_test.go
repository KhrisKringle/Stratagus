package Player

import "testing"

func TestModifier(t *testing.T) {
	p := Player{
		Strength:     20,
		Dexterity:    18,
		Constitution: 14,
		Intelligence: 10,
		Wisdom:       24,
		Charisma:     8,
		Inventory:    map[string]float32{
			"sword": 15.2,
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
		Inventory:    map[string]float32{
			"sword": 15.2,
			"potion": 0.5,
		},
	}
	got := p.WeightChecker(p.Inventory)
	want := true

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}