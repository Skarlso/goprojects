package main

//Boss represents the Boss
type Boss struct {
	hp, dmg, armor int
	effects        []Effect
}

//Effect which is either beneficial or harmfull
type Effect struct {
	duration  int
	remaining int
	active    bool
}

//Wizard wizard
type Wizard struct {
	hp, dmg, mana, armor int
	spells               []Spell
	effects              []Effect
}

//Spell spells
//effectDuration if 0 the spell is instantenious
type Spell struct {
	name           string
	cost           int
	dmg            int
	heal           int
	mana           int
	armor          int
	effectDuration int
}

var spell Spell
var itemCombinations []int

func main() {
	// turnCount := 0
	wizard := &Wizard{hp: 50, mana: 500, armor: 0}
	spells := []Spell{
		{name: "Magic Missile", cost: 53, dmg: 4, heal: 0, mana: 0, armor: 0, effectDuration: 0},
		{name: "Drain", cost: 73, dmg: 2, heal: 2, mana: 0, armor: 0, effectDuration: 0},
		{name: "Shield", cost: 113, dmg: 0, heal: 0, mana: 0, armor: 7, effectDuration: 6},
		{name: "Poison", cost: 173, dmg: 3, heal: 0, mana: 0, armor: 0, effectDuration: 6},
		{name: "Recharge", cost: 229, dmg: 0, heal: 0, mana: 101, armor: 0, effectDuration: 5},
	}
	wizard.spells = spells
	// boss := &Boss{hp: 55, dmg: 8}

	//If the wizard runs out of mana and no recharge is active -> Lose
	//If the wizard's hp is 0 or below -> lose
	//If the boss' hp is 0 or below -> win

	//Apply effects with a function
	//Apply to wizard and boss alike
	//Write the combinations of spells to use and effects to apply

}

func (b *Boss) attack(w *Wizard) {
	dmg := b.dmg - w.armor
	if dmg <= 0 {
		dmg = 1
	}
	w.hp -= dmg
}
