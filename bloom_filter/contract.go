package bloom_filter

type Contract interface {
	// Add
	// Postcondition: bytes have set to the mask according hash functions
	Add(value string)

	// IsValue
	// Check if value exists (possibility of false-positive result)
	IsValue(value string) bool

	// Hash1
	// Return hash-index of the given value
	Hash1(value string) int

	// Hash2
	// Return hash-index of the given value
	Hash2(value string) int

	// GetStatusIsValue
	// Get status of the last executed IsValue command
	GetStatusIsValue() int
}
