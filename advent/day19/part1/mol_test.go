package main

import (
	"fmt"
	"testing"
)

// func TestMoleculeCount(*testing.T) {
// 	molecule = "HOHOHO"
// 	replacements = map[string][]string{
// 		"H": []string{"HO", "OH"},
// 		"O": []string{"HH"},
// 	}
//
// 	replace()
// 	fmt.Println(combinations)
// }

func TestAllIndexesReturned(*testing.T) {
	molecule = "OOOOHdddOHddOH"
	indexes := allIndiciesForString("OH", molecule)
	fmt.Println(indexes)
}
