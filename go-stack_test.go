package main

import (
	"fmt"
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

func TestAddOne(t *testing.T) {
	var stk *stack.Stack = stack.New()
	var firstElement *stack.Element = new(stack.Element)

	expFirstValue := "firstElementTestValue"
	expLen := 1
	firstElement.Value = expFirstValue
	stk.Add(firstElement)

	if stk.Len() != expLen {
		t.Errorf("got %v, want %v", stk.Len(), expLen)
	}
	if stk.Elements() == nil {
		t.Errorf("got %v, want %v", stk.Len(), nil)
	}

	poppedElement := stk.Pop()
	poppedValue := poppedElement.Value
	fmt.Printf("first value was %v\n", poppedValue)
	if poppedValue != expFirstValue {
		t.Errorf("got first value %v, want %v", poppedElement.Value, expFirstValue)

	}
}
