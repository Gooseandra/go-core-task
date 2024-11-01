package main

import (
	"testing"
)

func createChannel(values ...int) <-chan int {
	ch := make(chan int)
	go func() {
		for _, v := range values {
			ch <- v
		}
		close(ch)
	}()
	return ch
}

func TestMergeChannels(t *testing.T) {
	tests := []struct {
		channels []<-chan int
		expected []int
		testName string
	}{
		{
			channels: []<-chan int{createChannel(1, 2, 3), createChannel(4, 5, 6), createChannel(7, 8, 9)},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			testName: "Multiple channels with integers",
		},
		{
			channels: []<-chan int{createChannel()},
			expected: []int{},
			testName: "Single empty channel",
		},
		{
			channels: []<-chan int{createChannel(), createChannel(1, 2, 3), createChannel()},
			expected: []int{1, 2, 3},
			testName: "Empty channels with one non-empty",
		},
		{
			channels: []<-chan int{createChannel(), createChannel()},
			expected: []int{},
			testName: "All channels empty",
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			merged := MergeChannels(tt.channels...)
			result := collectChannelValues(merged)

			// Проверка, что все ожидаемые значения содержатся в результате
			if !containsAll(result, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func collectChannelValues(ch <-chan int) []int {
	var result []int
	for val := range ch {
		result = append(result, val)
	}
	return result
}

func containsAll(result, expected []int) bool {
	if len(result) != len(expected) {
		return false
	}
	count := make(map[int]int)
	for _, val := range result {
		count[val]++
	}
	for _, val := range expected {
		if count[val] == 0 {
			return false
		}
		count[val]--
	}
	return true
}
