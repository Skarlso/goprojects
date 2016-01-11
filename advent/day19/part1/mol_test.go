package main

import (
	"fmt"
	"testing"
)

func TestMoleculeCount(*testing.T) {
	molecule = "HOH"
	replacements = map[string][]string{
		"H": []string{"HO", "OH"},
		"O": []string{"HH"},
	}

	replace()
	fmt.Println(combinations)
}

// func TestMoleculeCount(*testing.T) {
// 	molecule = "HOHOHO"
// 	replacements = map[string][]string{
// 		"H": []string{"HO", "OH"},
// 		"O": []string{"HH"},
// 	}
//
// 	replace()
// 	fmt.Println(combinations)
// 	fmt.Println(len(combinations))
// }
