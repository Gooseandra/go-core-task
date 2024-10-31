package main

import "fmt"

type StringIntMap struct {
	key   string
	value int
}

type StringIntMapSlice struct {
	data []StringIntMap
}

func NewStringIntMapSlice() *StringIntMapSlice {
	return &StringIntMapSlice{
		data: make([]StringIntMap, 0),
	}
}

func (m *StringIntMapSlice) Add(key string, value int) {
	for i, item := range m.data {
		if item.key == key {
			m.data[i].value = value
			return
		}
	}
	m.data = append(m.data, StringIntMap{key, value})
}

func (m *StringIntMapSlice) Remove(key string) {
	for i, item := range m.data {
		if item.key == key {
			m.data = append(m.data[:i], m.data[i+1:]...)
			return
		}
	}
}

func (m *StringIntMapSlice) Copy() map[string]int {
	newMap := make(map[string]int)
	for _, item := range m.data {
		newMap[item.key] = item.value
	}
	return newMap
}

func (m *StringIntMapSlice) Exists(key string) bool {
	for _, item := range m.data {
		if item.key == key {
			return true
		}
	}
	return false
}

func (m *StringIntMapSlice) Get(key string) (int, bool) {
	for _, item := range m.data {
		if item.key == key {
			return item.value, true
		}
	}
	return 0, false
}

func main() {
	m := NewStringIntMapSlice()

	m.Add("one", 1)
	m.Add("two", 2)
	m.Add("three", 3)

	if value, exists := m.Get("one"); exists {
		fmt.Printf("Key: one, Value: %d\n", value)
	} else {
		fmt.Println("Key: one does not exist")
	}

	if m.Exists("two") {
		fmt.Println("Key: two exists")
	} else {
		fmt.Println("Key: two does not exist")
	}

	m.Add("two", 22)
	if value, exists := m.Get("two"); exists {
		fmt.Printf("Updated Key: two, New Value: %d\n", value)
	}

	m.Remove("two")
	if !m.Exists("two") {
		fmt.Println("Key: two has been removed")
	}

	copiedMap := m.Copy()
	fmt.Println("Copied Map:", copiedMap)

	fmt.Println("Original Map:")
	for key, value := range copiedMap {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}
}
