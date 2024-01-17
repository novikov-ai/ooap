package native_dictionary

const (
	RemoveStatusOK = iota
	RemoveStatusKeyMissing
)

const (
	PutStatusOK = iota
	PutStatusValueOfCurrentKeyUpdated
)

const (
	IsKeyStatus = iota
	IsKeyStatusNotFound
)

type NativeDictionary struct {
	size         int
	count        int
	keys         []string
	values       []any
	removeStatus int
	putStatus    int
	isKeyStatus  int
}

// New
// Postcondition: a new empty NativeDictionary created
func New(size int) *NativeDictionary {
	return &NativeDictionary{
		size:         size,
		count:        0,
		keys:         make([]string, size),
		values:       make([]any, size),
		removeStatus: RemoveStatusOK,
		putStatus:    PutStatusOK,
		isKeyStatus:  IsKeyStatusNotFound,
	}
}

// Commands:

// Put
// Postcondition: a new pair of key-value added to the NativeDictionary or the current value of key is updated
func (nd *NativeDictionary) Put(key string, value any) {
	index := nd.Hash(key)

	exists := nd.IsKey(key)
	if exists {
		nd.putStatus = PutStatusValueOfCurrentKeyUpdated
	} else {
		nd.putStatus = PutStatusOK
	}

	nd.keys[index] = key
	nd.values[index] = value

	nd.count++

	return
}

// Remove
// Precondition: the NativeDictionary is not empty and provided key existed
// Postcondition: the provided key removed with it's value
func (nd *NativeDictionary) Remove(key string) {
	if !nd.IsKey(key) {
		nd.removeStatus = RemoveStatusKeyMissing
		return
	}

	index := nd.Hash(key)

	nd.keys[index] = ""
	nd.values[index] = nil

	nd.removeStatus = RemoveStatusOK
	nd.count--
}

// Queries:

// IsKey
// Precondition: the NativeDictionary is not empty
func (nd *NativeDictionary) IsKey(key string) bool {
	index := nd.Hash(key)

	exists := nd.keys[index] != ""
	if exists {
		nd.isKeyStatus = IsKeyStatus
	} else {
		nd.isKeyStatus = IsKeyStatusNotFound
	}

	return exists
}

// Get
// Precondition: the NativeDictionary is not empty and the key is presence
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

func (nd *NativeDictionary) Size() int {
	return nd.count
}

func (nd *NativeDictionary) GetPutStatus() int {
	return nd.putStatus
}

func (nd *NativeDictionary) GetRemoveStatus() int {
	return nd.removeStatus
}

func (nd *NativeDictionary) GetIsKeyStatus() int {
	return nd.isKeyStatus
}
