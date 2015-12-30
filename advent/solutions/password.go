package solutions

import (
	"fmt"
	"regexp"
)

var passwordInput = []byte("hxbxxyzz")

//GenerateNewPassword generates a new password for Santa
func GenerateNewPassword() {
	fmt.Println(string(incrementalPasswordGenerate(passwordInput)))
}

func checkIncreasingTriplet(s []byte) bool {
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
	return !checkForbiddenLetters(s) && checkIncreasingTriplet(s) && checkNonOverlappingDifferentPairs(s)
}

func checkForbiddenLetters(s []byte) bool {
	var cannotContain = regexp.MustCompile(`(i|o|l)`)
	return cannotContain.Match(s)
}

func checkNonOverlappingDifferentPairs(s []byte) bool {
	pairCount := 0
	skip := false
	for i := range s {
		if skip {
			skip = false
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

func incrementPassword(passwd []byte, i int) []byte {
	//If passwd[i] == 'a' then... The next character also needs to increase
	//1. Am I an 'a'
	//2. Yes -> I need to tell the next guy to increment -> Call myself with i-1
	//3. Increment myself
	//1.a If not -> increment
	//If I'm the last character -> return me incremented if I need to.

	if i == 0 {
		passwd[i] -= 'a'
		passwd[i] = (passwd[i] + 1) % (('z' - 'a') + 1)
		passwd[i] += 'a'
		return passwd
	}

	if passwd[i] == 'a' {
		incrementPassword(passwd, i-1)
	}

	passwd[i] -= 'a'
	passwd[i] = (passwd[i] + 1) % (('z' - 'a') + 1)
	passwd[i] += 'a'
	return passwd
}

func incrementalPasswordGenerate(in []byte) []byte {
	pass := in
	for {
		pass = incrementPassword(pass, len(pass)-1)
		if checkCorrectness(pass) {
			return pass
		}
	}
}
