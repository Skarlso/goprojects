package spiralmatrix

import "testing"

var spi1 = [][]int{{1, 2, 3},
	{8, 9, 4},
	{7, 6, 5}}

var spi2 = [][]int{{1, 2, 3, 4},
	{12, 13, 14, 5},
	{11, 16, 15, 6},
	{10, 9, 8, 7}}

var spi3 = [][]int{{1, 2, 3, 4, 5},
	{16, 17, 18, 19, 6},
	{15, 24, 25, 20, 7},
	{14, 23, 22, 21, 8},
	{13, 12, 11, 10, 9}}

func Test3x3Array(*testing.T) {
	OrderSpiralMatrix(spi1)
}

func Test4x4Array(*testing.T) {
	OrderSpiralMatrix(spi2)
}

func Test5x5Array(*testing.T) {
	OrderSpiralMatrix(spi3)
}

func BenchmarkRosetta(b *testing.B) {
	for i := 0; i < b.N; i++ {
		OrderSpiralMatrixRosetta()
	}
}

func BenchmarkMine(b *testing.B) {
	for i := 0; i < b.N; i++ {
		OrderSpiralMatrix(spi3)
	}
}
