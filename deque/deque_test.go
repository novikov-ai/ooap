package deque

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeque_AddFront(t *testing.T) {
	testCases := []struct {
		name     string
		in       Deque
		elements []int
		expected Deque
	}{
		{
			name:     "added 2 elements to front of deque",
			in:       Deque{deque: []interface{}{1, 2, 3, 4, 5}},
			elements: []int{6, 7},
			expected: Deque{deque: []interface{}{7, 6, 1, 2, 3, 4, 5}},
		},
		{
			name:     "added 2 elements to an empty deque",
			in:       Deque{deque: []interface{}{}},
			elements: []int{6, 7},
			expected: Deque{deque: []interface{}{7, 6}},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			deque := test.in

			for _, el := range test.elements {
				deque.AddFront(el)
			}

			assert.Equal(t, test.expected, deque)
		})
	}
}

func TestDeque_AddTail(t *testing.T) {
	testCases := []struct {
		name     string
		in       Deque
		elements []int
		expected Deque
	}{
		{
			name:     "added 2 elements to tail of deque",
			in:       Deque{deque: []interface{}{1, 2, 3, 4, 5}},
			elements: []int{6, 7},
			expected: Deque{deque: []interface{}{1, 2, 3, 4, 5, 6, 7}},
		},
		{
			name:     "added 2 elements to an empty deque",
			in:       Deque{deque: []interface{}{}},
			elements: []int{6, 7},
			expected: Deque{deque: []interface{}{6, 7}},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			deque := test.in

			for _, el := range test.elements {
				deque.AddTail(el)
			}

			assert.Equal(t, test.expected, deque)
		})
	}
}

func TestDeque_RemoveFront(t *testing.T) {
	testCases := []struct {
		name     string
		in       Deque
		expected Deque
	}{
		{
			name: "removed the front element",
			in:   Deque{deque: []interface{}{1, 2, 3, 4, 5, 6, 7}},
			expected: Deque{
				deque:             []interface{}{2, 3, 4, 5, 6, 7},
				removeFrontStatus: RemoveTailStatusOK,
			},
		},
		{
			name: "removed the only 1 element",
			in:   Deque{deque: []interface{}{6}},
			expected: Deque{
				deque:             []interface{}{},
				removeFrontStatus: RemoveFrontStatusOK,
			},
		},
		{
			name: "try to remove from empty deque",
			in:   Deque{deque: []interface{}{}},
			expected: Deque{
				deque:             []interface{}{},
				removeFrontStatus: RemoveFrontStatusError,
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			deque := test.in

			deque.RemoveFront()

			assert.Equal(t, test.expected, deque)
		})
	}
}

func TestDeque_RemoveTail(t *testing.T) {
	testCases := []struct {
		name     string
		in       Deque
		expected Deque
	}{
		{
			name: "removed the tail element",
			in:   Deque{deque: []interface{}{1, 2, 3, 4, 5, 6, 7}},
			expected: Deque{
				deque:            []interface{}{1, 2, 3, 4, 5, 6},
				removeTailStatus: RemoveTailStatusOK,
			},
		},
		{
			name: "removed the only 1 element",
			in:   Deque{deque: []interface{}{6}},
			expected: Deque{
				deque:            []interface{}{},
				removeTailStatus: RemoveTailStatusOK,
			},
		},
		{
			name: "try to remove from empty deque",
			in:   Deque{deque: []interface{}{}},
			expected: Deque{
				deque:            []interface{}{},
				removeTailStatus: RemoveTailStatusError,
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			deque := test.in

			deque.RemoveTail()

			assert.Equal(t, test.expected, deque)
		})
	}
}

func TestDeque_Size(t *testing.T) {
	deque := New()

	assert.Equal(t, 0, deque.Size())

	for _, el := range []int{1, 2, 3, 4, 5} {
		deque.AddFront(el)
	}

	assert.Equal(t, 5, deque.Size())

	for i := 0; i < 3; i++ {
		deque.RemoveFront()
	}

	assert.Equal(t, 2, deque.Size())
}
