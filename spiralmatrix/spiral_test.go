package spiralmatrix

import (
	"fmt"
	"testing"
)

var spiUnEven = [][]int{{1, 2, 3},
	{5, 6, 4}}

var spi2 = [][]int{{1, 2, 3, 4},
	{12, 13, 14, 5},
	{11, 16, 15, 6},
	{10, 9, 8, 7}}

var spi3 = [][]int{{1, 2, 3, 4, 5},
	{16, 17, 18, 19, 6},
	{15, 24, 25, 20, 7},
	{14, 23, 22, 21, 8},
	{13, 12, 11, 10, 9}}

func Test3x3ArrayV3(*testing.T) {
	var spi1 = [][]int{{1, 2, 3},
		{8, 9, 4},
		{7, 6, 5}}

	fmt.Println("========V3=========")
	OrderMatrix(spi1)
}

//
// func TestUnEvenArrayV3(*testing.T) {
// 	fmt.Println("========V3UnEven=========")
// 	OrderMatrix(spiUnEven)
// 	fmt.Println("========V3UnEven=========")
// }

// func Test4x4Array(*testing.T) {
// 	OrderSpiralMatrix(spi2)
// }
//
// func Test5x5Array(*testing.T) {
// 	OrderSpiralMatrix(spi3)
// }
//
func Test5x5ArrayV3(*testing.T) {
	var spi3 = [][]int{{1, 2, 3, 4, 5},
		{16, 17, 18, 19, 6},
		{15, 24, 25, 20, 7},
		{14, 23, 22, 21, 8},
		{13, 12, 11, 10, 9}}
	fmt.Println("Spi3 currently before:", spi3)
	fmt.Println("========V3MultipleTimes=========")
	OrderMatrix(spi3)
	fmt.Println("========V3MultipleTimes=========")
	fmt.Println("Spi3 currently after: ", spi3)
	OrderMatrix(spi3)
	fmt.Println("Spi3 after second call:", spi3)
}

// func BenchmarkRosetta(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		OrderSpiralMatrixRosetta()
// 	}
// }
//
// func BenchmarkMine(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		OrderSpiralMatrix(spi3)
// 	}
// }
//
// func BenchmarkMineV3(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		OrderMatrix(spi3)
// 	}
// }
