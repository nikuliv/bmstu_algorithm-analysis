package conv

import (
	"time"

	"github.com/brianvoe/gofakeit"
)

// Conveyor used to generate credit cards in parallel conveyor way.
func Conveyor(n int, ch chan int) *Queue {
	first_belt := make(chan *Message, 5)
	second_belt := make(chan *Message, 5)
	third_belt := make(chan *Message, 5)
	l := createQueue(n)

	massage := func() {
		for {
			select {
			case cc := <-first_belt:
				cc.genMsg = true

				cc.startMsgGen = time.Now()
				cc.message = gofakeit.HackerPhrase()
				for i := 0; i < 100; i++ {
					cc.message = gofakeit.HackerPhrase()
				}
				cc.publicKey, cc.priv = PublicKey, PrivKey
				cc.doneMsgGen = time.Now()

				second_belt <- cc
			}
		}
	}
	encrypt := func() {
		for {
			select {
			case cc := <-second_belt:
				cc.encrypt = true

				cc.startEncrypt = time.Now()
				cc.encrypted, _ = Encrypt(cc.message, cc.publicKey)
				cc.doneEncrypt = time.Now()

				third_belt <- cc
			}
		}
	}

	decrypt := func() {
		for {
			select {
			case cc := <-third_belt:
				cc.decrypt = true

				cc.startDecrypt = time.Now()
				cc.decrypted, _ = Decrypt(cc.encrypted, cc.priv)
				cc.doneDecrypt = time.Now()

				l.push(cc)
				if cc.num == n {
					ch <- 0
				}

			}
		}
	}

	go massage()
	go encrypt()
	go decrypt()
	for i := 0; i <= n; i++ {
		cc := new(Message)
		cc.num = i
		first_belt <- cc
	}

	return l
}

func massage(cc *Message) {
	cc.genMsg = true

	cc.startMsgGen = time.Now()
	cc.message = gofakeit.HackerPhrase()
	for i := 0; i < 100; i++ {
		cc.message = gofakeit.HackerPhrase()
	}
	cc.publicKey, cc.priv = PublicKey, PrivKey
	cc.doneMsgGen = time.Now()
}

func encrypt(cc *Message) {
	cc.encrypt = true

	cc.startEncrypt = time.Now()
	cc.encrypted, _ = Encrypt(cc.message, cc.publicKey)
	cc.doneEncrypt = time.Now()
}

func decrypt(cc *Message) {
	cc.decrypt = true

	cc.startDecrypt = time.Now()
	cc.decrypted, _ = Decrypt(cc.encrypted, cc.priv)
	cc.doneDecrypt = time.Now()
}