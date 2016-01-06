package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//Reindeer this is a reindeer, has a name, a speed and a time it requires to rest
type Reindeer struct {
	name  string
	speed int
	limit int
	sleep int
}

var reindeers []Reindeer = make([]Reindeer, 0)

func init() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	in := bufio.NewReader(file)
	for {
		var (
			name                string
			speed, limit, sleep int
		)
		n, err := fmt.Fscanf(in, "%s can fly %d km/s for %d seconds, but then must rest for %d seconds.", &name, &speed, &limit, &sleep)
		if err != nil {
			if err == io.EOF {
				break
			}
		}
		//Skip if no lines were parsed
		if n == 0 {
			continue
		}
		r := Reindeer{name, speed, limit, sleep}
		reindeers = append(reindeers, r)
	}
}

func main() {
	fmt.Println("Reindeers:", reindeers)
}

func startRace() {
}
