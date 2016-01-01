package solutions

import (
	"fmt"
	"regexp"
)

var passwordInputChan = []byte("hxbxxyzz")

//GenerateNewPasswordChan generates a new password for Santa
func GenerateNewPasswordChan() {
	//TODO: this did not work with a []byte channel. Why?
	//TODO: ANSWER => Because I'm using a []byte SLICE!! And the slice may be overwritten as per the documentation!
	//So using a primitive here works because that doesn't get overwritten!
	//SEE DOC HERE => https://golang.org/pkg/bufio/#Scanner.Bytes
	generatedPassword := make(chan string, 100)
	correctPassword := make(chan string)
	defer close(generatedPassword)
	defer close(correctPassword)
	go incrementalPasswordGenerateChan(generatedPassword)
	go checkCorrectnessChan(generatedPassword, correctPassword)
	pass := <-correctPassword
	fmt.Println(string(pass))
}

func checkCorrectnessChan(input <-chan string, output chan<- string) {
	//Could use range here as it's an infinite loop anyways
	for {
		s := []byte(<-input)
		// time.Sleep(time.Second)
		// fmt.Println("Checking:", string(s))
		correct := !checkForbiddenLettersChan(s) && checkIncreasingTripletChan(s) && checkNonOverlappingDifferentPairsChan(s)
		if correct {
			// fmt.Println("Good password:", string(s))
			// Return and stop with the first good password that was found
			output <- string(s)
			break
		}
	}
}

func incrementalPasswordGenerateChan(out chan<- string) {
	pass := passwordInputChan
	for {
		pass = incrementPasswordChan(pass, len(pass)-1)
		// fmt.Println("New password is:", string(pass))
		out <- string(pass)
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
