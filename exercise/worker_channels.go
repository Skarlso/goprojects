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
		// fmt.Println("Checking p:", p)
		if p > LIMIT {
			output <- p
		}
	}
}

func passwordIncrement(out chan<- int) {
	p := 0
	for {
		p++
		// fmt.Println("Generated Password:", p)
		out <- p
	}
}
