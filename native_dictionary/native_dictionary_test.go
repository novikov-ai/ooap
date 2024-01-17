package native_dictionary

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNativeDictionary_Put(t *testing.T) {
	testCases := []struct {
		name     string
		key      string
		value    any
		expected *NativeDictionary
	}{
		{
			name:  "put new key",
			key:   "hello",
			value: 777,
			expected: &NativeDictionary{
				size:   10,
				keys:   []string{"", "", "", "", "", "hello", "", "", "", ""},
				values: []any{nil, nil, nil, nil, nil, 777, nil, nil, nil, nil},
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			nd := New(10)
			nd.Put(tt.key, tt.value)

			assert.Equal(t, tt.expected, nd)
		})
	}
}

func TestNativeDictionary_IsKey(t *testing.T) {
	testCases := []struct {
		name     string
		key      string
		expected bool
	}{
		{
			name:     "found key",
			key:      "hello",
			expected: true,
		},
		{
			name:     "can't find the key",
			key:      "BOO",
			expected: false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			nd := &NativeDictionary{
				size:   10,
				keys:   []string{"", "", "", "", "", "hello", "", "", "", ""},
				values: []any{nil, nil, nil, nil, nil, 777, nil, nil, nil, nil},
			}

			got := nd.IsKey(tt.key)

			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestNativeDictionary_Get(t *testing.T) {
	testCases := []struct {
		name          string
		key           string
		expectedValue any
	}{
		{
			name:          "got the value by key",
			key:           "hello",
			expectedValue: 777,
		},
		{
			name:          "can't get the value by key",
			key:           "BOO",
			expectedValue: nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			nd := &NativeDictionary{
				size:   10,
				keys:   []string{"", "", "", "", "", "hello", "", "", "", ""},
				values: []any{nil, nil, nil, nil, nil, 777, nil, nil, nil, nil},
			}

			got := nd.Get(tt.key)

			assert.Equal(t, tt.expectedValue, got)
		})
	}
}
