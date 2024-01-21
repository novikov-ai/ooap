package power_set

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPowerSet_Intersection(t *testing.T) {
	tests := []struct {
		name     string
		main     *PowerSet
		with     PowerSet
		expected *PowerSet
	}{
		{
			name:     "new set is NOT empty",
			main:     powerSetWithData([]any{1, 2, 3, 4, nil, nil, nil, nil, nil}),
			with:     *powerSetWithData([]any{3, 4, nil, nil, nil, nil, nil}),
			expected: powerSetWithData([]any{nil, nil, nil, 3, 4, nil, nil, nil, nil}),
		},
		{
			name:     "new set is empty",
			main:     powerSetWithData([]any{1, 2, 3, 4, nil, nil, nil, nil, nil}),
			with:     *powerSetWithData([]any{6, nil, nil, nil, nil, nil, nil}),
			expected: powerSetWithData([]any{nil, nil, nil, nil, nil, nil, nil, nil, nil}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.main.Intersection(tt.with)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestPowerSet_Union(t *testing.T) {
	tests := []struct {
		name     string
		main     *PowerSet
		with     PowerSet
		expected *PowerSet
	}{
		{
			name:     "new set is NOT empty",
			main:     powerSetWithData([]any{1, 2, 3, 4, nil, nil, nil, nil, nil, nil}),
			with:     *powerSetWithData([]any{3, 4, nil, nil, nil, nil, nil}),
			expected: powerSetWithData([]any{nil, 1, 2, 3, 4, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil}),
		},
		{
			name:     "new set is empty with sum size of both sets",
			main:     powerSetWithData([]any{nil, nil, nil}),
			with:     *powerSetWithData([]any{nil, nil, nil, nil, nil}),
			expected: powerSetWithData([]any{nil, nil, nil, nil, nil, nil, nil, nil}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.main.Union(tt.with)
			assert.Equal(t, tt.expected.elements, got.elements)
		})
	}
}

func TestPowerSet_Difference(t *testing.T) {
	tests := []struct {
		name     string
		main     *PowerSet
		with     PowerSet
		expected *PowerSet
	}{
		{
			name:     "new set is NOT empty",
			main:     powerSetWithData([]any{1, 2, 3, 4, nil, nil, nil, nil, nil, nil}),
			with:     *powerSetWithData([]any{3, 4, nil, nil, nil, nil, nil}),
			expected: powerSetWithData([]any{nil, 1, 2, nil, nil, nil, nil, nil, nil, nil}),
		},
		{
			name:     "new set is empty",
			main:     powerSetWithData([]any{nil, nil, nil}),
			with:     *powerSetWithData([]any{nil, 1, 2, nil, nil}),
			expected: powerSetWithData([]any{nil, nil, nil}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.main.Difference(tt.with)
			assert.Equal(t, tt.expected.elements, got.elements)
		})
	}
}

func TestPowerSet_IsSubset(t *testing.T) {
	tests := []struct {
		name     string
		main     *PowerSet
		with     PowerSet
		expected bool
	}{
		{
			name:     "second set is subset",
			main:     powerSetWithData([]any{1, 2, 3, 4, nil, nil, nil, nil, nil, nil}),
			with:     *powerSetWithData([]any{3, 4, nil, nil, nil, nil, nil}),
			expected: true,
		},
		{
			name:     "second set is NOT subset",
			main:     powerSetWithData([]any{nil, nil, nil}),
			with:     *powerSetWithData([]any{nil, 1, 2, nil, nil}),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.main.IsSubset(tt.with)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func powerSetWithData(values []any) *PowerSet {
	pw := New(len(values))
	for _, v := range values {
		pw.Put(v)

		if v == nil {
			pw.count--
		}
	}

	return pw
}
