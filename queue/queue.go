package queue

const (
	DequeuededStatusNil = iota
	DequeuededStatusError
	DequeuededStatusOK
)

type Queue struct {
	queue         []any
	dequeueStatus int
}

// New
// Postcondition: new empty queue created
func New() *Queue {
	return &Queue{
		queue:         []any{},
		dequeueStatus: DequeuededStatusNil,
	}
}

// Commands:

// Enqueue
// Postcondition: the element added to the queue
func (q *Queue) Enqueue(value any) {
	q.queue = append(q.queue, value)
}

// Dequeue
// Precondition: queue is not empty
// Postcondition: head element dequeueded from the queue & it's size reduced by 1
func (q *Queue) Dequeue() any {
	if len(q.queue) == 0 {
		q.dequeueStatus = DequeuededStatusError
		return nil
	}

	dequeued := q.queue[0]
	q.queue = q.queue[1:]

	q.dequeueStatus = DequeuededStatusOK

	return dequeued
}

// Queries:

// Size
// Got length of the queue
func (q *Queue) Size() int {
	return len(q.queue)
}

// GetDequeuededStatus
// Got status of the last executed Dequeue() command
func (q *Queue) GetDequeuededStatus() int {
	return q.dequeueStatus
}
