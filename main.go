package main

import "fmt"

// Element definition
type Element struct {
	Value interface{}
}

// Stack definition
type Stack struct {
	elements []Element
	topElem  *Element
	len      int
}

func main() {
	stk := New()
	fmt.Print("At initialization...")
	stk.PrintLen()

	var e1 Element = Element{5}
	// fmt.Printf("Adding element with value %v...\n", e1.Value)
	stk.Add(&e1)
	// stk.PrintAll()

	var e2 Element = Element{9}
	// fmt.Printf("Adding element with value %v...\n", e2.Value)
	stk.Add(&e2)
	// stk.PrintAll()

	var e3 Element = Element{"Wow"}
	// fmt.Printf("Adding element with value %v...\n", e3.Value)
	stk.Add(&e3)
	stk.PrintAll()

	fmt.Println("Attempting to pop top element")
	pop1 := stk.Pop()
	fmt.Printf("Popped element with value %v\n", *pop1)

	fmt.Print("New stack comprises...")
	stk.PrintAll()
	stk.PrintTop()
}

// Pop returns the top value of the stack and removes it from the stack
func (s *Stack) Pop() (e *Element) {
	poppedElem := s.topElem

	if s.len >= 2 {
		s.elements = s.elements[1:s.len]
		s.len = s.len - 1
		topElemPtr := s.elements[0]
		s.topElem = &topElemPtr
	} else if s.len == 1 {
		s.elements = nil
		s.topElem = nil
		s.len = s.len - 1
	} else {
		fmt.Println("Cannot pop - no elements in stack")
		return s.topElem
	}

	return poppedElem
}

// PrintLen prints the length of the stack
func (s *Stack) PrintLen() {
	fmt.Printf("Len of stack is: %v\n", s.len)
}

// PrintAll prints all values in a stack
func (s *Stack) PrintAll() {
	fmt.Printf("Stack contains these values: %v\n", s.elements)
}

// PrintTop prints the top value in the stack
func (s *Stack) PrintTop() {
	fmt.Printf("Top element has value %v\n", s.topElem.Value)
}

// Add adds an element to the stack
func (s *Stack) Add(e *Element) {
	s.topElem = e
	var front []Element

	s.elements = append(append(front, *e), s.elements...)
	s.len = s.len + 1
}

// New will initialize a new stack
func New() *Stack { return new(Stack).Init() }

// Init will initialize the base length
func (s *Stack) Init() *Stack {
	s.len = 0
	return s
}
