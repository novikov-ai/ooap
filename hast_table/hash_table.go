package hast_table

const (
	PutNil = iota
	PutError
	PutOK

	SlotNotFound = -1
)

type HashTable struct {
	size      int
	elements  []any
	putStatus int
}

func New(size int) *HashTable {
	return &HashTable{
		size:      size,
		elements:  make([]any, size),
		putStatus: PutNil,
	}
}

// Commands:

// Put
// Precondition: current size is less than capacity & found a free slot
// Postcondition: new value inserted in the hashtable
func (ht *HashTable) Put(value any) bool {
	index := ht.FindFreeSlot(value)
	if len(ht.elements) == ht.size || index == SlotNotFound {
		ht.putStatus = PutError
		return false
	}

	ht.elements[index] = value

	ht.putStatus = PutOK
	return true
}

// Remove
// Postcondition: a given value removed from the hashtable if exists
func (ht *HashTable) Remove(value any) {
	index := ht.FindFreeSlot(value)
	if index == SlotNotFound {
		return
	}

	ht.elements[index] = nil
}

// Queries:

// Size
// Count elements of the hashtable
func (ht *HashTable) Size() int {
	return len(ht.elements)
}

// Hash
// Return index of the given value
func (ht *HashTable) Hash(value any) int {
	switch value.(type) {
	case int:
		v := value.(int)
		return v % (ht.size - 1)
	case string:
		v := value.(string)
		return len(v) % (ht.size - 1)
	}

	return 0
}

// Find
// Return slot index based on Hash
func (ht *HashTable) Find(value any) int {
	return ht.Hash(value)
}

// FindFreeSlot
// Return free slot index and SlotNotFound if all slots are occupied
func (ht *HashTable) FindFreeSlot(value any) int {
	if len(ht.elements) == ht.size {
		return SlotNotFound
	}

	hashIndex := ht.Hash(value)
	if ht.elements[hashIndex] == nil {
		return hashIndex
	}

	for i := 0; i <= ht.size; i++ {
		freeSlotIndex := (hashIndex + (ht.size - len(ht.elements))) % (ht.size - 1)
		if ht.elements[freeSlotIndex] == nil {
			return freeSlotIndex
		}
	}

	return SlotNotFound
}

// GetPutStatus
// Return status of the last executed command Put
func (ht *HashTable) GetPutStatus() int {
	return ht.putStatus
}
