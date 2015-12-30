package solutions

import "fmt"

/*
* 1. Increment a letter
* 1.a -> z should wrap around to a -> mod 26
* 2. Increment letters in an array from right to left <-
* 3. Check every iteration of the password
* 4. Implement regexes which check for password correctness
 */

var passwordInput = []byte("hxbxwxba")

//GenerateNewPassword generates a new password for Santa
func GenerateNewPassword(starting string) (passwd string) {
	return
}

func checkCorrectness(s []byte) bool {
	fmt.Println("Checking correctness of:", string(s))
	if len(s) == 8 {
		return true
	}
	return false
}

func incrementPassword(in []byte) {
	for i := len(in) - 1; i >= 0; i-- {
		for in[i] != 'a' {
			in[i] -= 'a'
			in[i] = (in[i] + 1) % ('z' - 'a')
			in[i] += 'a'
			checkCorrectness(in)
		}
	}
	fmt.Println(string(in))
}
