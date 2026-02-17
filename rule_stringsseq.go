package sample

import "strings"

// ExampleStringsSeq demonstrates the "stringsseq" rule.
// go fix replaces strings.Split in a for-range with strings.SplitSeq (iterator).
func ExampleStringsSeq(s string) []string {
	var result []string
	for _, part := range strings.Split(s, ",") {
		result = append(result, part)
	}
	return result
}

func ExampleFieldsSeq(s string) []string {
	var result []string
	for _, field := range strings.Fields(s) {
		result = append(result, field)
	}
	return result
}
