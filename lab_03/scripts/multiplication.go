package scripts

func BasicMatrixMultiplication(mat_1, mat_2 intMatrix) intMatrix {
	res := MakeResultMatrix(mat_1.n, mat_2.m)

	for i := 0; i < res.n; i++ {
		for j := 0; j < res.m; j++ {
			res.mat[i][j] = 0
			for k := 0; k < mat_1.m; k++ {
				res.mat[i][j] += mat_1.mat[i][k] * mat_2.mat[k][j]
			}
		}
	}
	return res
}


func WinogradMatrixMultiplication(mat_1, mat_2 intMatrix) intMatrix {
	res := MakeResultMatrix(mat_1.n, mat_2.m)

	rowf := make([]int, mat_1.n)
	for i := 0; i < mat_1.n; i++ {
		for j := 0; j < mat_1.m/2; j++ {
			rowf[i] += mat_1.mat[i][j*2] * mat_1.mat[i][j*2+1]
		}
	}

	colf := make ([]int, mat_2.m)
	for i := 0; i < mat_2.m; i++ {
		for j := 0; j < mat_2.n/2; j++ {
			colf[i] += mat_2.mat[j*2][i] * mat_2.mat[j*2+1][i]
		}
	}

	for i := 0; i < res.n; i++ {
		for j := 0; j < res.m; j++ {
			res.mat[i][j] = - rowf[i] - colf[j]
			for k := 0; k < mat_2.n/2; k++ {
				res.mat[i][j] += (mat_1.mat[i][k*2+1]+mat_2.mat[k*2][j])*(mat_1.mat[i][k*2]+mat_2.mat[k*2+1][j])
			}
		}
	}

	if mat_2.n%2 != 0 {
		for i := 0; i < res.n; i++ {
			for j := 0; j < res.m; j++ {
				res.mat[i][j] += mat_1.mat[i][mat_1.m-1] * mat_2.mat[mat_2.n-1][j]
			}
		}
	}

	return res
}

// WinogradMultImp used to multiply two matrixes using impproved Winograd method.
func WinogradMatrixMultiplicationImprove(mat_1, mat_2 intMatrix) intMatrix {
	res := MakeResultMatrix(mat_1.n, mat_2.m)

	rowf := make([]int, mat_1.n)
	for i := 0; i < mat_1.n; i++ {
		for j := 0; j < mat_1.m/2; j++ {
			rowf[i] += mat_1.mat[i][j*2] * mat_1.mat[i][j*2+1]
		}
	}

	colf := make ([]int, mat_2.m)
	for i := 0; i < mat_2.m; i++ {
		for j := 0; j < mat_2.n/2; j++ {
			colf[i] += mat_2.mat[j*2][i] * mat_2.mat[j*2+1][i]
		}
	}

	flag := false
	if mat_2.n % 2 == 1 {
		flag = true
	}

	for i := 0; i < res.n; i++ {
		for j := 0; j < res.m; j++ {
			res.mat[i][j] = -rowf[i] - colf[j]
			for k := 0; k < mat_2.n - 1; k += 2 {
				res.mat[i][j] += (mat_1.mat[i][k] + mat_2.mat[k+1][j]) * (mat_1.mat[i][k+1] + mat_2.mat[k][j])
			}
			if (flag) {
				res.mat[i][j] += mat_1.mat[i][mat_2.n - 1] * mat_2.mat[mat_2.n - 1][j]
			}
		}
	}

	return res
}
