package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/KhrisKringle/Stratagus/NPC/Enemies"
	"github.com/KhrisKringle/Stratagus/NPC/Nutrals"
	"github.com/KhrisKringle/Stratagus/combat"
	"github.com/KhrisKringle/Stratagus/player"
	"golang.org/x/exp/rand"
)

func main() {

	//var race string
	//race = "Elf"
	var race_list = [6]string{"Elf", "Human", "Orc", "Gnome", "Trent", "Dragonkin"}
	randRace := rand.Intn(5)
	p := player.Player{
		Race:         "",
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
		Deck:            make([]player.Spell, 0),
		SpellMod:        0,
		PlayerPos_Y:     1,
		PlayerPos_X:     0,
		AttackTurnState: false,
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Pick a race (Elf, Human, Orc, Gnome, Trent, Dragonkin): ")

		race, _ := reader.ReadString('\n')
		race = strings.TrimSpace(race)
		fmt.Println(race)

		switch race {
		case "Human", "human":
			p.Race = "Human"
		case "Elf", "elf":
			p.Race = "Elf"
		case "Orc", "orc":
			p.Race = "Orc"
		case "Gnome", "gnome":
			p.Race = "Gnome"
		case "Trent", "trent":
			p.Race = "Trent"
		case "Dragonkin", "dragonkin":
			p.Race = "Dragonkin"
		default:
			fmt.Println("invalid Race")
			continue
		}
		break

	}

	p.DeckSetter(p.Race)

	for {
		// Decides if they meet an enemy or a neutral or an empty spot
		landChance := rand.Intn(100)

		fmt.Println("Y:", p.PlayerPos_Y)
		fmt.Println("X:", p.PlayerPos_X)
		if p.PlayerPos_Y == 3 && p.PlayerPos_X == 3 {
			fmt.Println("Congrats you reached the village!!!")
			break
		}

		//landChance := 66
		if landChance < 33 {
			e := Enemies.Enemy{
				Race:         race_list[randRace],
				Resistances:  make([]string, 0, 5),
				Health:       20,
				Strength:     0,
				Dexterity:    0,
				Constitution: 0,
			}
			e.RandomAttributeSetter()

			fmt.Println(landChance)
			if p.Dexterity > e.Dexterity {
				p.ChangeTurnState(p.AttackTurnState)
			} else if p.Dexterity < e.Dexterity {
				e.ChangeTurnState(e.AttackTurnState)
			} else {
				p.ChangeTurnState(p.AttackTurnState)
			}

			fmt.Println("You Have entered combat!!!")

			for {
				fmt.Printf("Player Dex: %v\n", p.Dexterity)
				fmt.Printf("Enemy Dex: %v\n", e.Dexterity)
				fmt.Printf("Player attackstate: %v\n", p.AttackTurnState)
				fmt.Printf("Enemy attackstate: %v\n", e.AttackTurnState)
				if p.AttackTurnState {
					if p.Health <= 0 {
						fmt.Println("GAME OVER!!!")
						break
					}
					// Prints out the Deck
					for {
						fmt.Println("Your Hand:")
						available_input := make([]player.Spell, 0, len(p.Deck))
						player.PrintSpells(p.Deck, available_input)
						fmt.Print("Enter what spell you want to use: ")
						reader := bufio.NewReader(os.Stdin)
						input, _ := reader.ReadString('\n')

						inputv2 := strings.TrimSpace(input)

						foundSpell := false

						// Check if the input is in the deck
						for i, x := range p.Deck {
							if inputv2 == fmt.Sprintf("%s %d", x.DamageType, x.Damage) {
								foundSpell = true
								p.SpellMod = x.Damage
								p.RemoveSpellAtIndex(i)
								break
							}
						}

						if !foundSpell {
							fmt.Printf("You do not have the spell type %s\n", inputv2)
						} else {
							e.TakeDamage(combat.Attack(p.Strength, e.Strength, p, e) + p.SpellMod)
							fmt.Printf("You dealt %v damage\n", combat.Attack(p.Strength, e.Strength, p, e)+p.SpellMod)
							p.ChangeTurnState(p.AttackTurnState)
							e.ChangeTurnState(e.AttackTurnState)
							fmt.Println("Enemy Health: ", e.Health)
							break
						}
					}
				} else {
					if e.Health <= 0 {
						fmt.Println("You defeated the Beast!!!")
						break
					}
					p.TakeDamage(combat.Attack(e.Strength, p.Strength, p, e))
					fmt.Printf("You took %v damage\n", combat.Attack(e.Strength, p.Strength, p, e)+p.SpellMod)
					fmt.Println("Player Health: ", p.Health)
					e.ChangeTurnState(e.AttackTurnState)
					p.ChangeTurnState(p.AttackTurnState)
				}
			}
			p.PlayerMove()
		}

		if landChance >= 33 || landChance <= 66 {
			fmt.Println(landChance)
			fmt.Println("There is nothing here but trees...")
			p.PlayerMove()
			continue
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
			continue
		}
	}

	fmt.Println("Thanks for playing come again!!!")
}
