package native_dictionary

type Contract interface {
	Put(key string, value any)
	Remove(key string)
	IsKey(key string) bool
	Get(key string) any
	Hash(value any) int
	Size() int
	GetPutStatus() int
	GetRemoveStatus() int
	GetIsKeyStatus() int
}
