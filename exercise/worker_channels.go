package main

import "fmt"

//LIMIT limit
const LIMIT = 10000000

func main() {
	generatedPassword := make(chan int, 100)
	correctPassword := make(chan int)
	defer close(generatedPassword)
	defer close(correctPassword)
	go passwordIncrement(generatedPassword)
	go checkPassword(generatedPassword, correctPassword)
	pass := <-correctPassword
	fmt.Println(pass)
}

func checkPassword(input <-chan int, output chan<- int) {
	for {
		p := <-input
		//Introduce lengthy operation here
		// time.Sleep(time.Second)
		fmt.Println("Checking p:", p)
		correct := check(p)
		if correct {
			output <- p
		}
	}
}

func increment(i int) int {
	return i + 1
}

func check(i int) bool {

	for i := 0; i < 10; i++ {
		p := i * 2
		if p&1 == 1 {

		}
	}

	if i&1 == 1 && i > 50000 {
		return true
	}

	return false
}

func passwordIncrement(out chan<- int) {
	p := 0
	for {
		p = increment(p)
		fmt.Println("Generated Password:", p)
		out <- p
	}
}
