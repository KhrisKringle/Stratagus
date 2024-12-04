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
	}

	got := p.DeckSetter(p.Race)
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

/*func TestChangeTurnState(t *testing.T) {

	p := player.Player{
		AttackTurnState: false,
	}
	got := player.ChangeTurnState(p.AttackTurnState)
	want := true

	if got != want {
		t.Errorf("got %b want %b", got, want)
	}
}*/

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
