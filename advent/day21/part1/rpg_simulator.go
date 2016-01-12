package main

//Boss The Baus
type Boss struct {
	hp, dmg, armor int
}

//Inventory the inventory of the player. Items are represented by numbers only
//since he can have only one of each, except for rings. If ring is -1 it's armor
//if ring is +1 it's damage (gone with separate rings for now. Later I can optimize.)
type Inventory struct {
	weapond, armor int
	defRings       []int
	dmgRings       []int
}

//Player The Playa
type Player struct {
	Inventory
	hp, dmg, armor int
}

//Weapon weapon's representation with damage and cost
type Weapon struct {
	name   string
	cost   int
	damage int
}

//Armor armor representation with armor value
type Armor struct {
	name  string
	cost  int
	armor int
}

//DefenseRing rings which improve armor
type DefenseRing struct {
	name    string
	cost    int
	defense int
}

//DamageRing rings which improve damage
type DamageRing struct {
	name   string
	cost   int
	damage int
}

//Shop a shop which has a variaty of items
type Shop struct {
	weapons      []Weapon
	armors       []Armor
	defenseRings []DefenseRing
	damageRings  []DamageRing
}

var player Player
var boss Boss
var shop Shop

func init() {
	player = Player{hp: 100, dmg: 0, armor: 0, Inventory: Inventory{}}
	boss = Boss{hp: 103, dmg: 9, armor: 2}
	shop = Shop{
		weapons: []Weapon{
			Weapon{"Dagger", 4, 8},
			Weapon{"Shortsword", 10, 5},
			Weapon{"Warhammer", 25, 6},
			Weapon{"Longsword", 40, 7},
			Weapon{"Greataxe", 74, 8},
		},
		armors: []Armor{
			Armor{"Leather", 13, 1},
			Armor{"Chainmail", 31, 2},
			Armor{"Splintmail", 53, 3},
			Armor{"Bandedmail", 75, 4},
			Armor{"Platemail", 102, 5},
		},
		defenseRings: []DefenseRing{
			DefenseRing{"Defense +1", 20, 1},
			DefenseRing{"Defense +2", 40, 2},
			DefenseRing{"Defense +3", 80, 3},
		},
		damageRings: []DamageRing{
			DamageRing{"Damage +1", 25, 1},
			DamageRing{"Damage +2", 50, 2},
			DamageRing{"Damage +3", 100, 3},
		},
	}
}

func main() {

}
