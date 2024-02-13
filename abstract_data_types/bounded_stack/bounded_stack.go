package bounded_stack

import (
	"ooap/abstract_data_types/list"
)

const (
	PopNil = 0
	PopOK  = 1
	PopErr = 2

	PeekNil = 0
	PeekOK  = 1
	PeekErr = 2
)

const defaultStackMax = 32

type BoundedStack struct {
	bound      uint
	peekStatus int
	popStatus  int
	stack      *list.List
}

func New(bound *uint) *BoundedStack {
	bs := BoundedStack{
		stack: list.New(),
	}

	if bound == nil || bound != nil && *bound == 0 {
		bs.bound = defaultStackMax
	} else {
		bs.bound = *bound
	}

	return &bs
}

// Push
// Precondition: current size less than bound.
// Postcondition: new value added to stack.
func (bs *BoundedStack) Push(value any) {
	if bs.stack.Length() < int(bs.bound) {
		bs.stack.Append(value)
	}
}

// Clear
// Postcondition: Stack cleared, size is 0.
func (bs *BoundedStack) Clear() {
	bs.stack = list.New()
	bs.peekStatus = PeekNil
	bs.popStatus = PopNil
}

// Pop
// Precondition: Stack is not empty.
// Postcondition: last value removed from stack.
func (bs *BoundedStack) Pop() {
	if bs.stack.Length() > 0 {
		bs.stack.RemoveAt(-1)
		bs.popStatus = PopOK
	} else {
		bs.popStatus = PopErr
	}
}

// Peek
// Precondition: Stack is not empty.
// Postcondition: stack unchanged.
func (bs *BoundedStack) Peek() any {
	var result any

	if bs.stack.Length() > 0 {
		result = bs.stack.IndexAt(-1)
		bs.peekStatus = PeekOK
	} else {
		result = 0
		bs.peekStatus = PeekErr
	}

	return result
}

func (bs *BoundedStack) Size() int {
	return bs.stack.Length()
}

func (bs *BoundedStack) GetPopStatus() int {
	return bs.popStatus
}

func (bs *BoundedStack) GetPeekStatus() int {
	return bs.peekStatus
}
