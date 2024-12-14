package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/KhrisKringle/Stratagus/NPC/Enemies"
	"github.com/KhrisKringle/Stratagus/NPC/Nutrals"
	"github.com/KhrisKringle/Stratagus/player"
	"golang.org/x/exp/rand"
)

func main() {

	var race string
	var race_list = [6]string{"Elf", "Human", "Orc", "Gnome", "Trent", "Dragonkin"}
	randRace := rand.Intn(5)
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Pick a race (Elf, Human, Orc, Gnome, Trent, Dragonkin): ")

		race, _ := reader.ReadString('\n')
		race = strings.TrimSpace(race)

		switch race {
		case "Human", "human":
			race = "Human"
		case "Elf", "elf":
			race = "Elf"
		case "Orc", "orc":
			race = "Orc"
		case "Gnome", "gnome":
			race = "Gnome"
		case "Trent", "trent":
			race = "Trent"
		case "Dragonkin", "dragonkin":
			race = "Dragonkin"
		default:
			fmt.Println("invalid Race")
			continue
		}
		break

	}

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

	p.DeckSetter(p.Race)

	for {
		// Decides if they meet an enemy or a neutral or an empty spot
		//landChance := rand.Intn(100)

		fmt.Println("Y:", p.PlayerPos_Y)
		fmt.Println("X:", p.PlayerPos_X)
		if p.PlayerPos_Y == 3 && p.PlayerPos_X == 3 {
			fmt.Println("Congrats you reached the village!!!")
			break
		}

		landChance := 66
		if landChance < 33 {
			e := Enemies.Enemy{
				Race:         race_list[randRace],
				Resistances:  make([]string, 0, 5),
				Strength:     0,
				Dexterity:    0,
				Constitution: 0,
			}
			fmt.Println(landChance)
			e.RandomAttributeSetter()

			if p.Dexterity > e.Dexterity {
				p.ChangeTurnState(p.AttackTurnState)

			} else if p.Dexterity < e.Dexterity {
				e.ChangeTurnState(e.AttackTurnState)
			} else {
				p.ChangeTurnState(p.AttackTurnState)
			}
			fmt.Println("You Have entered combat!!!")

			for p.Health > 0 || e.Health > 0 {
				if p.AttackTurnState {
					reader := bufio.NewReader(os.Stdin)

					// Prints out the Deck
					fmt.Println("Your Hand:")
					available_input := make([]player.Spell, 0)
					for _, spell := range p.Deck {
						available_input = append(available_input, spell)
						fmt.Printf("+----------+\n")
						fmt.Printf("| [%s] %d  |\n", spell.DamageType, spell.Damage)
						fmt.Printf("+----------+\n")
					}

					fmt.Print("Enter what spell you want ot use: ")
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
				} else if e.AttackTurnState {
					p.TakeDamage(e.DoDamage())
				}
			}
			p.PlayerMove()
		}

		if landChance >= 33 || landChance <= 66 {
			fmt.Println(landChance)
			fmt.Println("There is nothing here but trees...")
			p.PlayerMove()
		}

		if landChance > 66 {
			n := Nutrals.Neutral{
				Race:         race_list[randRace],
				Resistances:  make([]string, 0, 5),
				Strength:     0,
				Dexterity:    0,
				Constitution: 0,
				Charisma:     0,
				Inventory:    make(map[string]float32, 0),
			}
			fmt.Println(landChance)
			n.RandomAttributeSetter()
			fmt.Println("There is nothing here but trees...")
			fmt.Println("This is where you would meet someone.")
			p.PlayerMove()
		}
	}

	fmt.Println("Thanks for playing come again!!!")
}
