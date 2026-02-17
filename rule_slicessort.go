package sample

import (
	"slices"
	"sort"
)

// Ensure slices is used so the file compiles before fix.
var _ = slices.Sort[[]int]
var _ = sort.Ints

// ExampleSlicesSort demonstrates the "slicessort" rule.
// go fix replaces sort.Slice/sort.Ints/sort.Strings with slices equivalents.
func ExampleSlicesSortInts(s []int) {
	sort.Ints(s)
}

func ExampleSlicesSortStrings(s []string) {
	sort.Strings(s)
}

func ExampleSlicesSortFunc(s []int) {
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
}
