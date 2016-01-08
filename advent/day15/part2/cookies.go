package main

import "fmt"

//Ingredient cookie ingredients
type Ingredient struct {
	name       string
	capacity   int
	durability int
	flavor     int
	texture    int
	calories   int
}

var ingredients = []Ingredient{
	{"Sprinkles", 2, 0, -2, 0, 3},
	{"Butterscotch", 0, 5, -3, 0, 3},
	{"Chocolate", 0, 0, 5, -1, 8},
	{"Candy", 0, -1, 0, 5, 8},
}

func getScore(v []int) (score int) {
	var (
		capacity   int
		durability int
		flavor     int
		texture    int
		calories   int
	)

	for i := range ingredients {
		capacity += ingredients[i].capacity * v[i]
		durability += ingredients[i].durability * v[i]
		flavor += ingredients[i].flavor * v[i]
		texture += ingredients[i].texture * v[i]
		calories += ingredients[i].calories * v[i]
	}

	// This is a more interesting approach to getting a zero value if it is a
	// negative number.
	capacity = (abs(capacity) + capacity) / 2
	durability = (abs(durability) + durability) / 2
	flavor = (abs(flavor) + flavor) / 2
	texture = (abs(texture) + texture) / 2

	recipe := capacity * durability * flavor * texture
	if calories == 500 {
		score = recipe
	}
	return
}

func generatePossibleIngredientCombinations(lenght int) {
	var limit = 100
	var score int
	currentSeq := make([]int, lenght)
	for {
		if currentSeq[len(currentSeq)-1] == 99 {
			fmt.Println("Best score:", score)
			break
		}
		currentSeq = incrementIngredientCount(currentSeq)
		sum := sum(currentSeq...)
		if sum == limit {
			currentScore := getScore(currentSeq)
			if currentScore > score {
				score = currentScore
			}
		}
	}
}

func incrementIngredientCount(arr []int) []int {
	if len(arr) == 1 {
		arr[0] = ((arr[0] + 1) % 99) + 1
		return arr
	}
	if arr[0] == 99 {
		incrementIngredientCount(arr[1:])
	}
	arr[0] = ((arr[0] + 1) % 99) + 1
	return arr
}

func sum(nums ...int) int {
	sum := 0
	for _, i := range nums {
		sum += i
	}
	return sum
}

func abs(x int) int {
	switch {
	case x < 0:
		return -x
	case x == 0:
		return 0 // return correctly abs(-0)
	}
	return x
}

func main() {
	generatePossibleIngredientCombinations(len(ingredients))
}
