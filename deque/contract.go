package deque

import "ooap/queue"

type Contract interface {
	queue.Contract

	AddFront(value any)
	RemoveTail() any
	GetRemovedTailStatus() int
}
