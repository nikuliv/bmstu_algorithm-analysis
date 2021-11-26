package scripts

import (
	"fmt"
	"math/rand"
)

type intMatrix struct {
	mat [][]int
	n, m int
}

func ReadMatrix(n, m int) intMatrix {
	mat := MakeResultMatrix(n, m)

	for i := 0; i < mat.n; i++ {
		for j := 0; j < mat.m; j++ {
			fmt.Scanf("%d", &mat.mat[i][j])
		}
	}

	return mat
}

func (matrix intMatrix) PrintMatrix() {
	for i := 0; i < matrix.n; i++ {
		for j := 0; j < matrix.m; j++ {
			fmt.Printf("%3d ", matrix.mat[i][j])
		}
		fmt.Printf("\n")
	}
}

func ReadNum() int {
	var num int

	fmt.Scanln(&num)

	return num
}

func RandomFill(mat intMatrix, max int) {
	for i := 0; i < mat.n; i++ {
		for j := 0; j < mat.m; j++ {
			mat.mat[i][j] = rand.Intn(max)
		}
	}
}

func MakeResultMatrix(n, m int) intMatrix {
	var rmat intMatrix

	rmat.n = n
	rmat.m = m
	rmat.mat = make([][]int, rmat.n)

	for i := range rmat.mat {
		rmat.mat[i] = make([]int, rmat.m)
	}

	return rmat
}
