package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	m := NewStringIntMapSlice()
	m.Add("key1", 10)

	value, exists := m.Get("key1")
	if !exists {
		t.Errorf("Expected key 'key1' to exist")
	}
	if value != 10 {
		t.Errorf("Expected value 10, got %d", value)
	}

	// Проверка обновления значения
	m.Add("key1", 20)
	value, _ = m.Get("key1")
	if value != 20 {
		t.Errorf("Expected updated value 20, got %d", value)
	}
}

func TestRemove(t *testing.T) {
	m := NewStringIntMapSlice()
	m.Add("key1", 10)
	m.Remove("key1")

	if m.Exists("key1") {
		t.Errorf("Expected key 'key1' to be removed")
	}
}

func TestExists(t *testing.T) {
	m := NewStringIntMapSlice()
	m.Add("key1", 10)

	if !m.Exists("key1") {
		t.Errorf("Expected key 'key1' to exist")
	}
	if m.Exists("key2") {
		t.Errorf("Expected key 'key2' to not exist")
	}
}

func TestGet(t *testing.T) {
	m := NewStringIntMapSlice()
	m.Add("key1", 10)

	value, exists := m.Get("key1")
	if !exists {
		t.Errorf("Expected key 'key1' to exist")
	}
	if value != 10 {
		t.Errorf("Expected value 10, got %d", value)
	}

	// Проверка отсутствующего ключа
	_, exists = m.Get("key2")
	if exists {
		t.Errorf("Expected key 'key2' to not exist")
	}
}

func TestCopy(t *testing.T) {
	m := NewStringIntMapSlice()
	m.Add("key1", 10)
	m.Add("key2", 20)

	copiedMap := m.Copy()

	// Проверка содержимого копии
	if len(copiedMap) != 2 {
		t.Errorf("Expected copied map to have 2 elements, got %d", len(copiedMap))
	}

	if copiedMap["key1"] != 10 {
		t.Errorf("Expected value for 'key1' in copied map to be 10, got %d", copiedMap["key1"])
	}

	if copiedMap["key2"] != 20 {
		t.Errorf("Expected value for 'key2' in copied map to be 20, got %d", copiedMap["key2"])
	}
}
