package scripts

import "fmt"

func ReadWord() string {
	var word string
	scanln, err := fmt.Scanln(&word)
	if err != nil || scanln == 0 {
		return ""
	}

	return word
}

type intMatrix [][]int

func (matrix intMatrix) PrintMatrix() {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			fmt.Printf("%3d ", matrix[i][j])
		}
		fmt.Printf("\n")
	}
}