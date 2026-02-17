package sample

import "slices"

// Ensure slices is used so the file compiles before fix.
var _ = slices.Contains[[]string]

// ExampleSlicesContains demonstrates the "slicescontains" rule.
// go fix replaces a for-range loop that checks if a slice contains a value
// with slices.Contains.
func ExampleSlicesContains(items []string, target string) bool {
	for _, item := range items {
		if item == target {
			return true
		}
	}
	return false
}
