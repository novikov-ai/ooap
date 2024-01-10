package deque

const (
	RemoveFrontStatusNil = iota
	RemoveFrontStatusError
	RemoveFrontStatusOK
)

const (
	RemoveTailStatusNil = iota
	RemoveTailStatusError
	RemoveTailStatusOK
)

type Deque struct {
	deque             []any
	removeFrontStatus int
	removeTailStatus  int
}

// New
// Postcondition: new empty deque created
func New() *Deque {
	return &Deque{
		deque:             []any{},
		removeFrontStatus: RemoveFrontStatusNil,
		removeTailStatus:  RemoveTailStatusNil,
	}
}

// Commands:

// AddFront
// Postcondition: the element added to the front of deque
func (d *Deque) AddFront(value any) {
	d.deque = append([]any{value}, d.deque...)
}

// AddTail
// Postcondition: the element added to the tail of deque
func (d *Deque) AddTail(value any) {
	d.deque = append(d.deque, value)
}

// RemoveFront
// Precondition: the deque is not empty
// Postcondition: the element removed from the front of deque
func (d *Deque) RemoveFront() {
	if len(d.deque) == 0 {
		d.removeFrontStatus = RemoveFrontStatusError
		return
	}

	if len(d.deque) == 1 {
		d.deque = []any{}
	} else {
		d.deque = d.deque[1:]
	}

	d.removeFrontStatus = RemoveFrontStatusOK
}

// RemoveTail
// Precondition: the deque is not empty
// Postcondition: the element removed from the tail of deque
func (d *Deque) RemoveTail() {
	if len(d.deque) == 0 {
		d.removeTailStatus = RemoveTailStatusError
		return
	}

	if len(d.deque) == 1 {
		d.deque = []any{}
	} else {
		d.deque = d.deque[:len(d.deque)-1]
	}

	d.removeTailStatus = RemoveTailStatusOK
}

// Queries:

// Size
// Got length of the deque
func (d *Deque) Size() int {
	return len(d.deque)
}

// GetRemoveFrontStatus
// Got status of the last executed RemoveFront(value any) command
func (d *Deque) GetRemoveFrontStatus() int {
	return d.removeFrontStatus
}

// GetRemoveTailStatus
// Got status of the last executed RemoveTail(value any) command
func (d *Deque) GetRemoveTailStatus() int {
	return d.removeTailStatus
}
