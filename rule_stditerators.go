package sample

import (
	"fmt"
	"reflect"
)

// ExampleStdIterators demonstrates the "stditerators" rule.
// go fix replaces for i := 0; i < x.NumMethod(); i++ { x.Method(i) }
// with for method := range x.Methods() {}.
func ExampleStdIterators(t reflect.Type) {
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Println(m.Name)
	}
}
