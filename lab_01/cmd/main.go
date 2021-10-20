package main

import (
	"fmt"
	"lab_01/scripts"
	"lab_01/tests"
	"os"
	"strings"
)

func main() {
	if (strings.Compare(os.Args[1], "user")) == 0 {
		fmt.Println("Расстояние Левенштейна")

		fmt.Printf("Введите первое слово: ")
		s1 := scripts.ReadWord()
		fmt.Printf("Введите второе слово: ")
		s2 := scripts.ReadWord()

		fmt.Println()

		distRec := scripts.LevenshteinRecursive(s1, s2)
		fmt.Println("Рекурсивный способ нахождения растояния Левенштейна:", distRec)

		distRC := scripts.LevenshteinRecursiveCache(s1, s2)
		fmt.Println("Рекурсивный способ нахождения растояния Левенштейна с кэшированием:", distRC)

		distMt, mat := scripts.LevenshteinMatrix(s1, s2)
		fmt.Println("Матричный способ нахождения растояния Левенштейна:", distMt)
		x := mat[0][0]
		if x == 0 {
			x -= 1
		}
		//mat.PrintMatrix()

		distDL := scripts.DamerauLevenshtein(s1, s2)
		fmt.Println("Рекурсивный способ нахождения растояния Дамерау-Левенштейна:", distDL)
	} else {
		tests.TimeComparsion(100000)
	}

}
