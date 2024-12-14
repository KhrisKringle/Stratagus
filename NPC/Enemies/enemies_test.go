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

func TestTakeDamage(t *testing.T) {
	e := Enemies.Enemy{
		Health: 20,
	}
	e.TakeDamage(1)
	got := e.Health
	want := 19

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

// func TestDoDamage(t *testing.T) {
// 	e := Enemies.Enemy{
// 		Strength: 16,
// 		Health:   20,
// 	}
// 	wantcheck := 20 - e.DoDamage()
// 	e.TakeDamage(e.DoDamage())
// 	got := e.Health
// 	want := wantcheck

// 	if got != want {
// 		t.Errorf("got %v want %v", got, want)
// 	}

// }
