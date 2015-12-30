package solutions

import (
	"fmt"
	"testing"
)

func TestIncrementingPasswordCharacterByCharacter(*testing.T) {
	incrementPassword(passwordInput)
}

func TestTriplet(t *testing.T) {
	sample := "aaabcaa"
	if !checkIncreasingTriplet([]byte(sample)) {
		t.Error("Triple was not found in:", sample)
	} else {
		t.Log("Triplet found in sample:", sample)
		fmt.Println("Triplet Found in sample:", sample)
	}

	errorSample := "abddd"
	if !checkIncreasingTriplet([]byte(errorSample)) {
		fmt.Println("Triplet not found.")
		t.Log("Triple was not found in:", errorSample)
	} else {
		t.Error("Triplet found in sample:", errorSample)
		fmt.Println("Triplet Found in sample:", errorSample)
	}
}

func TestForbiddenLetters(t *testing.T) {
	sample := "forbidden"

	if checkForbiddenLetters(sample) {
		fmt.Println("Found forbidden letters which is correct.")
	} else {
		t.Error("Did not find forbidden letters.")
	}

	errorSample := "abdsc"

	if checkForbiddenLetters(errorSample) {
		t.Error("Should have returned false for: ", errorSample)
	} else {
		fmt.Println("All Good.")
	}
}

func TestNonOverlappingPair(t *testing.T) {
	sample := "aacbb"

	if checkNonOverlappingDifferentPairs([]byte(sample)) {
		fmt.Println("All Good.")
	} else {
		t.Error("Did not find overlapping pairs in:", sample)
	}

	errorSample := "aabb"

	if checkNonOverlappingDifferentPairs([]byte(errorSample)) {
		t.Error("Found overlapping pairs in:", errorSample)
	} else {
		fmt.Println("All Good.")
	}
}
