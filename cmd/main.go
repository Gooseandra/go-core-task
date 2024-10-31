package main

import "fmt"

type StringIntMap struct {
	data map[string]int
}

func NewStringIntMap() *StringIntMap {
	return &StringIntMap{
		data: make(map[string]int),
	}
}

func (m *StringIntMap) Add(key string, value int) {
	m.data[key] = value
}

func (m *StringIntMap) Remove(key string) {
	delete(m.data, key)
}

func (m *StringIntMap) Copy() map[string]int {
	newMap := make(map[string]int)
	for k, v := range m.data {
		newMap[k] = v
	}
	return newMap
}

func (m *StringIntMap) Exists(key string) bool {
	_, exists := m.data[key]
	return exists
}

func (m *StringIntMap) Get(key string) (int, bool) {
	value, exists := m.data[key]
	return value, exists
}

func main() {
	m := NewStringIntMap()

	m.Add("ine", 1)
	m.Add("two", 2)
	m.Add("three", 3)

	if value, exists := m.Get("1"); exists {
		fmt.Printf("Key: one, Value: %d\n", value)
	} else {
		fmt.Println("Key: one does not exist")
	}

	if m.Exists("2") {
		fmt.Println("Key: two exists")
	} else {
		fmt.Println("Key: two does not exist")
	}

	m.Remove("2")
	if !m.Exists("2") {
		fmt.Println("Key: two has been removed")
	}

	copiedMap := m.Copy()
	fmt.Println("Copied Map:", copiedMap)

	fmt.Println("Original Map:")
	for key, value := range m.Copy() {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}
}
