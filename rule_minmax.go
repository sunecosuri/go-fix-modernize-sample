package sample

// ExampleMinMax demonstrates the "minmax" rule.
// go fix replaces if-then-assign min/max patterns with built-in min/max.
func ExampleMin(a, b int) int {
	x := a
	if b < x {
		x = b
	}
	return x
}

func ExampleMax(a, b int) int {
	x := a
	if b > x {
		x = b
	}
	return x
}
