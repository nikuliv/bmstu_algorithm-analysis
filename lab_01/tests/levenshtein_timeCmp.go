package tests

import (
	"fmt"
	"lab_01/scripts"
	"lab_01/tools"
	"os"
	"text/tabwriter"
)

func TimeComparsion(iterationsNumber int) {
	var start, end float64
	tab := tabwriter.NewWriter(os.Stdout, 4, 30, 1, ' ', 0)
	tab.Write([]byte("Len\tRecursion\tMatrix\tWith_cache\tDamerau\t\n"))

	for i := 5; i <= 9; i++ {
		string1 := tools.RandStringBytesMask(i)
		string2 := tools.RandStringBytesMask(i)
		fmt.Fprintf(tab, "%d\t", i)

		// Recursive
		if i <= 10 {
			start = tools.GetCPU()
			for j := 0; j < iterationsNumber/100; j++ {
				scripts.LevenshteinRecursive(string1, string2)
			}
			end = tools.GetCPU()
			fmt.Fprintf(tab, "%v\t", (end-start)/float64(iterationsNumber/100))
		} else {
			fmt.Fprintf(tab, "%s\t", "-")
		}

		// Matrix
		start = tools.GetCPU()
		for j := 0; j < iterationsNumber; j++ {
			scripts.LevenshteinMatrix(string1, string2)
		}
		end = tools.GetCPU()
		fmt.Fprintf(tab, "%v\t", (end - start) / float64(iterationsNumber))

		// With_cache
		start = tools.GetCPU()
		for j := 0; j < iterationsNumber; j++ {
			scripts.LevenshteinRecursiveCache(string1, string2)
		}
		end = tools.GetCPU()
		fmt.Fprintf(tab, "%v\t", (end - start) / float64(iterationsNumber))

		// Damerau
		if i <= 10 {
			start = tools.GetCPU()
			for j := 0; j < iterationsNumber/100; j++ {
				scripts.DamerauLevenshtein(string1, string2)
			}
			end = tools.GetCPU()
			fmt.Fprintf(tab, "%v\t\n", (end-start)/float64(iterationsNumber/100))
		} else {
			fmt.Fprintf(tab, "%s\t\n", "-")
		}
	}

	for i := 10; i < 101; i+= 10 {
		string1 := tools.RandStringBytesMask(i)
		string2 := tools.RandStringBytesMask(i)
		fmt.Fprintf(tab, "%d\t", i)

		// Recursive
		fmt.Fprintf(tab, "%s\t", "-")

		// Matrix
		start = tools.GetCPU()
		for j := 0; j < iterationsNumber; j++ {
			scripts.LevenshteinMatrix(string1, string2)
		}
		end = tools.GetCPU()
		fmt.Fprintf(tab, "%v\t", (end-start)/float64(iterationsNumber))

		// With_cache
		start = tools.GetCPU()
		for j := 0; j < iterationsNumber; j++ {
			scripts.LevenshteinRecursiveCache(string1, string2)
		}
		end = tools.GetCPU()
		fmt.Fprintf(tab, "%v\t", (end-start)/float64(iterationsNumber))

		// Damerau
		fmt.Fprintf(tab, "%s\t\n", "-")
	}
	fmt.Println("Time in milliseconds")
	tab.Flush()
}