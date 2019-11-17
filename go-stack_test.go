package main

import (
	"go-stack/stack"
	"testing"
)

func TestNew(t *testing.T) {
	var stk *stack.Stack = stack.New()
	expLen := 0
	if stk.Len() != expLen {
		t.Errorf("got %v, want %v", stk.Len(), expLen)
	}
	if stk.Elements() != nil {
		t.Errorf("got %v, want %v", stk.Len(), nil)
	}
}
