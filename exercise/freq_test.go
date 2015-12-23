package main

import "testing"

var data = "Writing good tests is not trivial, but in many situations a lot of ground can be covered with table-driven tests: Each table entry is a complete test case with inputs and expected results, and sometimes with additional information such as a test name to make the test output easily readable. If you ever find yourself using copy and paste when writing a test, think about whether refactoring into a table-driven test or pulling the copied code out into a helper function might be a better option."

func BenchmarkLoopFrequencyCount(t *testing.B) {

	for i := 0; i < t.N; i++ {
		countLettersLoop(data)
	}
}

func BenchmarkRecursiveFrequencyCount(t *testing.B) {
	for i := 0; i < t.N; i++ {
		countLettersRecursive(data)
	}
}
