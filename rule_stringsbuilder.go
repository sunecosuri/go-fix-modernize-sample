package sample

import "strings"

// ExampleStringsBuilder demonstrates the "stringsbuilder" rule.
// go fix replaces string concatenation in a loop with strings.Builder.
func ExampleStringsBuilder(parts []string) string {
	var result string
	for _, p := range parts {
		result += p
	}
	return result
}

// ExampleStringsBuilderJoin demonstrates another pattern.
// go fix may replace manual join with strings.Join.
func ExampleStringsBuilderJoin(parts []string) string {
	var result string
	for i, p := range parts {
		if i > 0 {
			result += ","
		}
		result += p
	}
	return result
}

// Keep strings import used.
var _ = strings.Builder{}
