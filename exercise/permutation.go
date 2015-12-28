package main

import "fmt"

var locations = []string{"London", "Dublin", "Belfast"}

func permute(s []string, n int) {
	if n == 1 {
		fmt.Println(s)
	}
	for i := 0; i < n; i++ {
		s[i], s[n-1] = s[n-1], s[i]
		permute(s, n-1)
		s[i], s[n-1] = s[n-1], s[i]
	}
}

func main() {
	permute(locations, len(locations))
}
