package solutions

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

//Mine Mine MD5 hashes for Santa
func Mine() {
	src := "bgvyzdsv"
	i := 1

	//Though this is an infinite loop in essence
	for {
		temp := src
		temp += strconv.Itoa(i)
		sum := md5.Sum([]byte(temp))

		firstFiveDigit := fmt.Sprintf("%x", sum)
		if firstFiveDigit[:6] == "000000" {
			break
		}
		i++
	}

	fmt.Println(i)
}
