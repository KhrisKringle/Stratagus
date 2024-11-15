package Player

type Player struct {
	Race         string
	Resistances  []string
	Journal      []string
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Inventory    map[string]float32
	Deck         []map[string]int
}

// Allways got the modifier on deck.
func (p Player) Modifier(attribute string) int {
	switch attribute {
	case "Strength":
		return (p.Strength - 10) / 2
	case "Dexterity":
		return (p.Dexterity - 10) / 2
	case "Constitution":
		return (p.Constitution - 10) / 2
	case "Intelligence":
		return (p.Intelligence - 10) / 2
	case "Wisdom":
		return (p.Wisdom - 10) / 2
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
	deck := make([]map[string]int, 20)
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
