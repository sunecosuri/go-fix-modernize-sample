package sample

import "unsafe"

// ExampleUnsafeFuncs demonstrates the "unsafefuncs" rule.
// modernize replaces unsafe.Pointer(uintptr(ptr) + uintptr(n)) with unsafe.Add(ptr, n).
func ExampleUnsafeAdd(ptr unsafe.Pointer, n int) unsafe.Pointer {
	return unsafe.Pointer(uintptr(ptr) + uintptr(n))
}
