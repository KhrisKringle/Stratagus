/*This is what is unique to the player
Like move and weight check
If you want Attack etc. go to entity.go
*/

package player

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
// defines the players stats and such
type Player struct {
	Race         string
	Resistances  []string
	Journal      []string
	Strength     int
	Dexterity    int
	Constitution int
	Charisma     int
	Health       int
	Alive        bool
	Inventory    map[string]float32
	Deck         []map[string]int
	PlayerPos_Y  int
	PlayerPos_X  int
	PlayerMap    [][]string
}

// 2D matrix of the world map
var WorldMap = [][]string{
	{"water", "land", "land", "land", "water"},
	{"land", "land", "land", "land", "water"},
	{"water", "land", "water", "land", "land"},
	{"water", "land", "land", "village", "land"},
	{"water", "water", "water", "water", "water"},
}

// Allways got the modifier on deck
func (p Player) Modifier(attribute string) int {
	switch attribute {
	case "Strength":
		return (p.Strength - 10) / 2
	case "Dexterity":
		return (p.Dexterity - 10) / 2
	case "Constitution":
		return (p.Constitution - 10) / 2
	case "Charisma":
		return (p.Charisma - 10) / 2
	default:
		return 0
	}
}

// Checks to see if the player is overweight
func (p Player) WeightChecker(inv map[string]float32) bool {
	var total float32 = 0
	for _, v := range p.Inventory {
		total += v
	}
	if total > 150 {
		return false
	} else {
		return true
	}
}

func (p *Player) DeckSetter(race string) []map[string]int {
	deck := make([]map[string]int, 0, 20)
	switch race {
	case "Human":
		deck = append(deck,
			map[string]int{"Holy": 3},
			map[string]int{"Holy": 4},
			map[string]int{"Force": 8},
			map[string]int{"Force": 7},
			map[string]int{"Force": 9})
	case "Elf":
		deck = append(deck,
			map[string]int{"Nature": 5},
			map[string]int{"Nature": 7},
			map[string]int{"Nature": 8},
			map[string]int{"Poison": 4},
			map[string]int{"Poison": 3})
	case "Orc":
		deck = append(deck,
			map[string]int{"Force": 8},
			map[string]int{"Force": 6},
			map[string]int{"Force": 10},
			map[string]int{"Force": 9})
	case "Gnome":
		deck = append(deck,
			map[string]int{"Arcane": 8},
			map[string]int{"Arcane": 9},
			map[string]int{"Arcane": 10},
			map[string]int{"Nature": 7},
			map[string]int{"Poison": 3})
	case "Trent":
		deck = append(deck,
			map[string]int{"Nature": 5},
			map[string]int{"Nature": 7},
			map[string]int{"Nature": 8},
			map[string]int{"Nature": 5},
			map[string]int{"Nature": 7},
			map[string]int{"Nature": 8})
	case "Dragonkin":
		deck = append(deck,
			map[string]int{"Fire": 5},
			map[string]int{"Fire": 7},
			map[string]int{"Fire": 8},
			map[string]int{"Fire": 5},
			map[string]int{"Fire": 7},
			map[string]int{"Fire": 8})
	}
	return deck
}

// Moves the player
func (p *Player) PlayerMove() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter a direction (north, south, east, west): ")

	// Read the user's input
	input, _ := reader.ReadString('\n')

	// Remove the newline character
	input = strings.TrimSpace(input)

	// Check the input and take appropriate action
	switch input {
	case "north", "n", "North", "N":
		if !PlayerPositionChecker(WorldMap[p.PlayerPos_Y-1][p.PlayerPos_X]) {
			fmt.Println("Thats water, you cant swim.")
			p.PlayerMove()
		} else {
			p.PlayerPos_Y = p.PlayerPos_Y - 1
		}
	case "south", "s", "South", "S":
		if !PlayerPositionChecker(WorldMap[p.PlayerPos_Y+1][p.PlayerPos_X]) {
			fmt.Println("Thats water, you cant swim.")
			p.PlayerMove()
		} else {
			p.PlayerPos_Y = p.PlayerPos_Y + 1
		}
	case "east", "e", "East", "E":
		if !PlayerPositionChecker(WorldMap[p.PlayerPos_Y][p.PlayerPos_X+1]) {
			fmt.Println("Thats water, you cant swim.")
			p.PlayerMove()
		} else {
			p.PlayerPos_X = p.PlayerPos_X + 1
		}
	case "west", "w", "West", "W":
		if !PlayerPositionChecker(WorldMap[p.PlayerPos_Y][p.PlayerPos_X-1]) {
			fmt.Println("Thats water, you cant swim.")
			p.PlayerMove()
		} else {
			p.PlayerPos_X = p.PlayerPos_X - 1
		}
	default:
		fmt.Println("Invalid direction")
		p.PlayerMove()
	}
}

// Checks the players position
func PlayerPositionChecker(s string) bool {
	if s != ("water") {
		return true
	} else {
		return false
	}
}

func (p *Player) TakeDamage(dmg int) {
	if p.Health != 0 {
		p.Health -= dmg
	} else {
		p.Alive = false
	}
}
