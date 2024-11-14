package Player


type Player struct {
	Strength int
	Dexterity int
	Constitution int
	Intelligence int
	Wisdom int
	Charisma int
	Inventory map[string]float32
}

func (p Player) Modifier(attribute string) int {
	switch attribute{
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