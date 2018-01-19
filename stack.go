package stack

import (
	"sync"
)

// Stack is the struct which implments the methods Push, Pop, Top, IsEmpty
type Stack struct {
	Stack []interface{}
	sync.Mutex
}

// NewStack returns a new Stack struct
func NewStack() *Stack {
	return &Stack{}
}

// Push is adding the given element onto the stack
func (s *Stack) Push(item interface{}) {
	s.Lock()
	s.Stack = append(s.Stack, item)
	s.Unlock()
}

// Pop returns and removes the last element of the Stack
func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	s.Lock()
	item := s.Stack[len(s.Stack)-1]
	s.Stack = s.Stack[:len(s.Stack)-1]
	s.Unlock()
	return item
}

// Top returns the last element of the Stack
func (s *Stack) Top() interface{} {
	if s.IsEmpty() {
		return nil
	}
	s.Lock()
	item := s.Stack[len(s.Stack)-1]
	s.Unlock()
	return item
}

// IsEmpty returns true if the stack is empty or nil
func (s *Stack) IsEmpty() bool {
	if len(s.Stack) == 0 {
		return true
	}
	return false
}
