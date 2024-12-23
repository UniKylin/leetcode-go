package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDailyTemperatures(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Normal case",
			input:    []int{73, 74, 75, 71, 69, 72, 76, 73},
			expected: []int{1, 1, 4, 2, 1, 1, 0, 0},
		},
		{
			name:     "All temperatures increasing",
			input:    []int{30, 40, 50, 60},
			expected: []int{1, 1, 1, 0},
		},
		{
			name:     "All temperatures decreasing",
			input:    []int{60, 50, 40, 30},
			expected: []int{0, 0, 0, 0},
		},
		{
			name:     "Temperatures with same value",
			input:    []int{30, 30, 30},
			expected: []int{0, 0, 0},
		},
		{
			name:     "Single element",
			input:    []int{42},
			expected: []int{0},
		},
		{
			name:     "Empty input",
			input:    []int{},
			expected: []int{},
		},
	}

	for _, test := range tests {
		result := DailyTemperatures(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			fmt.Printf("Test '%s' failed: input %v, expected %v, got %v\n",
				test.name, test.input, test.expected, result)
		} else {
			fmt.Printf("Test '%s' passed.\n", test.name)
		}
	}
}
