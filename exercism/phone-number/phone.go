package phonenumber

import "fmt"

//Number number
func Number(pnum string) (string, error) {
	if len(pnum) < 10 {
		return "", fmt.Errorf("Bla")
	}
	return "", nil
}

//AreaCode return area code from phone number
func AreaCode(pnum string) (string, error) {
	return pnum[:3], nil
}

//Format formats a given phone number
func Format(pnum string) (string, error) {
	return "", nil
}
