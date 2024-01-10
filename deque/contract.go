package deque

type Contract interface {
	AddFront(value any)
	AddTail(value any)
	RemoveFront()
	RemoveTail()
	Size() int
	GetRemoveFrontStatus() int
	GetRemoveTailStatus() int
}
