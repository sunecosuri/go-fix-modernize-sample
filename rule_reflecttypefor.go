package sample

import "reflect"

// ExampleReflectTypeFor demonstrates the "reflecttypefor" rule.
// go fix replaces reflect.TypeOf((*T)(nil)).Elem() with reflect.TypeFor[T]().
func ExampleReflectTypeFor() reflect.Type {
	return reflect.TypeOf((*error)(nil)).Elem()
}

func ExampleReflectTypeForInt() reflect.Type {
	return reflect.TypeOf((*int)(nil)).Elem()
}
