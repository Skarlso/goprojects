package main

func fib(n int) int {
	if n <= 0 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

// func main() {
// 	var fibList []int
//
// 	for i := 0; i < 10; i++ {
// 		fibList = append(fibList, fib(i))
// 	}
// 	fmt.Println(fibList)
// }
