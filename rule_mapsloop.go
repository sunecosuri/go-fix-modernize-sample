package sample

// ExampleMapsLoop demonstrates the "mapsloop" rule.
// go fix replaces for k, v := range src { dst[k] = v } with maps.Copy(dst, src).
func ExampleMapsLoopCopy(dst, src map[string]int) {
	for k, v := range src {
		dst[k] = v
	}
}
