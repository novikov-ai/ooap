package bloom_filter

const (
	IsValueOK = iota
	IsValueNotFound
)

type BloomFilter struct {
	byteMask      []bool
	isValueStatus int
}

// New
// Created new empty bloomfilter with given bytemask size
func New(size int) *BloomFilter {
	return &BloomFilter{
		byteMask:      make([]bool, size),
		isValueStatus: IsValueOK,
	}
}

// Command:

// Add
// Postcondition: bytes have set to the mask according hash functions
func (b *BloomFilter) Add(value string) {
	b.byteMask[b.Hash1(value)] = true
	b.byteMask[b.Hash2(value)] = true
}

// Queries:

// IsValue
// Check if value exists (possibility of false-positive result)
func (b *BloomFilter) IsValue(value string) bool {
	if !b.byteMask[b.Hash1(value)] {
		b.isValueStatus = IsValueNotFound
		return false
	}

	if !b.byteMask[b.Hash2(value)] {
		b.isValueStatus = IsValueNotFound
		return false
	}

	b.isValueStatus = IsValueOK
	return true
}

// Hash1
// Return hash-index of the given value
func (b *BloomFilter) Hash1(value string) int {
	return len(value) % (len(b.byteMask) - 1)
}

// Hash2
// Return hash-index of the given value
func (b *BloomFilter) Hash2(value string) int {
	var hash uint32
	for _, char := range value {
		hash += uint32(char)
	}

	return int(hash) % (len(b.byteMask) - 1)
}

// GetStatusIsValue
// Get status of the last executed IsValue command
func (b *BloomFilter) GetStatusIsValue() int {
	return b.isValueStatus
}
