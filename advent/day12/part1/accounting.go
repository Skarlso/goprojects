package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fileContent, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(fileContent)
}
