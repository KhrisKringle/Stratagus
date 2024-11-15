package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/KhrisKringle/Stratagus/Player"
)

func main() {
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
