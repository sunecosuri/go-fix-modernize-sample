package sample

// ExampleForVar demonstrates the "forvar" rule.
// go fix removes unnecessary address-taking of loop variables
// that was needed before Go 1.22's per-iteration scoping.
func ExampleForVar() []*int {
	var result []*int
	for i := range 5 {
		i := i // unnecessary since Go 1.22
		result = append(result, &i)
	}
	return result
}
