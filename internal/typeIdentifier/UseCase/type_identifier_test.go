package UseCase

import (
	"testing"
)

func TestIdentifyType(t *testing.T) {
	identifier := TypeIdentifier{}

	tests := []struct {
		input    any
		expected string
	}{
		{123, "int"},
		{3.14, "float64"},
		{"Hello", "string"},
		{true, "bool"},
		{[]int{1, 2, 3}, "[]int"},
		{map[string]int{"key": 1}, "map[string]int"},
		{nil, "nil"},
		{complex(1, 2), "complex128"},
		{struct{ A int }{1}, "struct { A int }"},
	}

	for _, test := range tests {
		result := identifier.IdentifyType(test.input)
		if result != test.expected {
			t.Errorf("IdentifyType(%v) = %s; want %s", test.input, result, test.expected)
		}
	}
}
