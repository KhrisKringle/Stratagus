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
		Deck:        nil,
		PlayerPos_Y: 1,
		PlayerPos_X: 0,
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


    playerTurn := false
    enemyTurn := false

    outerLoop:
	for player.PlayerPositionChecker(player.WorldMap[p.PlayerPos_Y][p.PlayerPos_X]) {
		for  p.Health > 0 {
			landChance := rand.Intn(100)

			if landChance <= 33 {
                
				e.RandomAttributeSetter()

				if p.Dexterity > e.Dexterity {
					playerTurn = true

				} else { 
                    enemyTurn = true
                }
                for p.Health > 0 || e.Health > 0{
				    
                }
			}

			if landChance >= 33 || landChance <= 66 {
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
