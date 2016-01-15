package main

import "testing"

func TestGenerateNextCode(t *testing.T) {
	base := 20151125
	expected := 31916031
	actual := generateNextCode(base)

	if actual != expected {
		t.Errorf("Actual: %d; Did not equal Expected:%d", actual, expected)
	}
}
