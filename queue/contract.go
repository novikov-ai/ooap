package queue

type Contract interface {
	Enqueue(value any)
	Dequeue() any
	Size() int
}
