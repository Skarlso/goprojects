package solutions

//LightCoordinates The location of lights which will be increased
type LightCoordinates struct {
	x, y int
}

var lightgridV2 = make(map[LightCoordinates]int, 1000)
