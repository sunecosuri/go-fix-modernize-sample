package sample

// ExampleAny demonstrates the "any" rule.
// go fix replaces interface{} with any.
func ExampleAny(v interface{}) interface{} {
	m := map[string]interface{}{
		"key": v,
	}
	return m["key"]
}
