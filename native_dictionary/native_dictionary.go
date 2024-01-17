package native_dictionary

type NativeDictionary struct {
	size   int
	keys   []string
	values []any
}

// New
// Postcondition: a new empty NativeDictionary created
func New(size int) *NativeDictionary {
	return &NativeDictionary{
		size:   size,
		keys:   make([]string, size),
		values: make([]any, size),
	}
}

// Commands:

// Put
// Postcondition: a new pair of key-value added to the NativeDictionary
func (nd *NativeDictionary) Put(key string, value any) {
	index := nd.Hash(key)

	nd.keys[index] = key
	nd.values[index] = value

	return
}

// Queries:

// IsKey
// Precondition: the NativeDictionary is not empty
func (nd *NativeDictionary) IsKey(key string) bool {
	index := nd.Hash(key)
	return nd.keys[index] != ""
}

// Get
// Precondition: the NativeDictionary is not empty
func (nd *NativeDictionary) Get(key string) any {
	index := nd.Hash(key)
	return nd.values[index]
}

// Hash
// Return index of the given value
func (nd *NativeDictionary) Hash(value any) int {
	switch value.(type) {
	case int:
		v := value.(int)
		return v % (nd.size - 1)
	case string:
		v := value.(string)
		return len(v) % (nd.size - 1)
	}

	return 0
}
