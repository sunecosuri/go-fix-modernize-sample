package sample_test

import (
	"context"
	"testing"
)

// TestTestingContext demonstrates the "testingcontext" rule.
// go fix replaces context.WithCancel(context.Background()) + defer cancel()
// with t.Context().
func TestTestingContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_ = ctx
}

func TestTestingContextTODO(t *testing.T) {
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()
	_ = ctx
}
