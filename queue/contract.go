package queue

type Contract interface {
	AddTail(value any)
	RemoveFront() any
	Size() int
	GetRemovedFrontStatus() int
}
