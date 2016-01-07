package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//Ingredient cookie ingredients
type Ingredient struct {
	name       string
	capacity   int
	durability int
	flavor     int
	texture    int
	calories   int
}

var ingredients []*Ingredient

func init() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	in := bufio.NewReader(file)
	for {
		var (
			name                                            string
			capacity, durability, flavor, texture, calories int
		)
		n, err := fmt.Fscanf(in, "%s: capacity %d, durability %d, flavor %d, texture %d, calories %d", &name, &capacity, &durability, &flavor, &texture, &calories)
		if err != nil {
			if err == io.EOF {
				break
			}
		}
		//Skip if no lines were parsed
		if n == 0 {
			continue
		}
		i := Ingredient{name, capacity, durability, flavor, texture, calories}
		ingredients = append(ingredients, &i)
	}
}

func main() {
	fmt.Println("Making cookies...")
}
