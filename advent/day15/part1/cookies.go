package main

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
var validIngredientCountCombinations = make([][]int, 0)

func generatePossibleIngredientCombinations() {
	var limit = 100
	currentSeq := []int{1, 1, 1, 1}
	for {
		if currentSeq[len(currentSeq)-1] == 99 {
			break
		}
		currentSeq = incrementIngredientCount(currentSeq)
		sum := sum(currentSeq...)
		if sum == limit {
			a := make([]int, len(currentSeq))
			copy(a, currentSeq)
			validIngredientCountCombinations = append(validIngredientCountCombinations, a)
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

func main() {
	generatePossibleIngredientCombinations()
}
