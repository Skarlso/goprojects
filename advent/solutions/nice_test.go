package solutions

import "testing"

type testCase struct {
	input    string
	expected bool
}

var niceTests = []testCase{
	{"qjhvhtzxzqqjkmpb", true},
	{"uurcxstgmygtbstg", false},
	{"ieodomkazucvgmuy", false},
	{"cyypypveppxxxfuq", true},
}

func TestSecondIteration(t *testing.T) {
	for _, test := range niceTests {
		actual := isNicePartTwo(test.input)
		if actual != test.expected {
			t.Errorf("Did not match expected output: actual: %v; expected:%v", actual, test.expected)
		}
	}
}
