// +build !example
package stringset

const TestVersion = 2

type Set struct {
	s []string
}

func New() Set {
	return Set{}
}
func NewFromSlice([]string) Set {
	return Set{}
}
func (s Set) Add(string) { // modify s
}
func (s Set) Delete(string) { // modify s
}
func (s Set) Has(string) bool {
	return true
}
func (s Set) IsEmpty() bool {
	return true
}
func (s Set) Len() int {
	return 1
}
func (s Set) Slice() []string {
	return []string{}
}
func (s Set) String() string {
	return ""
}
func Equal(s1, s2 Set) bool {
	return true
}
func Subset(s1, s2 Set) bool { // return s1 ⊆ s2
	return true
}
func Disjoint(s1, s2 Set) bool {
	return true
}
func Intersection(s1, s2 Set) Set {
	return Set{}
}
func Union(s1, s2 Set) Set {
	return Set{}
}
func Difference(s1, s2 Set) Set { // return s1 ∖ s2
	return Set{}
}
func SymmetricDifference(s1, s2 Set) Set {
	return Set{}
}
