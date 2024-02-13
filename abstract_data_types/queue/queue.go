package queue

const (
	RemovedFrontStatusNil = iota
	RemovedFrontStatusError
	RemovedFrontStatusOK
)

type Queue struct {
	queue              []any
	removedFrontStatus int
}

// New
// Postcondition: new empty queue created
func New() *Queue {
	return &Queue{
		queue:              []any{},
		removedFrontStatus: RemovedFrontStatusNil,
	}
}

// Commands:

// AddTail
// Postcondition: the element added to the queue tail
func (q *Queue) AddTail(value any) {
	q.queue = append(q.queue, value)
}

// RemoveFront
// Precondition: queue is not empty
// Postcondition: head element dequeueded from the queue & it's size reduced by 1
func (q *Queue) RemoveFront() any {
	if len(q.queue) == 0 {
		q.removedFrontStatus = RemovedFrontStatusError
		return nil
	}

	dequeued := q.queue[0]
	q.queue = q.queue[1:]

	q.removedFrontStatus = RemovedFrontStatusOK

	return dequeued
}

// Queries:

// Size
// Got length of the queue
func (q *Queue) Size() int {
	return len(q.queue)
}

// GetRemovedFrontStatus
// Got status of the last executed RemovedFront() command
func (q *Queue) GetRemovedFrontStatus() int {
	return q.removedFrontStatus
}
