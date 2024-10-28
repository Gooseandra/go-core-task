package UseCase

import (
	"errors"
	"math/rand"
	"time"
)

type sliceAssembler struct{}

func NewSliceAssembler() *sliceAssembler {
	return &sliceAssembler{}
}

func (s sliceAssembler) NewRandomSlice(length, numFrom, numTo int) []int {
	rand.NewSource(time.Now().UnixNano())
	var slice []int
	for i := 0; i < length; i++ {
		slice = append(slice, rand.Intn(numTo)+numFrom)
	}
	return slice
}

func (s sliceAssembler) SliceExample(slice []int) []int {
	var res []int
	for _, v := range slice {
		if v%2 == 0 {
			res = append(res, v)
		}
	}
	return res
}

func (s sliceAssembler) AddElements(additionalNum int, slice []int) []int {
	return append(slice, additionalNum)
}

func (s sliceAssembler) CopySlice(slice []int) []int {
	dst := make([]int, len(slice))
	copy(dst, slice)
	return dst
}

func (s sliceAssembler) RemoveElement(index int, slice []int) ([]int, error) {
	if index >= len(slice) {
		return nil, errors.New("index out of range")
	}
	res := slice[:index]
	res = append(res, slice[index+1:]...)
	return res, nil
}
