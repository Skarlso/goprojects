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

func countBestCookieRecipe(in chan []int, done chan int) {
	bestrecipe := 0
	for v := range in {
		var (
			capacity   int
			durability int
			flavor     int
			texture    int
		)

		for i := range ingredients {
			capacity += ingredients[i].capacity * v[i]
			durability += ingredients[i].durability * v[i]
			flavor += ingredients[i].flavor * v[i]
			texture += ingredients[i].texture * v[i]
		}

		// This is a more interesting approach to getting a zero value if it is a
		// negative number.
		capacity = (abs(capacity) + capacity) / 2
		durability = (abs(durability) + durability) / 2
		flavor = (abs(flavor) + flavor) / 2
		texture = (abs(texture) + texture) / 2

		recipe := capacity * durability * flavor * texture
		if recipe > bestrecipe {
			bestrecipe = recipe
		}
	}
	done <- bestrecipe
}

func generatePossibleIngredientCombinations(lenght int, out chan []int) {
	defer close(out)
	var limit = 100
	currentSeq := make([]int, lenght)
	for {
		if currentSeq[len(currentSeq)-1] == 99 {
			break
		}
		currentSeq = incrementIngredientCount(currentSeq)
		sum := sum(currentSeq...)
		if sum == limit {
			a := make([]int, len(currentSeq))
			copy(a, currentSeq)
			out <- a
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
	ch := make(chan []int, 100)
	best := make(chan int)
	defer close(best)
	go generatePossibleIngredientCombinations(len(ingredients), ch)
	go countBestCookieRecipe(ch, best)
	fmt.Println("Best combination: ", <-best)
}
