package power_set

const (
	RemoveStatusOK = iota
	RemoveStatusNotFound
)

const (
	PutStatusOK = iota
	PutStatusDuplicated
)

const SlotNotFound = -1
const ValueDuplicated = -2

type PowerSet struct {
	size         int
	count        int
	elements     []any
	removeStatus int
	putStatus    int
}

// New
// Postcondition: created an empty PowerSet with a given limit
func New(size int) *PowerSet {
	return &PowerSet{
		size:         size,
		count:        0,
		elements:     make([]any, size),
		removeStatus: RemoveStatusNotFound,
		putStatus:    PutStatusOK,
	}
}

// Commands:

// Remove
// Precondition: set is not empty and value exists
// Postcondition: value removed from set
func (pw *PowerSet) Remove(value any) {
	if len(pw.elements) == 0 || !pw.ValueExists(value) {
		pw.removeStatus = RemoveStatusNotFound
		return
	}

	index := pw.Hash(value)

	pw.elements[index] = nil
	pw.count--
	pw.removeStatus = RemoveStatusOK
}

// Put
// Precondition: current elements count less than size
// Postcondition: provided value put to the set
func (pw *PowerSet) Put(value any) int {
	if pw.count == pw.size || pw.ValueExists(value) {
		pw.putStatus = PutStatusDuplicated
		return ValueDuplicated
	}

	index := pw.Hash(value)

	pw.elements[index] = value
	pw.putStatus = PutStatusOK
	pw.count++

	return index
}

// Queries:

// Size
// Count elements of set
func (pw *PowerSet) Size() int {
	return pw.count
}

// Hash
// Got index of given value
func (pw *PowerSet) Hash(value any) int {
	switch value.(type) {
	case int:
		v := value.(int)
		return v % (pw.size - 1)
	case string:
		v := value.(string)
		return len(v) % (pw.size - 1)
	}

	return 0
}

// Find
// Precondition: set is not empty and given value exists
func (pw *PowerSet) Find(value any) int {
	return pw.Hash(value)
}

// FindFreeSlot
// Return free slot index and SlotNotFound if all slots are occupied
func (pw *PowerSet) FindFreeSlot(value any) int {
	if pw.count == pw.size {
		return SlotNotFound
	}

	hashIndex := pw.Hash(value)
	if pw.elements[hashIndex] == nil {
		return hashIndex
	}

	return SlotNotFound
}

// ValueExists
// Check whether given value exists in the set
func (pw *PowerSet) ValueExists(value any) bool {
	if value == nil {
		return false
	}

	index := pw.Find(value)
	return pw.elements[index] != nil
}

// GetPutStatus
// Return status of the last executed command Put
func (pw *PowerSet) GetPutStatus() int {
	return pw.putStatus
}

// GetRemoveStatus
// Return status of the last executed command Remove
func (pw *PowerSet) GetRemoveStatus() int {
	return pw.removeStatus
}

// Intersection
// Got new PowerSet with intersected elements
func (pw *PowerSet) Intersection(with PowerSet) *PowerSet {
	intersected := New(pw.size)

	for _, l := range pw.elements {
		if with.ValueExists(l) {
			intersected.Put(l)
		}
	}

	return intersected
}

// Union
// Got new PowerSet with united elements
func (pw *PowerSet) Union(with PowerSet) *PowerSet {
	united := New(pw.size + with.size)
	for _, l := range append(pw.elements, with.elements...) {
		if l == nil {
			continue
		}
		united.Put(l)
	}

	return united
}

// Difference
// Got new PowerSet without elements intersected with given
func (pw *PowerSet) Difference(with PowerSet) *PowerSet {
	diff := New(pw.size)

	for _, l := range pw.elements {
		if with.ValueExists(l) {
			continue
		}

		diff.Put(l)
	}

	return diff
}

// IsSubset
// Is each element of the given PowerSet belongs to the main PowerSet
func (pw *PowerSet) IsSubset(with PowerSet) bool {
	for _, l := range with.elements {
		if l == nil {
			continue
		}

		if !pw.ValueExists(l) {
			return false
		}
	}

	return true
}
