package bounded_stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_BoundedStack(t *testing.T) {
	// BOUNDED STACK MAX SIZE DEFAULT
	boundedStackDefault := New(nil)

	assert.Equal(t, 0, boundedStackDefault.Size())

	for i := 0; i < defaultStackMax+100; i++ {
		boundedStackDefault.Push(i)
	}

	assert.Equal(t, defaultStackMax, boundedStackDefault.Size())
	boundedStackDefault.Push(111)
	assert.Equal(t, defaultStackMax, boundedStackDefault.Size())

	boundedStackDefault.Clear()
	assert.Equal(t, 0, boundedStackDefault.Size())

	insertedValue := 111
	boundedStackDefault.Push(insertedValue)
	assert.Equal(t, 1, boundedStackDefault.Size())

	assert.Equal(t, PeekNil, boundedStackDefault.GetPeekStatus())

	peekedValue := boundedStackDefault.Peek()
	assert.Equal(t, insertedValue, peekedValue)
	assert.Equal(t, PeekOK, boundedStackDefault.GetPeekStatus())

	assert.Equal(t, PopNil, boundedStackDefault.GetPopStatus())

	boundedStackDefault.Pop()
	assert.Equal(t, 0, boundedStackDefault.Size())
	assert.Equal(t, PopOK, boundedStackDefault.GetPopStatus())

	boundedStackDefault.Pop()
	assert.Equal(t, PopErr, boundedStackDefault.GetPopStatus())

	boundedStackDefault.Peek()
	assert.Equal(t, PeekErr, boundedStackDefault.GetPeekStatus())

	// BOUNDED STACK MAX SIZE 10 (TEST CUSTOM BOUND)
	size10 := uint(10)
	boundedStackSize10 := New(&size10)

	assert.Equal(t, 0, boundedStackSize10.Size())

	for i := 0; i < 100; i++ {
		boundedStackSize10.Push(i)
	}

	assert.Equal(t, int(size10), boundedStackSize10.Size())
}
