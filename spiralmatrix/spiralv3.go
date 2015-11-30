package spiralmatrix

import "fmt"

//Matrix a matrix representation
type Matrix struct {
	matrix [][]int
}

//OrderMatrix orders a matrix
func OrderMatrix(spiV3 [][]int) {
	m := Matrix{}
	m.matrix = make([][]int, 0, len(spiV3))
	for _, v := range spiV3 {
		tmp := make([]int, len(v))
		copy(tmp, v)
		m.matrix = append(m.matrix, tmp)
	}
	for len(m.matrix) > 0 {
		m.rARFirstRow()
		m.rARLastColumn()
		m.rARLastRow()
		m.rARFirstColumn()
	}
	fmt.Println("Matrix matrix:", m.matrix)
	fmt.Println("spiV3:", spiV3)
}

func (m *Matrix) rARFirstRow() {
	if len(m.matrix) == 0 {
		return
	}
	fmt.Println(m.matrix[0])
	m.matrix = m.matrix[:0+copy(m.matrix[0:], m.matrix[0+1:])]
}

func (m *Matrix) rARLastColumn() {
	var app []int
	for _, v := range m.matrix {
		app = append(app, v[len(v)-1])
	}
	if len(app) > 0 {
		fmt.Println(app)
	}
	for i, v := range m.matrix {
		m.matrix[i] = m.matrix[i][:len(v)-1+copy(m.matrix[i][len(v)-1:], m.matrix[i][len(v):])]
	}
}

func (m *Matrix) rARLastRow() {
	lastRow := m.matrix[len(m.matrix)-1]
	var arr []int
	for i := len(lastRow) - 1; i >= 0; i-- {
		arr = append(arr, lastRow[i])
	}
	if len(arr) > 0 {
		fmt.Println(arr)
	}
	m.matrix = m.matrix[:len(m.matrix)-1+copy(m.matrix[len(m.matrix)-1:], m.matrix[len(m.matrix):])]
}

func (m *Matrix) rARFirstColumn() {
	if len(m.matrix) == 1 {
		fmt.Println(m.matrix[0])
		m.matrix = make([][]int, 0)
		return
	}

	var app []int
	for i := len(m.matrix) - 1; i >= 0; i-- {
		app = append(app, m.matrix[i][0])
	}
	if len(app) > 0 {
		fmt.Println(app)
	}
	for i := range m.matrix {
		m.matrix[i] = m.matrix[i][:0+copy(m.matrix[i][0:], m.matrix[i][1:])]
	}
}
