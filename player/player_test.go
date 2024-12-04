package player_test

import (
	"testing"

	"github.com/KhrisKringle/Stratagus/player"
)

func TestModifier(t *testing.T) {
	p := player.Player{
		Strength: 20,
	}

	got := p.Modifier("Strength")
	var want int = 5

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestWeightChecker(t *testing.T) {

	p := player.Player{
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

	p := player.Player{
		Race: "Human",
		Deck: make([]player.Spell, 0),
	}

	p.DeckSetter(p.Race)
	got := p.Deck
	want := []player.Spell{
		{player.HolyDamage, 3},
		{player.HolyDamage, 4},
		{player.ForceDamage, 8},
		{player.ForceDamage, 7},
		{player.ForceDamage, 9}}

	if !compareSlices(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
func TestPlayerPosition(t *testing.T) {

	p := player.Player{
		PlayerPos_Y: 1,
		PlayerPos_X: 0,
	}

	got := player.PlayerPositionChecker(player.WorldMap[p.PlayerPos_Y][p.PlayerPos_X])
	want := true

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestPlayerMove(t *testing.T) {

	p := player.Player{
		PlayerPos_Y: 2,
		PlayerPos_X: 1,
	}

	p.PlayerMove()
	got := player.PlayerPositionChecker(player.WorldMap[p.PlayerPos_Y][p.PlayerPos_X])
	want := true

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestChangeTurnState(t *testing.T) {

	p := player.Player{
		AttackTurnState: false,
	}
	p.ChangeTurnState(p.AttackTurnState)
	got := p.AttackTurnState
	want := true

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

// Need this function to compare the slices
func compareSlices(slice1, slice2 []player.Spell) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}

	}

	return true
}
