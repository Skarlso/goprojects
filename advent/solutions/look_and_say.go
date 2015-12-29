package solutions

import (
	"fmt"
	"strconv"
)

const (
	//LIMIT limit
	LIMIT = 50
)

//INPUT puzzle input
//This used to be a string until I was reminded that BYTE ARRAY IS ALWAYS FASTER!
var INPUT = []byte("1321131112")

//LookAndSay translates numbers according to Look and Say algo
func LookAndSay(s []byte) (look []byte) {
	charCount := 1
	for i := range s {
		if i+1 < len(s) {
			if s[i] == s[i+1] {
				charCount++
			} else {
				// look += []byte(fmt.Sprintf("%d%s", charCount, string(s[i])))
				b := []byte(strconv.FormatInt(int64(charCount), 10))
				look = append(look, b[0], s[i])
				charCount = 1
			}
		} else {
			b := []byte(strconv.FormatInt(int64(charCount), 10))
			look = append(look, b[0], s[i])
			// look += []byte(fmt.Sprintf("%d%s", charCount, string(s[i])))
		}
	}
	return
}

//GetLengthOfLookAndSay Retrieve the Length of a lookandsay done Limit times
func GetLengthOfLookAndSay() {
	finalString := INPUT
	for i := 0; i <= LIMIT-1; i++ {
		finalString = LookAndSay(finalString)
	}
	fmt.Println("Lenght of final String:", len(finalString))
}
