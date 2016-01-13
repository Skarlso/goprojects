package main

import (
	"fmt"
	"math"
)

//Boss The Baus
type Boss struct {
	hp, dmg, armor int
}

//Player The Playa
type Player struct {
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
	weapons      map[int]Weapon
	armors       map[int]Armor
	defenseRings map[int]DefenseRing
	damageRings  map[int]DamageRing
}

var shop Shop
var itemCombinations []int

func init() {
	shop = Shop{
		weapons: map[int]Weapon{
			0: {"Dagger", 8, 4},
			1: {"Shortsword", 10, 5},
			2: {"Warhammer", 25, 6},
			3: {"Longsword", 40, 7},
			4: {"Greataxe", 74, 8},
		},
		armors: map[int]Armor{
			1: {"Leather", 13, 1},
			2: {"Chainmail", 31, 2},
			3: {"Splintmail", 53, 3},
			4: {"Bandedmail", 75, 4},
			5: {"Platemail", 102, 5},
		},
		defenseRings: map[int]DefenseRing{
			1: {"Defense +1", 20, 1},
			2: {"Defense +2", 40, 2},
			3: {"Defense +3", 80, 3},
		},
		damageRings: map[int]DamageRing{
			1: {"Damage +1", 25, 1},
			2: {"Damage +2", 50, 2},
			3: {"Damage +3", 100, 3},
		},
	}
}

func main() {
	smallestCost := math.MaxInt64
	var (
		weapondmg int
		armor     int
		defring   int
		dmgring   int
		cWeapond  int
		cArmor    int
		cDefRing  int
		cDmgRing  int
	)

	for _, v := range shop.weapons {
		fmt.Println("Starting game with weapond:", v.name)
		weapondmg = v.damage
		cWeapond = v.cost
		for a := 0; a <= len(shop.armors); a++ {
			if _, ok := shop.armors[a]; ok {
				fmt.Println("Starting game with armor:", shop.armors[a].name)
				armor = shop.armors[a].armor
				cArmor = shop.armors[a].cost
			} else {
				armor = 0
				cArmor = 0
			}
			for defr := 0; defr <= len(shop.defenseRings); defr++ {
				if _, ok := shop.defenseRings[defr]; ok {
					fmt.Println("Starting game with defense ring:", shop.defenseRings[defr].name)
					defring = shop.defenseRings[defr].defense
					cDefRing = shop.defenseRings[defr].cost
				} else {
					defring = 0
					cDefRing = 0
				}
				for dmgr := 0; dmgr <= len(shop.damageRings); dmgr++ {
					if _, ok := shop.damageRings[dmgr]; ok {
						fmt.Println("Starting game with damage ring:", shop.damageRings[dmgr].name)
						dmgring = shop.damageRings[dmgr].damage
						cDmgRing = shop.damageRings[dmgr].cost
					} else {
						dmgring = 0
						cDmgRing = 0
					}
					moneySpent := cWeapond + cArmor + cDefRing + cDmgRing
					playersTurn := true

					player := &Player{hp: 100, dmg: weapondmg + dmgring, armor: armor + defring}
					boss := &Boss{hp: 103, dmg: 9, armor: 2}
					fmt.Println("Player:", *player)
					fmt.Println("Boss:", *boss)
					for {
						// fmt.Printf("Player's hp:%d | Boss hp:%d \n", player.hp, boss.hp)
						switch playersTurn {
						case true:
							player.attack(boss)
							playersTurn = false
						case false:
							boss.attack(player)
							playersTurn = true
						}

						if player.hp <= 0 || boss.hp <= 0 {
							break
						}
					}

					if player.hp > 0 {
						if moneySpent < smallestCost {
							smallestCost = moneySpent
						}
					}

				}
			}
		}

		fmt.Println("Smallest cost spent on a win:", smallestCost)
	}
}
func (p *Player) attack(b *Boss) {
	dmg := p.dmg - b.armor
	if dmg <= 0 {
		dmg = 1
	}
	b.hp -= dmg
}

func (b *Boss) attack(p *Player) {
	dmg := b.dmg - p.armor
	if dmg <= 0 {
		dmg = 1
	}
	p.hp -= dmg
}
