package main

import (
	"fmt"
	"lab5/conv"
	"time"

	aurora "github.com/logrusorgru/aurora"
)

func main() {
	fmt.Printf("%v", "Параллельные конвейерные вычисления\n\n")

	n := 5

	start := time.Now()
	ch := make(chan int)
	lineP := conv.Conveyor(n, ch)
	<-ch
	conv.Log(lineP, true)
	_ = lineP
	end := time.Now()
	fmt.Println(
		aurora.Sprintf(
			("\nКодирование и декодирование %d сообщений заняло %s\n"),
			n,
			end.Sub(start),
		))
}
