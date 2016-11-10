package main

import (
	"fmt"

	"github.com/Skarlso/goprojects/algorithms/memoization"
)

func main() {
	fmt.Println("vim-go")
	fmt.Println(memoization.FibFast(10))
	fmt.Println(memoization.FibFast(30))
	fmt.Println(memoization.FibFast(5))
}
