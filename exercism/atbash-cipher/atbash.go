package atbash

import (
	"strings"
	"unicode"
)

//DELIMITER delimiting block of text
const DELIMITER = " "

//Atbash implements atbash cypher
func Atbash(s string) (cypher string) {
	s = strings.ToLower(s)
	// s = strings.Replace(s, " ", "", -1)
	var code []string
	var block int
	for _, v := range s {
		if unicode.IsLetter(v) || unicode.IsNumber(v) {
			if block == 5 {
				code = append(code, DELIMITER)
				block = 0
			}

			if unicode.IsLetter(v) {
				index := v - 'a'
				cypherIndex := (('z' - 'a') - index)
				code = append(code, string(cypherIndex+'a'))
			} else {
				code = append(code, string(v))

			}

			block++
		}
	}
	cypher = strings.Join(code, "")
	return
}
