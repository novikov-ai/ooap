package linked_list

type Contract interface {
	// Commands:

	// Head
	// Precondition: list is not empty
	// Postcondition: cursor set on the first node in the list
	Head()

	// Tail
	// Precondition: list is not empty
	// Postcondition: cursor set on the last node in the list
	Tail()

	// Right
	// Precondition: next node exists
	// Postcondition: cursor shifted 1 node right
	Right()

	// PutRight
	// Precondition: current node exists
	// Postcondition: new node with value inserted after current node
	PutRight(value any)

	// PutLeft
	// Precondition: current node exists
	// Postcondition: new node with value inserted before current node
	PutLeft(value any)

	// Remove
	// Precondition: current node exists
	// Postcondition: current node removed and cursor shifted right if node exists or left
	Remove()

	// Clear
	// Postcondition: linked list is empty
	Clear()

	// AddTail
	// Postcondition: new node inserted in tail
	AddTail(value any)

	// Replace
	// Precondition: current node exists
	// Postcondition: current node value replaced with given
	Replace(value any)

	// Find
	// Postcondition: cursor set to the next node with given value if exists
	Find(value any)

	// RemoveAll
	// Postcondition: all nodes which equal given value removed
	RemoveAll(value any)

	// Queries:

	// Get current node value
	// Precondition: list is not empty
	Get() any

	// Size Quantity of nodes
	Size() int

	// IsHead Cursor set at the start of the list
	IsHead() bool

	// IsTail Cursor set at the end of the list
	IsTail() bool

	// IsValue Cursor set on the node of the list
	IsValue() bool

	// Additional queries:

	GetRightStatus() bool
}
