package main

import (
	"fmt"
	"lab_02/scripts"
	"lab_02/tests"
)

func main() {
	test := false
	if test {
		tests.TimeComparsion()
	} else {
		fmt.Println("Умножение матриц\n")
		fmt.Println("Матрица 1")
		fmt.Print("Введите количество строк   : ")
		an := scripts.ReadNum()
		fmt.Print("Введите количество столбцов: ")
		am := scripts.ReadNum()
		fmt.Println("Введите матрицу:")
		amat := scripts.ReadMatrix(an, am)

		fmt.Println("\nМатрица 2")
		fmt.Print("Введите количество строк   : ")
		bn := scripts.ReadNum()

		if bn != am {
			fmt.Println("Количество строк матрицы 1 B не равно количеству" +
				" столбцов матрицы 2. Умножение невозможно.")
			return
		}

		fmt.Print("Введите количество столбцов: ")
		bm := scripts.ReadNum()
		fmt.Println("Введите матрицу:")
		bmat := scripts.ReadMatrix(bn, bm)

		fmt.Println("\nРезультирующая матрица")
		fmt.Println("Обычное умножение:")
		smmat := scripts.BasicMatrixMultiplication(amat, bmat)
		smmat.PrintMatrix()
		fmt.Println("Умножение по алгоритму Винограда:")
		wmmat := scripts.WinogradMatrixMultiplication(amat, bmat)
		wmmat.PrintMatrix()
		fmt.Println("Умножение по улучшенному алгоритму Винограда:")
		wimmat := scripts.WinogradMatrixMultiplicationImprove(amat, bmat)
		wimmat.PrintMatrix()
	}
}
