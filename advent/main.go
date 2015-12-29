package main

import "github.com/skarlso/goprojects/advent/solutions"
import "runtime"

//
func main() {
	// solutions.MoveSanta()
	// solutions.WrappingPaper()
	// solutions.DeliverPresents()
	// solutions.Mine()
	// solutions.NiceOrNaughty()
	// solutions.TurnOnTheLightsV2()
	// solutions.BobbyTable()
	// solutions.EscapeV2()
	// solutions.CalculateDistance()
	runtime.GOMAXPROCS(runtime.NumCPU())
	solutions.GetLengthOfLookAndSay()
}
