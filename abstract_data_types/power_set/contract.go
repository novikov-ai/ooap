package power_set

import (
	"ooap/abstract_data_types/hash_table"
)

type Contract interface {
	hash_table.Contract

	// ValueExists
	// Check whether given value exists in the set
	ValueExists(value any) bool

	// Intersection
	// Got new PowerSet with intersected elements
	Intersection(with PowerSet) *PowerSet

	// Union
	// Got new PowerSet with united elements
	Union(with PowerSet) *PowerSet

	// Difference
	// Got new PowerSet without elements intersected with given
	Difference(with PowerSet) *PowerSet

	// IsSubset
	// Is each element of the given PowerSet belongs to the main PowerSet
	IsSubset(with PowerSet) bool
}
