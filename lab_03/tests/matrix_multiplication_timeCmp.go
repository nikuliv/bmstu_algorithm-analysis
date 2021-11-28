package tests

import (
	"fmt"
	"lab_02/scripts"
	"lab_02/tools"
	"os"
	"text/tabwriter"
)

func TimeComparsion() {
	var start, end float64
	tab := tabwriter.NewWriter(os.Stdout, 4, 30, 1, ' ', 0)
	fmt.Println("Random")
	tab.Write([]byte("Len\tBasic\tWinograd\tImproved Winograd\t\n"))

	for i := 100; i <= 1000; i += 100 {

		mat_1 := scripts.MakeResultMatrix(i, i)
		scripts.RandomFill(mat_1, i)

		mat_2 := scripts.MakeResultMatrix(i, i)
		scripts.RandomFill(mat_2, i)
		fmt.Fprintf(tab, "%d\t", i)

		// Basic
		start = tools.GetCPU()
		for j := 0; j < 10; j++ {
			scripts.BasicMatrixMultiplication(mat_1, mat_2)
		}
		end = tools.GetCPU()
		fmt.Fprintf(tab, "%v\t", end - start)

		// Winograd
		start = tools.GetCPU()
		for j := 0; j < 10; j++ {
			scripts.WinogradMatrixMultiplication(mat_1, mat_2)
		}
		end = tools.GetCPU()
		fmt.Fprintf(tab, "%v\t", end - start)

		// Improved Winograd
		start = tools.GetCPU()
		for j := 0; j < 10; j++ {
			scripts.WinogradMatrixMultiplicationImprove(mat_1, mat_2)
		}
		end = tools.GetCPU()
		fmt.Fprintf(tab, "%v\t\n", end - start)
	}
	tab.Flush()
}