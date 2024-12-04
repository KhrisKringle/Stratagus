package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"

	"github.com/KhrisKringle/Stratagus/NPC/Enemies"
	"github.com/KhrisKringle/Stratagus/NPC/Nutrals"
	"github.com/KhrisKringle/Stratagus/player"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Pick a race (Elf, Human, Orc, Gnome, Trent, Dragonkin): ")

	var race_list = []string{"Elf", "Human", "Orc", "Gnome", "Trent", "Dragonkin"}

	randRace := rand.Intn(5)

	race, _ := reader.ReadString('\n')

	// Remove the newline character
	race = strings.TrimSpace(race)

	p := player.Player{
		Race:         race,
		Resistances:  make([]string, 0, 5),
		Journal:      make([]string, 0),
		Strength:     20,
		Dexterity:    18,
		Constitution: 14,
		Charisma:     8,
		Health:       20,
		Inventory: map[string]float32{
			"sword":  15.2,
			"potion": 0.5,
		},
		Deck:            nil,
		PlayerPos_Y:     1,
		PlayerPos_X:     0,
		AttackTurnState: false,
	}

	n := Nutrals.Neutral{
		Race:         race_list[randRace],
		Resistances:  make([]string, 0, 5),
		Strength:     0,
		Dexterity:    0,
		Constitution: 0,
		Charisma:     0,
		Inventory:    make(map[string]float32, 0),
	}

	e := Enemies.Enemy{
		Race:         race_list[randRace],
		Resistances:  make([]string, 0, 5),
		Strength:     0,
		Dexterity:    0,
		Constitution: 0,
	}

	p.Deck = p.DeckSetter(p.Race)

outerLoop:
	for player.PlayerPositionChecker(player.WorldMap[p.PlayerPos_Y][p.PlayerPos_X]) {
		for p.Health > 0 {

			// Decides if they meet an enemy or a neutral or an empty spot
			landChance := rand.Intn(100)

			if landChance <= 33 {

				e.RandomAttributeSetter()

				if p.Dexterity > e.Dexterity {
					p.ChangeTurnState(p.AttackTurnState)

				} else {
					e.ChangeTurnState(e.AttackTurnState)
				}
				for p.Health > 0 || e.Health > 0 {
					reader := bufio.NewReader(os.Stdin)

					fmt.Println("You Have entered combat!!!")

					input, _ := reader.ReadString('\n')

					input = strings.TrimSpace(input)

					//available_input := make([]string, 0)

					/*for k, v := range p.Deck {
						fmt.Println(k, ":", v)

						if !strings.Contains(available_input[k]) {
							available_input = append(available_input, k)
						}
					}

					/*switch input{
						case
					}*/
				}
			}

			if landChance >= 33 || landChance <= 66 {
				fmt.Println("There is nothing here but trees...")
				continue
			}

			if landChance >= 66 {
				n.RandomAttributeSetter()
			}
			p.PlayerMove()

		}
		if player.PlayerPositionChecker(player.WorldMap[p.PlayerPos_Y][p.PlayerPos_X]) {
			fmt.Println("Congrats you reached the village!!!")
			break outerLoop
		}
	}

	fmt.Println("Thanks for playing come again!!!")
}
