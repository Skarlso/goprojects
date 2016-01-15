package main

import "fmt"

const (
	searchRow    = 2981
	searchColumn = 3075
)

var grid = make([][]int, (searchRow+searchColumn)+1)

func main() {
	row := searchRow + searchColumn
	column := searchColumn + searchRow
	for i := 0; i <= row; i++ {
		grid[i] = make([]int, column+1)
	}

	grid[1][1] = 20151125
	previousCode := 20151125
	for i := 2; i <= row; i++ {
		innerI := i
		for j := 1; j <= column; j++ {
			grid[innerI][j] = generateNextCode(previousCode)
			previousCode = grid[innerI][j]
			innerI--
			if innerI < 1 {
				break
			}
		}
	}

	fmt.Println(grid[searchRow][searchColumn])
}

//generateNextCode generate the next code based on the previous code
func generateNextCode(prevCode int) (newCode int) {
	return (prevCode * 252533) % 33554393
}
