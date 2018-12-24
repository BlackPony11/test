package main

import (
	"fmt"
	"strings"
)

//ReturnInt comment
func ReturnInt() int {
	i := 1
	return i
}

//ReturnFloat comment
func ReturnFloat() float32 {
	var f float32 = 1.1
	return f
}

//ReturnIntArray com
func ReturnIntArray() [3]int {
	arr := [3]int{1, 3, 4}
	return arr
}

//ReturnIntSlice com
func ReturnIntSlice() []int {
	sl := []int{1, 2, 3}
	return sl
}

//IntSliceToString com
func IntSliceToString(y []int) string {
	return strings.Trim(strings.Replace(fmt.Sprint(y), " ", "", -1), "[]")

}

//MergeSlices com
func MergeSlices(msl []float32, msl2 []int32) []int {
	var merged []int
	for _, value := range msl {
		merged = append(merged, int(value))
	}
	for _, value := range msl2 {
		merged = append(merged, int(value))
	}
	return merged
}

func GetMapValuesSortedByKey(input map[int]string) []string {
	var result []string

	for len(input) > 0 {
		key := getMaxKeyFromMap(input)
		result = append(result, input[key])
		delete(input, key)
	}

	return flipSlice(result)
}

func getMaxKeyFromMap(input map[int]string) int {
	var max = 0
	for key := range input {
		if key > max {
			max = key
		}
	}
	return max
}

func flipSlice(input []string) []string {
	var result []string
	for i := len(input) - 1; i >= 0; i-- {
		result = append(result, input[i])
	}
	return result
}

func main() {

}
