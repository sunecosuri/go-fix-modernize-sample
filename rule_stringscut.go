package sample

import "strings"

// ExampleStringsCut demonstrates the "stringscut" rule.
// go fix replaces strings.Index + slice pattern with strings.Cut.
func ExampleStringsCut(s string) (string, string, bool) {
	i := strings.Index(s, "=")
	if i < 0 {
		return s, "", false
	}
	return s[:i], s[i+1:], true
}
