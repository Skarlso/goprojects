package solutions

import (
	"fmt"
	"regexp"
)

var passwordInputChan = []byte("hxbxwxba")

//GenerateNewPasswordChan generates a new password for Santa
func GenerateNewPasswordChan() {
	generatedPassword := make(chan []byte, 100)
	correctPassword := make(chan []byte)
	defer close(generatedPassword)
	defer close(correctPassword)
	go incrementalPasswordGenerateChan(generatedPassword)
	go checkCorrectnessChan(generatedPassword, correctPassword)
	pass := <-correctPassword
	fmt.Println(string(pass))
}

func checkCorrectnessChan(input <-chan []byte, output chan<- []byte) {
	//Could use range here as it's an infinite loop anyways
	for in := range input {
		s := in
		fmt.Println("Checking:", string(s))
		correct := !checkForbiddenLettersChan(s) && checkIncreasingTripletChan(s) && checkNonOverlappingDifferentPairsChan(s)
		if correct {
			// fmt.Println("Good password:", string(s))
			// Return and stop with the first good password that was found
			output <- s
		}
	}
}

func checkIncreasingTripletChan(s []byte) bool {
	for i := range s {
		if i+2 < len(s) {
			if s[i]+1 == s[i+1] && s[i]+2 == s[i+2] {
				return true
			}
		}
	}
	return false
}

func checkForbiddenLettersChan(s []byte) bool {
	var cannotContain = regexp.MustCompile(`(i|o|l)`)
	return cannotContain.Match(s)
}

func checkNonOverlappingDifferentPairsChan(s []byte) bool {
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

func incrementPasswordChan(passwd []byte, i int) []byte {
	if i == 0 {
		passwd[i] -= 'a'
		passwd[i] = (passwd[i] + 1) % (('z' - 'a') + 1)
		passwd[i] += 'a'
		return passwd
	}

	if passwd[i] == 'z' {
		incrementPasswordChan(passwd, i-1)
	}

	passwd[i] -= 'a'
	passwd[i] = (passwd[i] + 1) % (('z' - 'a') + 1)
	passwd[i] += 'a'
	return passwd
}

func incrementalPasswordGenerateChan(out chan<- []byte) {
	pass := passwordInputChan
	for {
		pass = incrementPasswordChan(pass, len(pass)-1)
		fmt.Println("New password is:", string(pass))
		out <- pass
	}
}
