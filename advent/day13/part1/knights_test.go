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

	table = map[string][]map[string]int{
		"Alice": {
			map[string]int{"Bob": 54},
			map[string]int{"David": -2},
		},
	}

	for _, v := range testCases {
		actual := getLikeForTargetConnect([]byte(v.input), []byte(v.neighbour))
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
