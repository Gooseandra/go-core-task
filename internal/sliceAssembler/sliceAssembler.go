package sliceAssembler

type SliceAssembler interface {
	NewRandomSlice(length, numFrom, numTo int) []int
	SliceExample(slice []int) []int
	AddElements(additionalNum int, slice []int) []int
	CopySlice(slice []int) []int
	RemoveElement(index int, slice []int) ([]int, error)
}
