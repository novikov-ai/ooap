package hast_table

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHashTable_Hash(t *testing.T) {
	const hashTableSize = 12

	tests := []struct {
		name     string
		value    any
		expected int
	}{
		{
			name:     "hash from string",
			value:    "hello",
			expected: 5,
		},
		{
			name:     "hash from empty string",
			value:    "",
			expected: 0,
		},
		{
			name:     "hash from int",
			value:    162,
			expected: 8,
		},
		{
			name:     "hash from bool",
			value:    true,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ht := New(hashTableSize)

			got := ht.Hash(tt.value)
			assert.Equal(t, tt.expected, got)
		})
	}
}
