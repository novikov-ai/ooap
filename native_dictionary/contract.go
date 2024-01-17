package native_dictionary

type Contract interface {
	Put(key string, value any)
	IsKey(key string) bool
	Get(key string) any
	Hash(value any) int
}
