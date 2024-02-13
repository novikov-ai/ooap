package hash_table

type Contract interface {
	// Put
	// Precondition: current size is less than capacity & found a free slot
	// Postcondition: new value inserted in the hashtable
	Put(value any) int

	// Remove
	// Precondition: a given value exists in the hashtable
	// Postcondition: a given value removed from the hashtable if exists
	Remove(value any)

	// Size
	// Count elements of the hashtable
	Size() int

	// Hash
	// Return index of the given value
	Hash(value any) int

	// Find
	// Return slot index based on Hash
	Find(value any) int

	// FindFreeSlot
	// Return free slot index and SlotNotFound if all slots are occupied
	FindFreeSlot(value any) int

	// GetPutStatus
	// Return status of the last executed command Put
	GetPutStatus() int

	// GetRemoveStatus
	// Return status of the last executed command Remove
	GetRemoveStatus() int
}
