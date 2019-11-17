package main

import "fmt"

// Element definition
type Element struct {
	Value interface{}
}

// Stack definition
type Stack struct {
	root Element
	len  int
}

func main() {
	stk := New()
	fmt.Printf("Len of stack is: %v\n", stk.len)
}

// New will initialize a new stack
func New() *Stack { return new(Stack).Init() }

// Init will initialize the base length
func (s *Stack) Init() *Stack {
	s.len = 0
	return s
}
