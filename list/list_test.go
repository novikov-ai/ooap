package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestList_RemoveAt(t *testing.T) {
	list := New()

	elements := 10
	for i := 0; i < elements; i++ {
		list.Append(i)
	}

	tests := []struct {
		name     string
		index    int
		expected []any
	}{
		{
			name:     "Remove first element",
			index:    0,
			expected: []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name:     "Remove last element",
			index:    -1,
			expected: []interface{}{0, 1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			name:     "Remove middle element",
			index:    5,
			expected: []interface{}{0, 1, 2, 3, 4, 6, 7, 8, 9},
		},
		{
			name:     "Remove out of range element",
			index:    15,
			expected: []interface{}{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			copiedList := list

			copiedList.RemoveAt(test.index)
			assert.Equal(t, test.expected, copiedList.list)
		})
	}
}

func TestList_Length(t *testing.T) {
	list := New()

	elements := 10
	for i := 0; i < elements; i++ {
		list.Append(i)
	}

	assert.Equal(t, elements, list.Length())
}

func TestList_InsertAfter(t *testing.T) {
	tests := []struct {
		name     string
		index    int
		value    any
		expected []any
	}{
		{
			name:     "Insert in the middle",
			index:    5,
			value:    56,
			expected: []interface{}{0, 1, 2, 3, 4, 5, 56, 6, 7, 8, 9, 10},
		},
		{
			name:     "Insert after the head",
			index:    0,
			value:    100,
			expected: []interface{}{0, 100, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name:     "Insert after the tail",
			index:    10,
			value:    90,
			expected: []interface{}{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 90},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			list := New()

			elements := 10
			for i := 0; i <= elements; i++ {
				list.Append(i)
			}

			list.InsertAfter(test.index, test.value)
			assert.Equal(t, test.expected, list.list)
		})
	}
}
