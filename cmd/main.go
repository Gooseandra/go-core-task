package main

import "fmt"

func FindSliceCoincidences(slice1, slice2 []int) ([]int, bool) {
	elementCount := make(map[int]int)

	for _, v := range slice1 {
		elementCount[v] = 1
	}

	for _, v := range slice2 {
		_, ok := elementCount[v]
		if !ok {
			elementCount[v] = 1
		} else {
			elementCount[v] = elementCount[v] + 1
		}
	}

	var result []int
	for key, val := range elementCount {
		if val == 2 {
			result = append(result, key)
		}
	}
	if len(result) == 0 {
		return nil, false
	}
	return result, true
}

func main() {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}
	fmt.Println(FindSliceCoincidences(a, b))
}
