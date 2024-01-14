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

func TestHashTable_FindFreeSlot(t *testing.T) {
	tests := []struct {
		name      string
		hashTable HashTable
		value     int
		expected  int
	}{
		{
			name: "found free slot from Hash1",
			hashTable: HashTable{
				size:     10,
				count:    3,
				elements: []any{100, nil, nil, nil, nil, nil, nil, nil, nil, 0},
			},
			value:    77,
			expected: 5,
		},
		{
			name: "found free slot from Hash2",
			hashTable: HashTable{
				size:     10,
				count:    3,
				elements: []any{100, nil, nil, nil, nil, 15, nil, nil, nil, 0},
			},
			value:    77,
			expected: 3,
		},
		{
			name: "can't found free slot",
			hashTable: HashTable{
				size:     10,
				count:    3,
				elements: []any{100, 22, 333, 1, 2, 15, 555, 123, 444, 0},
			},
			value:    77,
			expected: SlotNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.hashTable.FindFreeSlot(tt.value)
			assert.Equal(t, tt.expected, got)
		})
	}
}
