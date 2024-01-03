package dyn_array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDynArray_Remove(t *testing.T) {
	testCase := []struct {
		name         string
		index        int
		input        []any
		expected     []any
		decreasedCap int
	}{
		{
			name:         "remove the only one item",
			index:        0,
			input:        []any{1},
			expected:     []any{},
			decreasedCap: DefaultCapacity,
		},
		{
			name:         "remove the middle item",
			index:        3,
			input:        []any{0, 1, 2, 3, 4, 5},
			expected:     []any{0, 1, 2, 4, 5},
			decreasedCap: DefaultCapacity,
		},
		{
			name:         "remove the first item",
			index:        0,
			input:        []any{0, 1, 2, 3, 4, 5},
			expected:     []any{1, 2, 3, 4, 5},
			decreasedCap: DefaultCapacity,
		},
		{
			name:         "remove the last item",
			index:        5,
			input:        []any{0, 1, 2, 3, 4, 5},
			expected:     []any{0, 1, 2, 3, 4},
			decreasedCap: DefaultCapacity,
		},
	}

	for _, test := range testCase {
		t.Run(test.name, func(t *testing.T) {
			da := New()

			for _, item := range test.input {
				da.Append(item)
			}

			da.Remove(test.index)

			assert.Equal(t, test.expected, da.array)
			assert.Equal(t, test.decreasedCap, da.capacity)
		})
	}
}
