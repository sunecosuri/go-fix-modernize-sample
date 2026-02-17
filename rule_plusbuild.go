//go:build linux
// +build linux

package sample

// ExamplePlusBuild demonstrates the "plusbuild" rule.
// go fix removes the old-style // +build line when //go:build is present.
func ExamplePlusBuild() string {
	return "linux only"
}
