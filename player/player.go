package player

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type DamageType string

const (
	ForceDamage  DamageType = "Force"
	HolyDamage   DamageType = "Holy"
	NatureDamage DamageType = "Nature"
	PoisonDamage DamageType = "Poison"
	FireDamage   DamageType = "Fire"
	ArcaneDamage DamageType = "Arcane"
)

type Spell struct {
	DamageType DamageType
	Damage     int
}

// defines the players stats and such
type Player struct {
	Race            string
	Resistances     []string
	Journal         []string
	Strength        int
	Dexterity       int
	Constitution    int
	Charisma        int
	Health          int
	Alive           bool
	Inventory       map[string]float32
	Deck            []Spell
	PlayerPos_Y     int
	PlayerPos_X     int
	PlayerMap       [][]string
	AttackTurnState bool
}

/*This is what is unique to the player
Like move and weight check
If you want Attack etc. go to entity.go
*/

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

func (p *Player) DeckSetter(race string) {
	switch race {
	case "Human":

		//p.Deck = append(p.Deck, Spell{HolyDamage, 3})

		p.Deck = append(p.Deck,
			Spell{HolyDamage, 3},
			Spell{HolyDamage, 4},
			Spell{ForceDamage, 8},
			Spell{ForceDamage, 7},
			Spell{ForceDamage, 9})
	case "Elf":
		p.Deck = append(p.Deck,
			Spell{NatureDamage, 5},
			Spell{NatureDamage, 7},
			Spell{NatureDamage, 8},
			Spell{PoisonDamage, 4},
			Spell{PoisonDamage, 3})
	case "Orc":
		p.Deck = append(p.Deck,
			Spell{ForceDamage, 8},
			Spell{ForceDamage, 6},
			Spell{ForceDamage, 10},
			Spell{ForceDamage, 9})
	case "Gnome":
		p.Deck = append(p.Deck,
			Spell{ArcaneDamage, 8},
			Spell{ArcaneDamage, 9},
			Spell{ArcaneDamage, 10},
			Spell{NatureDamage, 7},
			Spell{PoisonDamage, 3})
	case "Trent":
		p.Deck = append(p.Deck,
			Spell{NatureDamage, 5},
			Spell{NatureDamage, 7},
			Spell{NatureDamage, 8},
			Spell{NatureDamage, 5},
			Spell{NatureDamage, 7},
			Spell{NatureDamage, 8})
	case "Dragonkin":
		p.Deck = append(p.Deck,
			Spell{FireDamage, 5},
			Spell{FireDamage, 7},
			Spell{FireDamage, 8},
			Spell{FireDamage, 5},
			Spell{FireDamage, 7},
			Spell{FireDamage, 8})
	}

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
			//p.PlayerMove()
		} else {
			p.PlayerPos_Y = p.PlayerPos_Y - 1
		}
	case "south", "s", "South", "S":
		if !PlayerPositionChecker(WorldMap[p.PlayerPos_Y+1][p.PlayerPos_X]) {
			fmt.Println("Thats water, you cant swim.")
			//p.PlayerMove()
		} else {
			p.PlayerPos_Y = p.PlayerPos_Y + 1
		}
	case "east", "e", "East", "E":
		if !PlayerPositionChecker(WorldMap[p.PlayerPos_Y][p.PlayerPos_X+1]) {
			fmt.Println("Thats water, you cant swim.")
			//p.PlayerMove()
		} else {
			p.PlayerPos_X = p.PlayerPos_X + 1
		}
	case "west", "w", "West", "W":
		if !PlayerPositionChecker(WorldMap[p.PlayerPos_Y][p.PlayerPos_X-1]) {
			fmt.Println("Thats water, you cant swim.")
			//p.PlayerMove()
		} else {
			p.PlayerPos_X = p.PlayerPos_X - 1
		}
	default:
		fmt.Println("Invalid direction")
		//p.PlayerMove()
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

func (p *Player) ChangeTurnState(pt bool) {

	if pt == true {
		p.AttackTurnState = false
	} else {
		p.AttackTurnState = true
	}
}
