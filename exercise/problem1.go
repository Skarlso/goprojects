package main

import "fmt"

var list = []int{1, 2, 3, 4, 5}

func loopCount() {
	sum := 0
	for _, v := range list {
		sum += v
	}

	fmt.Println("Sum is:", sum)
}

func recuriseCount(index int) (sum int) {
	if index == len(list) {
		return sum
	}
	sum = list[index]
	sum += recuriseCount(index + 1)
	return
}

// func main() {
// 	loopCount()
// 	fmt.Println("Sum rec:", recuriseCount(0))
// }
