package deque

import (
	"ooap/abstract_data_types/queue"
)

type Contract interface {
	queue.Contract

	// AddFront
	// Postcondition: the element added to the front of deque
	AddFront(value any)

	// RemoveTail
	// Precondition: the deque is not empty
	// Postcondition: the element removed from the tail of deque
	RemoveTail() any

	// GetRemoveTailStatus
	// Got status of the last executed RemoveTail(value any) command
	GetRemovedTailStatus() int
}
