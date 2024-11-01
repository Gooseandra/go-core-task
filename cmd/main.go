package main

import "fmt"

func FindSliceDifference(slice1, slice2 []string) []string {
	elementCount := make(map[string]int)

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

	var result []string
	for key, val := range elementCount {
		if val == 1 {
			result = append(result, key)
		}
	}
	return result
}

func main() {
	slice1 := []string{"1", "2"}
	slice2 := []string{"1", "2"}
	fmt.Println(FindSliceDifference(slice1, slice2))
}
