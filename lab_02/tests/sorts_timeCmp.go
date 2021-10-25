package tests

import (
	"fmt"
	"lab_02/scripts"
	"lab_02/tools"
	"math/rand"
	"os"
	"sort"
	"text/tabwriter"
)

func TimeComparsion(iterationsNumber int) {
	var start, end float64
	tab := tabwriter.NewWriter(os.Stdout, 4, 30, 1, ' ', 0)
	fmt.Println("Random")
	tab.Write([]byte("Len\tBubble\tSelection\tMerge\t\n"))

	for i := 10; i <= 100; i += 10 {
		arr1 := rand.Perm(i)
		//sort.Ints(arr1)
		temp := rand.Perm(i)
		copy(temp, arr1)
		fmt.Fprintf(tab, "%d\t", i)

		// Bubble
		start = tools.GetCPU()
		for j := 0; j < iterationsNumber; j++ {
			scripts.BubbleSort(arr1)
		}
		end = tools.GetCPU()
		fmt.Fprintf(tab, "%v\t", (end-start)/float64(iterationsNumber))

		copy(arr1, temp)

		// Selection
		start = tools.GetCPU()
		for j := 0; j < iterationsNumber; j++ {
			scripts.SelectionSort(arr1)
		}
		end = tools.GetCPU()
		fmt.Fprintf(tab, "%v\t", (end - start) / float64(iterationsNumber))

		copy(arr1, temp)

		// Merge
		start = tools.GetCPU()
		for j := 0; j < iterationsNumber; j++ {
			scripts.MergeSort(arr1)
		}
		end = tools.GetCPU()
		fmt.Fprintf(tab, "%v\t\n", (end - start) / float64(iterationsNumber))
	}
	tab.Flush()

	fmt.Println("Sorted")
	tab.Write([]byte("Len\tBubble\tSelection\tMerge\t\n"))

	for i := 100; i < 501; i += 100 {
		arr1 := rand.Perm(i)
		sort.Ints(arr1)
		temp := rand.Perm(i)
		copy(temp, arr1)
		fmt.Fprintf(tab, "%d\t", i)

		// Bubble
		start = tools.GetCPU()
		for j := 0; j < iterationsNumber; j++ {
			scripts.BubbleSort(arr1)
		}
		end = tools.GetCPU()
		fmt.Fprintf(tab, "%v\t", (end-start)/float64(iterationsNumber))

		copy(temp, arr1)

		// Selection
		start = tools.GetCPU()
		for j := 0; j < iterationsNumber; j++ {
			scripts.SelectionSort(arr1)
		}
		end = tools.GetCPU()
		fmt.Fprintf(tab, "%v\t", (end - start) / float64(iterationsNumber))

		copy(temp, arr1)

		// Merge
		start = tools.GetCPU()
		for j := 0; j < iterationsNumber; j++ {
			scripts.MergeSort(arr1)
		}
		end = tools.GetCPU()
		fmt.Fprintf(tab, "%v\t\n", (end - start) / float64(iterationsNumber))
	}
	tab.Flush()

	fmt.Println("Backsorted")
	tab.Write([]byte("Len\tBubble\tSelection\tMerge\t\n"))

	for i := 100; i <= 500; i += 100 {
		arr1 := rand.Perm(i)
		sort.Ints(arr1)
		tools.Reverse(arr1)
		temp := rand.Perm(i)
		copy(temp, arr1)
		fmt.Fprintf(tab, "%d\t", i)

		// Bubble
		start = tools.GetCPU()
		for j := 0; j < iterationsNumber; j++ {
			scripts.BubbleSort(arr1)
		}
		end = tools.GetCPU()
		fmt.Fprintf(tab, "%v\t", (end-start)/float64(iterationsNumber))

		copy(temp, arr1)

		// Selection
		start = tools.GetCPU()
		for j := 0; j < iterationsNumber; j++ {
			scripts.SelectionSort(arr1)
		}
		end = tools.GetCPU()
		fmt.Fprintf(tab, "%v\t", (end - start) / float64(iterationsNumber))

		copy(temp, arr1)

		// Merge
		start = tools.GetCPU()
		for j := 0; j < iterationsNumber; j++ {
			scripts.MergeSort(arr1)
		}
		end = tools.GetCPU()
		fmt.Fprintf(tab, "%v\t\n", (end - start) / float64(iterationsNumber))
	}
	tab.Flush()
}