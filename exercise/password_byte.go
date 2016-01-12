package main

import "regexp"

var passwordInputChan = [8]byte{'h', 'x', 'b', 'x', 'x', 'y', 'z', 'z'}

//GenerateNewPasswordChan generates a new password for Santa
// func main() {
// 	generatedPassword := make(chan [8]byte, 100)
// 	correctPassword := make(chan [8]byte)
// 	defer close(generatedPassword)
// 	defer close(correctPassword)
// 	go incrementalPasswordGenerateChan(generatedPassword)
// 	go checkCorrectnessChan(generatedPassword, correctPassword)
// 	pass := <-correctPassword
// 	fmt.Println(string(pass[:]))
// }

func checkCorrectnessChan(input <-chan [8]byte, output chan<- [8]byte) {
	//Could use range here as it's an infinite loop anyways
	for {
		s := [8]byte(<-input)
		// time.Sleep(time.Second)
		// fmt.Println("Checking:", string(s))
		correct := !checkForbiddenLettersChan(s) && checkIncreasingTripletChan(s) && checkNonOverlappingDifferentPairsChan(s)
		if correct {
			// fmt.Println("Good password:", string(s))
			// Return and stop with the first good password that was found
			output <- s
			break
		}
	}
}

func incrementalPasswordGenerateChan(out chan<- [8]byte) {
	pass := passwordInputChan
	for {
		pass = incrementPasswordChan(pass, len(pass)-1)
		// fmt.Println("New password is:", string(pass))
		out <- pass
	}
}

func checkIncreasingTripletChan(s [8]byte) bool {
	for i := range s {
		if i+2 < len(s) {
			if s[i]+1 == s[i+1] && s[i]+2 == s[i+2] {
				return true
			}
		}
	}
	return false
}

func checkForbiddenLettersChan(s [8]byte) bool {
	var cannotContain = regexp.MustCompile(`(i|o|l)`)
	return cannotContain.Match(s[:])
}

func checkNonOverlappingDifferentPairsChan(s [8]byte) bool {
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

func incrementPasswordChan(passwd [8]byte, i int) [8]byte {
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
