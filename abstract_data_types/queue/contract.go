package queue

type Contract interface {
	// AddTail
	// Postcondition: the element added to the queue tail
	AddTail(value any)

	// RemoveFront
	// Precondition: queue is not empty
	// Postcondition: head element dequeueded from the queue & it's size reduced by 1
	RemoveFront() any

	// Size
	// Got length of the queue
	Size() int

	// GetRemovedFrontStatus
	// Got status of the last executed RemovedFront() command
	GetRemovedFrontStatus() int
}
