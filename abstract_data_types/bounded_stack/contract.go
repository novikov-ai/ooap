package bounded_stack

type Contract interface {
	// Push
	// Precondition: current size less than bound.
	// Postcondition: new value added to stack.
	Push(value any)

	// Clear
	// Postcondition: Stack cleared, size is 0.
	Clear()

	// Pop
	// Precondition: Stack is not empty.
	// Postcondition: last value removed from stack.
	Pop()

	// Peek
	// Precondition: Stack is not empty.
	// Postcondition: stack unchanged.
	Peek() any

	// Size
	// Quantity of elements
	Size() int

	// GetPopStatus
	// Get status of the last executed command
	GetPopStatus() int

	// GetPeekStatus
	// Get status of the last executed command
	GetPeekStatus() int
}
