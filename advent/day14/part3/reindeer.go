package main

import (
	"fmt"
	"sync"
)

const FINISH = 2503

//Reindeer this is a reindeer, has a name, a speed and a time it requires to rest
type Reindeer struct {
	name          string
	speed         int
	limit         int
	sleepDuration int
	distanceMoved int
	timeMoved     int
	rested        int
}

var reindeers = []*Reindeer{
	{"Rudolph", 3, 15, 28, 0, 0, 0},
	{"Donner", 19, 9, 164, 0, 0, 0},
	{"Blitzen", 19, 9, 158, 0, 0, 0},
	{"Comet", 13, 7, 82, 0, 0, 0},
	{"Vixen", 19, 7, 124, 0, 0, 0},
	{"Cupid", 25, 6, 145, 0, 0, 0},
	{"Dasher", 14, 3, 38, 0, 0, 0},
	{"Dancer", 3, 16, 37, 0, 0, 0},
	{"Prancer", 25, 6, 143, 0, 0, 0},
}

func (r Reindeer) String() string {
	return fmt.Sprintf("Name:%s; Speed: %d; Limit: %d; SleepDuration:%d; distanceMoved:%d; timeMoved:%d; rested:%d", r.name, r.speed, r.limit, r.sleepDuration, r.distanceMoved, r.timeMoved, r.rested)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(len(reindeers))

	for _, r := range reindeers {
		go func(rein *Reindeer) {
			defer wg.Done()
			for i := 0; i < FINISH; i++ {
				if rein.canMove() {
					rein.move()
				} else {
					rein.rest()
				}
			}
		}(r)
	}

	wg.Wait()

	mostMoved := reindeers[0]
	for _, r := range reindeers {
		if r.distanceMoved > mostMoved.distanceMoved {
			mostMoved = r
		}
	}

	fmt.Println("Most moved reindeer:", mostMoved)
}

func (r Reindeer) canMove() bool {
	if r.timeMoved < r.limit {
		return true
	}
	return false
}

func (r *Reindeer) move() {
	r.distanceMoved += r.speed
	r.timeMoved++
	r.rested = 0
}

func (r *Reindeer) rest() {
	r.rested++
	if r.rested == r.sleepDuration {
		r.timeMoved = 0
	}
}
