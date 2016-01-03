package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	//420
	fileContent, err := ioutil.ReadFile("smaller_input.txt")
	if err != nil {
		panic(err)
	}
	// sum := 0
	var jsonMessage interface{}

	err = json.Unmarshal(fileContent, &jsonMessage)
	if err != nil {
		panic(err)
	}
	m := jsonMessage.(map[string]interface{})
	fmt.Println(m)
	//Recursive travel through the tree of maps.
}
