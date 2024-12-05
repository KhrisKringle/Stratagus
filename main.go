package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
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

	p.DeckSetter(p.Race)

outerLoop:
	for player.PlayerPositionChecker(player.WorldMap[p.PlayerPos_Y][p.PlayerPos_X]) {
		for p.Health > 0 {

			// Decides if they meet an enemy or a neutral or an empty spot
			landChance := rand.Intn(100)

			if landChance <= 33 {
				fmt.Println(landChance)
				e.RandomAttributeSetter()

				if p.Dexterity > e.Dexterity {
					p.ChangeTurnState(p.AttackTurnState)

				} else {
					e.ChangeTurnState(e.AttackTurnState)
				}
				fmt.Println("You Have entered combat!!!")
				for p.Health > 0 || e.Health > 0 {
					reader := bufio.NewReader(os.Stdin)

					fmt.Println("Your Hand:")
					available_input := make([]player.Spell, 0)
					for _, spell := range p.Deck {
						available_input = append(available_input, spell)
						fmt.Printf("+----------+\n")
						fmt.Printf("| [%s] %d  |\n", spell.DamageType, spell.Damage)
						fmt.Printf("+----------+\n")
					}

					input, _ := reader.ReadString('\n')

					inputv2 := strings.TrimSpace(input)

					parts := strings.Split(inputv2, " ")
					damageType := player.DamageType(parts[0])
					damage, _ := strconv.Atoi(parts[1])

					spell := player.Spell{
						DamageType: damageType,
						Damage:     damage,
					}

					for _, x := range available_input {
						if spell != x {
							fmt.Printf("You do not have the spell type %s\n", inputv2)
							break
						}
					}
					e.Health = 0
				}
				for {
					reader := bufio.NewReader(os.Stdin)

					fmt.Print("Enter a direction (north, south, east, west): ")

					// Read the user's input
					input, _ := reader.ReadString('\n')

					// Remove the newline character
					inputv2 := strings.TrimSpace(input)

					// Check the input and take appropriate action
					p.PlayerMove(inputv2)
				}
			}

			if landChance >= 33 || landChance <= 66 {
				fmt.Println(landChance)
				fmt.Println("There is nothing here but trees...")
				reader := bufio.NewReader(os.Stdin)

				fmt.Print("Enter a direction (north, south, east, west): ")

				// Read the user's input
				input, _ := reader.ReadString('\n')

				// Remove the newline character
				inputv2 := strings.TrimSpace(input)

				// Check the input and take appropriate action
				p.PlayerMove(inputv2)
			}

			if landChance >= 66 {
				fmt.Println(landChance)
				n.RandomAttributeSetter()
				fmt.Println("There is nothing here but trees...")
				for {
					reader := bufio.NewReader(os.Stdin)

					fmt.Print("Enter a direction (north, south, east, west): ")

					// Read the user's input
					input, _ := reader.ReadString('\n')

					// Remove the newline character
					inputv2 := strings.TrimSpace(input)

					// Check the input and take appropriate action
					p.PlayerMove(inputv2)
				}
			}

		}
		if player.PlayerPositionChecker(player.WorldMap[p.PlayerPos_Y][p.PlayerPos_X]) {
			if p.PlayerPos_Y == 3 && p.PlayerPos_X == 3 {
				fmt.Println("Congrats you reached the village!!!")

				break outerLoop
			}
		}
	}

	fmt.Println("Thanks for playing come again!!!")
}
