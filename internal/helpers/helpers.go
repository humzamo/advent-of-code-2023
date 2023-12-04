package helpers

import "fmt"

// Intersection return the intersection between two comparable slices
func Intersection[T comparable](s1, s2 []T) (inter []T) {
	hash := make(map[T]struct{})
	for _, e := range s1 {
		hash[e] = struct{}{}
	}
	for _, e := range s2 {
		if _, ok := hash[e]; ok {
			inter = append(inter, e)
		}
	}
	return
}

// PrintStructWithFields prints a struct with field names
// This is useful for debugging!
func PrintStructWithFields(s any) {
	fmt.Printf("%+v\n", s)
}
