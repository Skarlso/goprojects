package stringset

import "strings"

//TestVersion TestVersion
const TestVersion = 2

//Set set
type Set map[string]bool

//New new set
func New() Set {
	return Set{}
}

//NewFromSlice new Set from Slice
func NewFromSlice(fromSlice []string) Set {
	set := New()
	for _, v := range fromSlice {
		set[v] = true
	}
	return set
}

//Add adds an element
func (s Set) Add(newS string) { // modify s
	s[newS] = true
}

//Delete delets an element
func (s Set) Delete(toDelete string) { // modify s
	delete(s, toDelete)
}

//Has has element
func (s Set) Has(has string) bool {
	_, ok := s[has]
	return ok
}

//IsEmpty is empty
func (s Set) IsEmpty() bool {
	if len(s) == 0 {
		return true
	}
	return false
}

//Len length of Set
func (s Set) Len() int {
	return len(s)
}

//Slice slice representation of map
func (s Set) Slice() (slice []string) {
	for k := range s {
		slice = append(slice, k)
	}
	return
}

func (s Set) String() string {
	str := "{"
	for k := range s {
		str += "\"" + k + "\", "
	}
	str = strings.Trim(str, ", ")
	str += "}"
	return str
}

//Equal equal method for Set
func Equal(s1, s2 Set) bool {
	if s1.IsEmpty() && s2.IsEmpty() {
		return true
	}

	if s1.IsEmpty() || s2.IsEmpty() {
		return false
	}

	for k := range s1 {
		if !s2.Has(k) {
			return false
		}
	}
	return true
}

//Subset Subset
func Subset(s1, s2 Set) bool { // return s1 ⊆ s2
	for k1 := range s1 {
		if !s2.Has(k1) {
			return false
		}
	}

	return true
}

//Disjoint Disjoint
func Disjoint(s1, s2 Set) bool {
	if s1.IsEmpty() && s2.IsEmpty() {
		return true
	}
	if s1.IsEmpty() || s2.IsEmpty() {
		return true
	}

	for k1 := range s1 {
		if s2.Has(k1) {
			return false
		}
	}

	return true
}

//Intersection Intersection
func Intersection(s1, s2 Set) Set {
	interSet := make(map[string]bool)

	for k := range s1 {
		if s2.Has(k) {
			interSet[k] = true
		}
	}

	return interSet
}

//Union Union
func Union(s1, s2 Set) (newSet Set) {
	newSet = make(map[string]bool)
	for k := range s1 {
		newSet[k] = true
	}

	for k := range s2 {
		newSet[k] = true
	}
	return newSet
}

//Difference Difference
func Difference(s1, s2 Set) Set { // return s1 ∖ s2
	diffSet := make(map[string]bool)

	for k := range s1 {
		if !s2.Has(k) {
			diffSet[k] = true
		}
	}

	return diffSet

}

//SymmetricDifference SymmetricDifference
func SymmetricDifference(s1, s2 Set) Set {
	symSet := make(map[string]bool)

	for k := range s1 {
		if !s2.Has(k) {
			symSet[k] = true
		}
	}

	for k := range s2 {
		if !s1.Has(k) {
			symSet[k] = true
		}
	}

	// This should work but it returns {} as not a good Union.
	// return Union(Difference(s1, s2), Difference(s1, s2))
	return symSet
}
