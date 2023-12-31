package two_way_list

import "ooap/list"

type TwoWayList struct {
	list        *list.List
	cursor      int
	rightExists bool
}

// New Postcondition: new empty linked list created
func New() *TwoWayList {
	return &TwoWayList{
		list:        list.New(),
		cursor:      -1,
		rightExists: false,
	}
}

// Commands:

// Head
// Precondition: list is not empty
// Postcondition: cursor set on the first node in the list
func (ll *TwoWayList) Head() {
	if ll.list.Length() < 1 {
		return
	}

	ll.cursor = 0
}

// Tail
// Precondition: list is not empty
// Postcondition: cursor set on the last node in the list
func (ll *TwoWayList) Tail() {
	if ll.list.Length() < 1 {
		return
	}

	ll.cursor = ll.list.Length() - 1
}

// Left
// Precondition: list is not empty & previous node exists
// Postcondition: cursor shifted 1 node left
func (ll *TwoWayList) Left() {
	if ll.list.Length() < 1 {
		return
	}

	if ll.cursor <= 0 {
		return
	}

	ll.cursor -= 1
}

// Right
// Precondition: list is not empty & next node exists
// Postcondition: cursor shifted 1 node right
func (ll *TwoWayList) Right() {
	if ll.list.Length() < 1 && ll.cursor >= ll.list.Length()-1 {
		ll.rightExists = false
		return
	}

	ll.cursor += 1
	ll.rightExists = true
}

// PutRight
// Precondition: list is not empty
// Postcondition: new node with value inserted after current node
func (ll *TwoWayList) PutRight(value any) {
	if ll.list.Length() < 1 {
		return
	}

	insertingIndex := ll.cursor + 1
	if insertingIndex >= ll.list.Length() {
		return
	}

	ll.list.InsertAfter(insertingIndex, value)
}

// PutLeft
// Precondition: list is not empty
// Postcondition: new node with value inserted before current node
func (ll *TwoWayList) PutLeft(value any) {
	if ll.list.Length() < 1 {
		return
	}

	insertingIndex := ll.cursor - 1
	if insertingIndex < 0 {
		return
	}

	ll.list.InsertAfter(insertingIndex, value)
}

// Remove
// Postcondition: current node removed and cursor shifted right if node exists or left
func (ll *TwoWayList) Remove() {
	if ll.list.Length() < 1 {
		return
	}

	ll.list.RemoveAt(ll.cursor)

	ll.Right()

	if ll.GetRightStatus() {
		return
	}

	ll.Left()
}

// Clear
// Postcondition: linked list is empty
func (ll *TwoWayList) Clear() {
	if ll.list.Length() == 0 {
		return
	}

	for i := 0; i < ll.list.Length(); i++ {
		ll.list.RemoveAt(i)
	}
}

// AddTail
// Postcondition: new node inserted in tail
func (ll *TwoWayList) AddTail(value any) {
	ll.list.Append(value)
}

// Replace
// Precondition: list is not empty
// Postcondition: current node value replaced with given
func (ll *TwoWayList) Replace(value any) {
	if ll.list.Length() < 1 {
		return
	}

	ll.list.InsertAfter(ll.cursor, value)
	ll.list.RemoveAt(ll.cursor)
}

// Find
// Postcondition: cursor set to the next node with given value
func (ll *TwoWayList) Find(value any) {
	if ll.cursor+1 >= ll.list.Length() {
		return
	}

	for i := ll.cursor + 1; i < ll.list.Length(); i++ {
		if value == ll.list.IndexAt(i) {
			ll.cursor = i
			return
		}
	}
}

// RemoveAll
// Postcondition: all nodes which equal given value removed
func (ll *TwoWayList) RemoveAll(value any) {
	updatedList := list.New()
	for i := 0; i < ll.list.Length(); i++ {
		if ll.list.IndexAt(i) != value {
			updatedList.Append(value)
		}
	}

	ll.list = updatedList
}

// Queries:

// Get current node value
func (ll *TwoWayList) Get() any {
	if ll.cursor < 0 {
		return nil
	}

	return ll.list.IndexAt(ll.cursor)
}

// Size Quantity of nodes
func (ll *TwoWayList) Size() int {
	return ll.list.Length()
}

// IsHead Cursor set at the start of the list
func (ll *TwoWayList) IsHead() bool {
	return ll.cursor == 0
}

// IsTail Cursor set at the end of the list
func (ll *TwoWayList) IsTail() bool {
	return ll.cursor == ll.list.Length()-1
}

// IsValue Cursor set on the node of the list
func (ll *TwoWayList) IsValue() bool {
	return ll.cursor >= 0
}

// Additional queries:

func (ll *TwoWayList) GetRightStatus() bool {
	return ll.rightExists
}
