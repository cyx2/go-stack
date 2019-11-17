package stack

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

// Elements returns the array of elements in s
func (s *Stack) Elements() []Element {
	return s.elements
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
	if s.topElem != nil {
		fmt.Printf("Top element has value %v\n", s.topElem.Value)
	} else {
		fmt.Println("Stack is empty")
	}

}

// Len returns the length of the stack
func (s *Stack) Len() int {
	return s.len
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
