package sample

import "fmt"

// ExampleFmtAppendf demonstrates the "fmtappendf" rule.
// go fix replaces []byte(fmt.Sprintf(...)) with fmt.Appendf(nil, ...).
func ExampleFmtAppendf(name string, age int) []byte {
	return []byte(fmt.Sprintf("name: %s, age: %d", name, age))
}

// ExampleFmtAppendSprint demonstrates another fmtappendf pattern.
// go fix replaces []byte(fmt.Sprint(...)) with fmt.Append(nil, ...).
func ExampleFmtAppendSprint(a, b int) []byte {
	return []byte(fmt.Sprint(a, b))
}
