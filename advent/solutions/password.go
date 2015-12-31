package solutions

import (
	"fmt"
	"regexp"
)

var passwordInput = []byte("hxbxwxba")

//GenerateNewPassword generates a new password for Santa
func GenerateNewPassword() {
	generatedPassword := make(chan []byte, 100)
	checkedPassword := make(chan []byte)
	go incrementalPasswordGenerate(generatedPassword)
	go checkCorrectness(generatedPassword, checkedPassword)
	pass := <-checkedPassword
	fmt.Println(string(pass))
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

func checkCorrectness(input chan []byte, output chan []byte) {
	for in := range input {
		s := in
		checked := !checkForbiddenLetters(s) && checkIncreasingTriplet(s) && checkNonOverlappingDifferentPairs(s)
		if checked {
			output <- s
		}
	}
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

func incrementalPasswordGenerate(out chan []byte) {
	pass := passwordInput
	for {
		pass = incrementPassword(pass, len(pass)-1)
		out <- pass
	}
}
