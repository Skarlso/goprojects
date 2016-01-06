package main

import "testing"

func TestConnectionRetrieve(t *testing.T) {
	testCases := []struct {
		input        string
		neighbour    string
		expectedLike int
	}{
		{"Alice", "Bob", 54},
		{"Alice", "David", -2},
	}

	table = map[int][]map[int]int{
		1: {
			map[int]int{2: 54},
			map[int]int{3: -2},
		},
	}

	nameMapping = map[string]int{
		"Alice": 1,
		"Bob":   2,
		"David": 3,
	}

	for _, v := range testCases {
		actual := getLikeForNeighbour(nameMapping[v.input], nameMapping[v.neighbour])
		if actual != v.expectedLike {
			t.Errorf("Actual was: %d, Expected is:%d. For name: %s, with connection: %s\n", actual, v.expectedLike, v.input, v.neighbour)
		}
	}
}

func BenchmarkCalculateSeating(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculatePerfectSeating()
	}
}
