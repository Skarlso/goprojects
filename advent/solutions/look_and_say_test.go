package solutions

import "testing"

type LookAndSayTestCase struct {
	input    string
	expected string
}

var lookTests = []LookAndSayTestCase{
	{"1", "11"},
	{"11", "21"},
	{"21", "1211"},
	{"1211", "111221"},
	{"111221", "312211"},
}

func TestBasicLookAndSay(t *testing.T) {
	for _, test := range lookTests {
		actual := LookAndSay([]byte(test.input))
		if string(actual) != test.expected {
			t.Errorf("input:%s actual: %s; expected:%s", test.input, actual, test.expected)
		}
	}
}
