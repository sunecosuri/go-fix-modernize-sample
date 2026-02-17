package sample

import "sync"

// ExampleWaitGroup demonstrates the "waitgroup" rule.
// go fix replaces wg.Add(1) followed by go func() { defer wg.Done(); ... }()
// with wg.Go(func() { ... }).
func ExampleWaitGroup() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		// do work
		_ = "work"
	}()

	wg.Wait()
}
