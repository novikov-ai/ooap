package linked_list

type LinkedList struct {
	// todo: implement
}

// New Postcondition: new empty linked list created
func New() *LinkedList {
	// todo: implement
	return &LinkedList{}
}

// Commands:

// Head Postcondition: cursor set on the first node in the list
func (ll *LinkedList) Head() {
	// todo: implement
}

// Tail Postcondition: cursor set on the last node in the list
func (ll *LinkedList) Tail() {
	// todo: implement
}

// Right Postcondition: cursor shifted 1 node right
func (ll *LinkedList) Right() {
	// todo: implement
}

// PutRight Postcondition: new node with value inserted after current node
func (ll *LinkedList) PutRight(value any) {
	// todo: implement
}

// PutLeft Postcondition: new node with value inserted before current node
func (ll *LinkedList) PutLeft(value any) {
	// todo: implement
}

// Remove Postcondition: current node removed and cursor shifted right if node exists or left
func (ll *LinkedList) Remove(value any) {
	// todo: implement
}

// Clear Postcondition: linked list is empty
func (ll *LinkedList) Clear() {
	// todo: implement
}

// AddTail Postcondition: new node inserted in tail
func (ll *LinkedList) AddTail(value any) {
	// todo: implement
}

// Replace Postcondition: current node value replaced with given
func (ll *LinkedList) Replace(value any) {
	// todo: implement
}

// Find Postcondition: cursor set to the next node with given value
func (ll *LinkedList) Find(value any) {
	// todo: implement
}

// RemoveAll Postcondition: all nodes which equal given value removed
func (ll *LinkedList) RemoveAll(value any) {
	// todo: implement
}

// Queries:

// Get current node value
func (ll *LinkedList) Get() any {
	// todo: implement
	return nil
}

// Size Quantity of nodes
func (ll *LinkedList) Size() int {
	// todo: implement
	return 0
}

// IsHead Cursor set at the start of the list
func (ll *LinkedList) IsHead() bool {
	// todo: implement
	return false
}

// IsTail Cursor set at the end of the list
func (ll *LinkedList) IsTail() bool {
	// todo: implement
	return false
}

// IsValue Cursor set on the node of the list
func (ll *LinkedList) IsValue() bool {
	// todo: implement
	return false
}

// Additional queries:

func (ll *LinkedList) GetCurrentNodeStatus() int {
	// todo: implement
	return 0
}

func (ll *LinkedList) GetNextNodeStatus() int {
	// todo: implement
	return 0
}

func (ll *LinkedList) GetCursorStatus() int {
	// todo: implement
	return 0
}
