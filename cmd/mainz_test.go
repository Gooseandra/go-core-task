package main

import (
	"reflect"
	"testing"
)

// почему то тесты неправильно работают, "For [] and [], expected [], got []", говорят, что fail, но функция работает правильно
func TestFindSliceDifference(t *testing.T) {
	tests := []struct {
		slice1   []string
		slice2   []string
		expected []string
		testName string
	}{
		{
			slice1:   []string{"a", "b", "c"},
			slice2:   []string{"b", "c", "d"},
			expected: []string{"a", "d"},
			testName: "Common elements with unique differences",
		},
		{
			slice1:   []string{"apple", "banana"},
			slice2:   []string{"banana", "cherry"},
			expected: []string{"apple", "cherry"},
			testName: "Two unique differences",
		},
		{
			slice1:   []string{"a", "b"},
			slice2:   []string{"a", "b"},
			expected: []string{},
			testName: "No differences",
		},
		{
			slice1:   []string{},
			slice2:   []string{"a", "b"},
			expected: []string{"a", "b"},
			testName: "Empty first slice",
		},
		{
			slice1:   []string{"a", "b"},
			slice2:   []string{},
			expected: []string{"a", "b"},
			testName: "Empty second slice",
		},
		{
			slice1:   []string{},
			slice2:   []string{},
			expected: []string{},
			testName: "Both slices empty",
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			result := FindSliceDifference(tt.slice1, tt.slice2)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("For %v and %v, expected %v, got %v", tt.slice1, tt.slice2, tt.expected, result)
			}
		})
	}
}
