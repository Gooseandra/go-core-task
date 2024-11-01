package main

import (
	"reflect"
	"testing"
)

func TestFindSliceCoincidences(t *testing.T) {
	tests := []struct {
		slice1     []int
		slice2     []int
		expected   []int
		expectedOk bool
		testName   string
	}{
		{
			slice1:     []int{1, 2, 3},
			slice2:     []int{2, 3, 4},
			expected:   []int{2, 3},
			expectedOk: true,
			testName:   "Common elements with unique differences",
		},
		{
			slice1:     []int{5, 6, 7},
			slice2:     []int{6, 7, 8},
			expected:   []int{6, 7},
			expectedOk: true,
			testName:   "Two unique differences",
		},
		{
			slice1:     []int{1, 2, 3},
			slice2:     []int{1, 2, 3},
			expected:   []int{1, 2, 3},
			expectedOk: true,
			testName:   "No differences, all elements match",
		},
		{
			slice1:     []int{},
			slice2:     []int{10, 20},
			expected:   nil,
			expectedOk: false,
			testName:   "First slice empty",
		},
		{
			slice1:     []int{30, 40},
			slice2:     []int{},
			expected:   nil,
			expectedOk: false,
			testName:   "Second slice empty",
		},
		{
			slice1:     []int{},
			slice2:     []int{},
			expected:   nil,
			expectedOk: false,
			testName:   "Both slices empty",
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			result, ok := FindSliceCoincidences(tt.slice1, tt.slice2)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("For slices %v and %v, expected result %v, got %v", tt.slice1, tt.slice2, tt.expected, result)
			}

			if ok != tt.expectedOk {
				t.Errorf("For slices %v and %v, expected ok %v, got %v", tt.slice1, tt.slice2, tt.expectedOk, ok)
			}
		})
	}
}
