package native_dictionary

type Contract interface {
	// Put
	// Postcondition: a new pair of key-value added to the NativeDictionary or the current value of key is updated
	Put(key string, value any)

	// Remove
	// Precondition: the NativeDictionary is not empty and provided key existed
	// Postcondition: the provided key removed with it's value
	Remove(key string)

	// IsKey
	// Precondition: the NativeDictionary is not empty
	IsKey(key string) bool

	// Get
	// Precondition: the NativeDictionary is not empty and the key is presence
	Get(key string) any

	// Hash
	// Return index of the given value
	Hash(value any) int

	// Size
	// Get amount of inserted elements
	Size() int

	// GetPutStatus
	// Get status of the last executed Put command
	GetPutStatus() int

	// GetRemoveStatus
	// Get status of the last executed GetRemoveStatus command
	GetRemoveStatus() int

	// GetIsKeyStatus
	// Get status of the last executed GetIsKeyStatus command
	GetIsKeyStatus() int
}
