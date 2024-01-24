package bloom_filter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBloomFilter_Add(t *testing.T) {
	tests := []struct {
		name                string
		value               string
		bloomFilter         BloomFilter
		expectedBloomFilter BloomFilter
	}{
		{
			name:  "ok",
			value: "123456",
			bloomFilter: BloomFilter{
				byteMask:      []bool{false, false, false, false, false, false, false, false, false, false},
				isValueStatus: 0,
			},
			expectedBloomFilter: BloomFilter{
				byteMask:      []bool{false, false, false, true, false, false, true, false, false, false},
				isValueStatus: 0,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.bloomFilter.Add(tt.value)
			assert.Equal(t, tt.expectedBloomFilter, tt.bloomFilter)
		})
	}
}

func TestBloomFilter_IsValue(t *testing.T) {
	tests := []struct {
		name                string
		value               string
		bloomFilter         BloomFilter
		expectedBloomFilter BloomFilter
		expected            bool
	}{
		{
			name:  "found",
			value: "123456",
			bloomFilter: BloomFilter{
				byteMask:      []bool{false, false, false, true, false, false, true, false, false, false},
				isValueStatus: 0,
			},
			expectedBloomFilter: BloomFilter{
				byteMask:      []bool{false, false, false, true, false, false, true, false, false, false},
				isValueStatus: 0,
			},
			expected: true,
		},
		{
			name:  "false positive",
			value: "654321",
			bloomFilter: BloomFilter{
				byteMask:      []bool{false, false, false, true, false, false, true, false, false, false},
				isValueStatus: 0,
			},
			expectedBloomFilter: BloomFilter{
				byteMask:      []bool{false, false, false, true, false, false, true, false, false, false},
				isValueStatus: 0,
			},
			expected: true,
		},
		{
			name:  "not found",
			value: "43332",
			bloomFilter: BloomFilter{
				byteMask:      []bool{false, false, false, true, false, false, true, false, false, false},
				isValueStatus: 0,
			},
			expectedBloomFilter: BloomFilter{
				byteMask:      []bool{false, false, false, true, false, false, true, false, false, false},
				isValueStatus: IsValueNotFound,
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.bloomFilter.IsValue(tt.value)
			assert.Equal(t, tt.expected, got)
			assert.Equal(t, tt.expectedBloomFilter, tt.bloomFilter)
		})
	}
}
