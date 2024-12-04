package Enemies_test

import (
	"testing"

	"github.com/KhrisKringle/Stratagus/NPC/Enemies"
)

func TestRandomAttributeSetter(t *testing.T) {

}

func TestChangeTurnState(t *testing.T) {

	e := Enemies.Enemy{
		AttackTurnState: false,
	}
	e.ChangeTurnState(e.AttackTurnState)
	got := e.AttackTurnState
	want := true

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
