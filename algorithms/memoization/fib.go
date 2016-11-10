package memoization

var seq = []int{0, 1}

func FibFast(n int) []int {
	l := len(seq)
	if n < l {
		return seq[0:n]
	}
	for i := l; i < n; i++ {
		seq = append(seq, seq[i-2]+seq[i-1])
	}
	return seq
}
