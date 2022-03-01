package tests

import (
	"fmt"
	"lab_04/scripts"
	"lab_04/utils"
	"math/rand"
	"os"
	"text/tabwriter"
)

func TimeComparsion(iterationsNumber int) {
	var start, end float64
	tab := tabwriter.NewWriter(os.Stdout, 6, 24, 1, ' ', 0)
	fmt.Println("Random")
	tab.Write([]byte("Len\tSequential\tParallel\n"))

	for i := 100000; i <= 1000000; i += 100000 {
		arr1 := rand.Perm(i)
		//sort.Ints(arr1)
		temp := rand.Perm(i)
		copy(temp, arr1)
		fmt.Fprintf(tab, "%d\t", i)

		// Sequential
		start = utils.GetCPU()
		for j := 0; j < iterationsNumber; j++ {
			scripts.MergeSort(arr1)
		}
		end = utils.GetCPU()
		fmt.Fprintf(tab, "%v\t", (end - start) / float64(iterationsNumber))

		copy(arr1, temp)

		// Parallel
		start = utils.GetCPU()
		for j := 0; j < iterationsNumber; j++ {
			scripts.ParallelMergeSort(arr1)
		}
		end = utils.GetCPU()
		fmt.Fprintf(tab, "%v\t\n", (end - start) / float64(iterationsNumber))

		copy(arr1, temp)
	}
	tab.Flush()
}