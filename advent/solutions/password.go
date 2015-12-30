package solutions

import (
	"fmt"
	"strings"
)

/*
* 1. Increment a letter
* 1.a -> z should wrap around to a -> mod 26
* 2. Increment letters in an array from right to left <-
* 3. Check every iteration of the password
* 4. Implement regexes which check for password correctness
 */

var passwordInput = []byte("hxbxwxba")

//GenerateNewPassword generates a new password for Santa
func GenerateNewPassword() {
	fmt.Println(string(incrementalPasswordGenerate(passwordInput)))
}

func checkIncreasingTriplet(s []byte) bool {
	fmt.Println("Checking for Increasing triplets in:", string(s))
	for i := range s {
		if i+2 < len(s) {
			if s[i]+1 == s[i+1] && s[i]+2 == s[i+2] {
				return true
			}
		}
	}
	return false
}

func checkCorrectness(s []byte) bool {
	return checkForbiddenLetters(string(s)) && checkIncreasingTriplet(s) && checkNonOverlappingDifferentPairs(s)
}

func checkForbiddenLetters(s string) bool {
	fmt.Println("Checking for forbidden letters in:", s)
	return strings.ContainsAny(s, "i & o & l")
}

func checkNonOverlappingDifferentPairs(s []byte) bool {
	fmt.Println("Checking for Non Overlapping Pairs in:", string(s))
	pairCount := 0
	skipCount := 0
	skip := false
	for i := range s {
		if skip {
			skipCount++
			if skipCount == 2 {
				skip = false
			}
			continue
		}

		if i+1 < len(s) {
			if s[i] == s[i+1] {
				pairCount++
				skip = true
			}
		}
	}

	return pairCount > 1
}

func incrementalPasswordGenerate(in []byte) []byte {
	for i := len(in) - 1; i >= 0; i-- {
		origin := in[i]
		in[i]++
		for in[i] != origin {
			in[i] -= 'a'
			in[i] = (in[i] + 1) % ('z' - 'a')
			in[i] += 'a'
			if checkCorrectness(in) {
				return in
			}
		}
	}

	return nil
}
