package queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueue_Enqueue(t *testing.T) {
	testCases := []struct {
		name     string
		in       Queue
		elements []int
		expected Queue
	}{
		{
			name:     "enqueue 2 elements",
			in:       Queue{queue: []interface{}{1, 2, 3, 4, 5}},
			elements: []int{6, 7},
			expected: Queue{queue: []interface{}{1, 2, 3, 4, 5, 6, 7}},
		},
		{
			name:     "enqueue 2 elements to an empty queue",
			in:       Queue{queue: []interface{}{}},
			elements: []int{6, 7},
			expected: Queue{queue: []interface{}{6, 7}},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			queue := test.in

			for _, el := range test.elements {
				queue.Enqueue(el)
			}

			assert.Equal(t, test.expected, queue)
		})
	}
}

func TestQueue_Dequeue(t *testing.T) {
	testCases := []struct {
		name             string
		in               Queue
		times            int
		expected         Queue
		expectedElements []any
	}{
		{
			name:             "dequeue 2 elements",
			in:               Queue{queue: []interface{}{1, 2, 3, 4, 5}},
			times:            2,
			expected:         Queue{queue: []interface{}{3, 4, 5}},
			expectedElements: []interface{}{1, 2},
		},
		{
			name:             "dequeue only 1 element",
			in:               Queue{queue: []interface{}{1}},
			times:            1,
			expected:         Queue{queue: []interface{}{}},
			expectedElements: []interface{}{1},
		},
		{
			name:             "dequeue element from empty queue",
			in:               Queue{queue: []interface{}{}},
			times:            1,
			expected:         Queue{queue: []interface{}{}},
			expectedElements: []interface{}{},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			queue := test.in

			gotElements := make([]interface{}, 0, 0)
			for i := 0; i < test.times; i++ {
				got := queue.Dequeue()
				if got == nil {
					continue
				}

				v := got.(interface{})
				gotElements = append(gotElements, v)
			}

			assert.Equal(t, test.expected, queue)
			assert.Equal(t, test.expectedElements, gotElements)
		})
	}
}

func TestQueue_Size(t *testing.T) {
	queue := New()

	assert.Equal(t, 0, queue.Size())

	for _, el := range []int{1, 2, 3, 4, 5} {
		queue.Enqueue(el)
	}

	assert.Equal(t, 5, queue.Size())

	for i := 0; i < 3; i++ {
		queue.Dequeue()
	}

	assert.Equal(t, 2, queue.Size())
}
