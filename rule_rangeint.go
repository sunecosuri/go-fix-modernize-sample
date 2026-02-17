package sample

// ExampleRangeInt demonstrates the "rangeint" rule.
// go fix replaces for i := 0; i < n; i++ {} with for i := range n {}.
func ExampleRangeInt(n int) int {
	sum := 0
	for i := 0; i < n; i++ {
		sum += i
	}
	return sum
}

func ExampleRangeIntUnused(n int) {
	for i := 0; i < n; i++ {
		_ = i
	}
}
