package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//Reindeer this is a reindeer, has a name, a speed and a time it requires to rest
type Reindeer struct {
	name          string
	speed         int
	limit         int
	sleepDuration int
	distanceMoved int
	timeMoved     int
	rested        int
	award         int
}

var reindeers = make([]*Reindeer, 0)

func (r Reindeer) String() string {
	return fmt.Sprintf("Name:%s; Speed: %d; Limit: %d; SleepDuration:%d; distanceMoved:%d; timeMoved:%d; rested:%d", r.name, r.speed, r.limit, r.sleepDuration, r.distanceMoved, r.timeMoved, r.rested)
}

func init() {
	file, _ := os.Open("test_input.txt")
	defer file.Close()
	in := bufio.NewReader(file)
	for {
		var (
			name                        string
			speed, limit, sleepDuration int
		)
		n, err := fmt.Fscanf(in, "%s can fly %d km/s for %d seconds, but then must rest for %d seconds.", &name, &speed, &limit, &sleepDuration)
		if err != nil {
			if err == io.EOF {
				break
			}
		}
		//Skip if no lines were parsed
		if n == 0 {
			continue
		}
		r := Reindeer{name, speed, limit, sleepDuration, 0, 0, 0, 0}
		reindeers = append(reindeers, &r)
	}
}

func main() {
	startRace()
	mostMoved := reindeers[0]
	for _, r := range reindeers {
		if r.distanceMoved > mostMoved.distanceMoved {
			mostMoved = r
		}
	}

	fmt.Println("Most moved reindeer:", mostMoved)
}

func startRace() {
	for i := 0; i <= 1000; i++ {
		award := gatherLeadingReindeers()
		for _, v := range award {
			v.distanceMoved++
		}
		//Check if a reindeer can move
		//If yes, move the reindeer
		for _, rein := range reindeers {
			if rein.canMove() {
				rein.move()
			} else {
				rein.rest()
			}
		}
	}
	//Reindeer with greatest distance travelled, wins
}

func gatherLeadingReindeers() []*Reindeer {
	var leaders []*Reindeer
	leaders = append(leaders, reindeers[0])
	for _, v := range reindeers {
		if v.distanceMoved > leaders[0].distanceMoved {
			leaders = nil
			leaders = make([]*Reindeer, 0)
			leaders = append(leaders, v)
		} else if v.distanceMoved == leaders[0].distanceMoved {
			leaders = append(leaders, v)
		}
	}
	return leaders
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
