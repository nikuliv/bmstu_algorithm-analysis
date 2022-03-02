package conv

import (
	"fmt"
	"strings"
)

// Log used to log conveyor tasks time.
func Log(q *Queue, logCreditCard bool) {

	l := q.q
	start := l[0].startMsgGen
	fmt.Printf("%v%43v%v\n", "+", strings.Repeat("-", 42), "+")
	fmt.Printf(
		"|%5s|%5s|%15s|%15s|\n",
		"Tape", "Task", "Start", "End",
	)
	fmt.Printf("%v%43v%v\n", "+", strings.Repeat("-", 42), "+")
	for i := 0; i < len(l); i++ {
		if l[i] != nil {
			fmt.Printf(
				"|%5d|%5v|%15v|%15v|\n",
				1, i + 1, l[i].startMsgGen.Sub(start), l[i].doneMsgGen.Sub(start),
			)
			fmt.Printf(
				"|%5d|%5v|%15v|%15v|\n",
				2, i + 1, l[i].startEncrypt.Sub(start), l[i].doneEncrypt.Sub(start),
			)
			fmt.Printf(
				"|%5d|%5v|%15v|%15v|\n",
				3, i + 1, l[i].startDecrypt.Sub(start), l[i].doneDecrypt.Sub(start),
			)
			fmt.Printf("%v%43v%v\n", "+", strings.Repeat("-", 42), "+")
		}
	}
	if logCreditCard {
		fmt.Printf("\n%v:\n", "Generated data")
		for i := range l {
			fmt.Printf("CreditCard: %d\nMassage: %s\nEncrypted: %0x\nDecrypted: %s\n\n", l[i].num+1, l[i].message, l[i].encrypted, l[i].decrypted)
		}
	}
}
