package main

import "fmt"

//Character represents the Boss and a Player
type Character struct {
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
			0: {"Nothing", 0, 0},
			1: {"Leather", 13, 1},
			2: {"Chainmail", 31, 2},
			3: {"Splintmail", 53, 3},
			4: {"Bandedmail", 75, 4},
			5: {"Platemail", 102, 5},
		},
		defenseRings: map[int]DefenseRing{
			0: {"Nothing", 0, 0},
			1: {"Defense +1", 20, 1},
			2: {"Defense +2", 40, 2},
			3: {"Defense +3", 80, 3},
		},
		damageRings: map[int]DamageRing{
			0: {"Nothing", 0, 0},
			1: {"Damage +1", 25, 1},
			2: {"Damage +2", 50, 2},
			3: {"Damage +3", 100, 3},
		},
	}
}

func main() {
	leastSpent := 0
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
			fmt.Println("Starting game with armor:", shop.armors[a].name)
			armor = shop.armors[a].armor
			cArmor = shop.armors[a].cost
			for defr := 0; defr <= len(shop.defenseRings); defr++ {
				fmt.Println("Starting game with defense ring:", shop.defenseRings[defr].name)
				defring = shop.defenseRings[defr].defense
				cDefRing = shop.defenseRings[defr].cost
				for dmgr := 0; dmgr <= len(shop.damageRings); dmgr++ {

					fmt.Println("Starting game with damage ring:", shop.damageRings[dmgr].name)
					dmgring = shop.damageRings[dmgr].damage
					cDmgRing = shop.damageRings[dmgr].cost

					moneySpent := cWeapond + cArmor + cDefRing + cDmgRing
					playersTurn := true

					player := &Character{hp: 100, dmg: weapondmg + dmgring, armor: armor + defring}
					boss := &Character{hp: 103, dmg: 9, armor: 2}
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

					if player.hp <= 0 {
						if moneySpent > leastSpent {
							leastSpent = moneySpent
						}
					}

				}
			}
		}

		fmt.Println("Smallest cost spent on a win:", leastSpent)
	}
}

func (c1 *Character) attack(c2 *Character) {
	dmg := c1.dmg - c2.armor
	if dmg <= 0 {
		dmg = 1
	}
	c2.hp -= dmg
}
