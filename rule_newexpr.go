package sample

// ExampleNewExpr demonstrates the "newexpr" rule.
// go fix replaces func ptrOf(x T) *T { return &x } with return new(x),
// and inlines call sites: ptrOf(123) → new(123).

// intPtr is a common helper to get a pointer to a value.
func intPtr(v int) *int { return &v }

// stringPtr is another pointer helper.
func stringPtr(v string) *string { return &v }

func ExampleNewExpr() (*int, *string) {
	a := intPtr(42)
	b := stringPtr("hello")
	return a, b
}
