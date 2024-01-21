package power_set

import (
	"ooap/hast_table"
)

type Contract interface {
	hast_table.Contract

	ValueExists(value any) bool
	Intersection(with PowerSet) *PowerSet
	Union(with PowerSet) *PowerSet
	Difference(with PowerSet) *PowerSet
	IsSubset(with PowerSet) bool
}
