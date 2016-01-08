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

//ingredients all the ingredients we need to make a delicious cookie
var ingredients = []Ingredient{
	{"Sprinkles", 2, 0, -2, 0, 3},
	{"Butterscotch", 0, 5, -3, 0, 3},
	{"Chocolate", 0, 0, 5, -1, 8},
	{"Candy", 0, -1, 0, 5, 8},
}

//getScore calculate the score for a given recipe combination
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
	// negative number than having 4 lines of ifs. Though in line size, it's
	// the same.
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

//generatePossibleIngredientCombinations Generates combinations of possible ingredient
//counts by simple brute force counting.
func generatePossibleIngredientCombinations(lenght int) {
	var limit = 100
	var score int
	currentSeq := make([]int, lenght)
	for {
		//If the last variable is 99, we have exhausted our combination possibilities
		if currentSeq[len(currentSeq)-1] == 99 {
			fmt.Println("Best score:", score)
			break
		}
		//Increase the next number with a recursive function
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

//incrementIngredientCount one by one increments number combinations
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

//sum sums an arbitary number of ints
func sum(nums ...int) int {
	sum := 0
	for _, i := range nums {
		sum += i
	}
	return sum
}

//abs absolute value
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
