package Enemies_test

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

func TestChangeTurnState(t *testing.T) {

	p := Enemy.enemy{
		AttackTurnState: false,
	}
	got := Enemy.ChangeTurnState(p.AttackTurnState)
	want := true

	if got != want {
		t.Errorf("got %b want %b", got, want)
	}
}
