package main

import (
	"fmt"
	"lab_04/scripts"
	"lab_04/tests"
	"lab_04/utils"
)

func main() {
	if (false) {
		fmt.Print("Input array size: ")
		n := utils.ReadNum()
		fmt.Print("Input array: ")
		arr := utils.ReadArray(n)
		marr := make([]int, n)
		parr := make([]int, n)

		fmt.Print("Merge sort: ")
		copy(marr, arr)
		scripts.MergeSort(marr)
		fmt.Println(marr)

		fmt.Print("Parallel Merge sort: ")
		copy(parr, arr)
		scripts.ParallelMergeSort(parr)
		fmt.Println(parr)
	}
	tests.TimeComparsion(100)
}