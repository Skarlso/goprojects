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

func recuriseCount(innerList []int) (sum int) {
	if len(innerList) == 0 {
		return sum
	}
	sum = innerList[0] + recuriseCount(innerList[1:])
	return
}

//func main() {
//	loopCount()
//	fmt.Println("Sum rec:", recuriseCount(list))
//}
