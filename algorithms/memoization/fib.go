package memoization

func FibFast(n int) []int {
	seq := []int{0, 1}
	for i := 2; i < n; i++ {
		seq = append(seq, seq[i-2]+seq[i-1])
	}
	return seq
}
