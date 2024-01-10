package deque

type Contract interface {
	AddFront(value any)
	AddTail(value any)
	RemoveFront() any
	RemoveTail() any
	Size() int
	GetRemoveFrontStatus() int
	GetRemoveTailStatus() int
}
