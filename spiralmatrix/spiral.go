package spiralmatrix

import "fmt"

var orderedMatrix [][]int

//OrderSpiralMatrix orders a matrix
func OrderSpiralMatrix(spi [][]int) {
	orderedMatrix = make([][]int, 0)
	for len(spi) > 0 {
		spi = readAndRemoveFirstRow(spi)
		spi = readAndRemoveLastColumn(spi)
		spi = readAndRemoveLastRow(spi)
		spi = readAndRemoveFirstColumn(spi)
	}
	fmt.Println(orderedMatrix)
}

func readAndRemoveFirstRow(spi [][]int) [][]int {
	orderedMatrix = append(orderedMatrix, spi[0])
	spi = spi[:0+copy(spi[0:], spi[0+1:])]
	return spi
}

func readAndRemoveLastColumn(spi [][]int) [][]int {
	var app []int
	for _, v := range spi {
		app = append(app, v[len(v)-1])
	}
	if len(app) > 0 {
		orderedMatrix = append(orderedMatrix, app)
	}
	for i, v := range spi {
		spi[i] = spi[i][:len(v)-1+copy(spi[i][len(v)-1:], spi[i][len(v):])]
	}
	return spi
}

func readAndRemoveLastRow(spi [][]int) [][]int {
	lastRow := spi[len(spi)-1]
	var arr []int
	for i := len(lastRow) - 1; i >= 0; i-- {
		arr = append(arr, lastRow[i])
	}
	if len(arr) > 0 {
		orderedMatrix = append(orderedMatrix, arr)
	}
	spi = spi[:len(spi)-1+copy(spi[len(spi)-1:], spi[len(spi):])]
	return spi
}

func readAndRemoveFirstColumn(spi [][]int) [][]int {
	if len(spi) == 1 {
		orderedMatrix = append(orderedMatrix, spi[0])
		spi = make([][]int, 0)
		return spi
	}

	var app []int
	for i := len(spi) - 1; i >= 0; i-- {
		app = append(app, spi[i][0])
	}
	if len(app) > 0 {
		orderedMatrix = append(orderedMatrix, app)
	}
	for i := range spi {
		spi[i] = spi[i][:0+copy(spi[i][0:], spi[i][1:])]
	}
	return spi
}
