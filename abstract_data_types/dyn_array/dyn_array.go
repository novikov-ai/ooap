package dyn_array

const (
	DefaultCapacity = 10

	ExtensionValue = 1.618

	CollapsingValue  = 1.5
	OccupancyPercent = 0.25
)

const (
	RemoveNil = iota
	RemoveOK
	RemoveIndexOutOfRange
)

const (
	GetItemNil = iota
	GetItemOK
	GetItemArrayIsEmpty
	GetItemIndexOutOfRange
)

type DynArray struct {
	array            []any
	count            int
	capacity         int
	statusArrayEmpty bool
	statusRemove     int
	statusGetItem    int
}

// New
// Postcondition: a new empty dyn-array created with default capacity
func New() *DynArray {
	return &DynArray{
		array:            []any{},
		count:            0,
		capacity:         DefaultCapacity,
		statusArrayEmpty: true,
		statusRemove:     RemoveNil,
		statusGetItem:    GetItemNil,
	}
}

// Commands

// Append
// Postcondition: element inserted to the end of the array and the capacity increased if array was full
func (da *DynArray) Append(value any) {
	da.array = append(da.array, value)
	da.count++

	if da.count == da.capacity {
		da.capacity = int(ExtensionValue * float64(da.capacity))
	}

	da.statusArrayEmpty = false
}

// Remove
// Precondition: element exists at provided index
// Postcondition: element removed at provided index and the capacity decreased if occupancy is bellow than reference value
func (da *DynArray) Remove(index int) {
	if index >= da.count {
		da.statusRemove = RemoveIndexOutOfRange
		return
	}

	left := da.array[:index]
	da.array = append(left, da.array[index+1:]...)
	da.count--

	if float64(da.count/da.capacity) <= OccupancyPercent {
		decreasedCapacity := int(float64(da.capacity) / CollapsingValue)
		if decreasedCapacity < DefaultCapacity {
			decreasedCapacity = DefaultCapacity
		}

		da.capacity = decreasedCapacity
	}

	if da.count == 0 {
		da.statusArrayEmpty = true
	}

	da.statusRemove = RemoveOK
}

// Queries

// GetItem
// Precondition: the array is not empty & index <= size of array
// Postcondition: got item by provided index
func (da *DynArray) GetItem(index int) any {
	if da.count == 0 {
		da.statusGetItem = GetItemArrayIsEmpty
		return nil
	}

	if index >= da.count {
		da.statusGetItem = GetItemIndexOutOfRange
		return nil
	}

	da.statusGetItem = GetItemOK
	return da.array[index]
}

// Size
// Get items amount in the array
func (da *DynArray) Size() int {
	return da.count
}

// GetArrayEmptyStatus
// Check if array is empty
func (da *DynArray) GetArrayEmptyStatus() bool {
	return da.statusArrayEmpty
}

// GetItemStatus
// Get status of last executed command GetItem()
func (da *DynArray) GetItemStatus() int {
	return da.statusGetItem
}

// GetRemoveStatus
// Get status of last executed command Remove()
func (da *DynArray) GetRemoveStatus() int {
	return da.statusRemove
}
