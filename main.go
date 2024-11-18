package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/KhrisKringle/Stratagus/Player"
)

func main() {
	worldMap := [][]string{
		{"water", "land", "land", "water", "water"},
		{"land", "land", "land", "land", "water"},
		{"water", "land", "water", "land", "land"},
		{"water", "water", "land", "land", "land"},
		{"water", "water", "water", "water", "water"},
	}

	player_postition = worldMap[2][1]

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Pick a race (Elf, Human, Orc, Gnome, Trent, Dragonkin): ")

	race, _ := reader.ReadString('\n')

	p := Player.Player{
		Race:         race,
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
	p.Deck = p.DeckSetter(p.Race)

}
