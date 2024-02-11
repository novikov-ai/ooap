package dyn_array

type Contract interface {
	// Append
	// Postcondition: element inserted to the end of the array and the capacity increased if array was full
	Append(value any)

	// Remove
	// Precondition: element exists at provided index
	// Postcondition: element removed at provided index and the capacity decreased if occupancy is bellow than reference value
	Remove(index int)

	// GetItem
	// Precondition: the array is not empty & index <= size of array
	// Postcondition: got item by provided index
	GetItem(index int) any

	// Size
	// Get items amount in the array
	Size() int

	// GetArrayEmptyStatus
	// Check if array is empty
	GetArrayEmptyStatus() bool

	// GetItemStatus
	// Get status of last executed command GetItem()
	GetItemStatus() int

	// GetRemoveStatus
	// Get status of last executed command Remove()
	GetRemoveStatus() int
}
