package hast_table

type Contract interface {
	Put(value any) int
	Remove(value any)
	Size() int
	Hash(value any) int
	Find(value any) int
	FindFreeSlot(value any) int
	GetPutStatus() int
}
