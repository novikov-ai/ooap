package dyn_array

type Contract interface {
	Append(value any)
	Remove(index int)
	GetItem(index int) any
	Size() int
	GetArrayEmptyStatus() bool
	GetItemStatus() int
	GetRemoveStatus() int
}
