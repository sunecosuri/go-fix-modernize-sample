package sample

import "strings"

// ExampleStringsCutPrefix demonstrates the "stringscutprefix" rule.
// go fix replaces strings.HasPrefix + strings.TrimPrefix with strings.CutPrefix.
func ExampleStringsCutPrefix(s string) (string, bool) {
	if strings.HasPrefix(s, "prefix_") {
		return strings.TrimPrefix(s, "prefix_"), true
	}
	return s, false
}

// ExampleStringsCutSuffix demonstrates the same for suffix.
// go fix replaces strings.HasSuffix + strings.TrimSuffix with strings.CutSuffix.
func ExampleStringsCutSuffix(s string) (string, bool) {
	if strings.HasSuffix(s, "_suffix") {
		return strings.TrimSuffix(s, "_suffix"), true
	}
	return s, false
}
