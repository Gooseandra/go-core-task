package UseCase

import (
	"crypto/sha256"
	"encoding/hex"
	"testing"
)

func TestToString(t *testing.T) {
	assembler := StringAssembler{}

	tests := []struct {
		input    any
		expected string
	}{
		{123, "123"},
		{3.14, "3.14"},
		{"Hello, World!", "Hello, World!"},
		{true, "true"},
	}

	for _, test := range tests {
		result, _ := assembler.ToString(test.input)
		if result != test.expected {
			t.Errorf("ToString(%v) = %s; want %s", test.input, result, test.expected)
		}
	}
}

func TestStringStream(t *testing.T) {
	assembler := StringAssembler{}

	result := assembler.StringStream("Hello", ", ", "World!")
	expected := "Hello, World!"

	if result != expected {
		t.Errorf("StringStream() = %s; want %s", result, expected)
	}
}

func TestRuneSlice(t *testing.T) {
	assembler := StringAssembler{}

	result := assembler.RuneSlice("Hello")
	expected := []rune{'H', 'e', 'l', 'l', 'o'}

	if len(result) != len(expected) {
		t.Errorf("RuneSlice() length = %d; want %d", len(result), len(expected))
	}

	for i, r := range expected {
		if result[i] != r {
			t.Errorf("RuneSlice()[%d] = %c; want %c", i, result[i], r)
		}
	}
}

func TestAddSalt(t *testing.T) {
	assembler := StringAssembler{}
	assembler.cfg.Hex.Salt = "SALT"

	result := assembler.AddSalt("HelloWorld", 5)
	expected := "HelloSALTWorld"

	if result != expected {
		t.Errorf("AddSalt() = %s; want %s", result, expected)
	}
}

func TestHexRunes(t *testing.T) {
	assembler := StringAssembler{}

	input := []rune{'H', 'e', 'l', 'l', 'o'}
	expectedHash := sha256.Sum256([]byte(string(input)))
	expected := hex.EncodeToString(expectedHash[:])

	result := assembler.HexRunes(input)

	if result != expected {
		t.Errorf("HexRunes() = %s; want %s", result, expected)
	}
}
