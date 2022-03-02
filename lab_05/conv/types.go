package conv

import (
	"time"
)

type Message struct {
	num				int
	genMsg  		bool
	encrypt    	 	bool
	decrypt    		bool
	startMsgGen 	time.Time
	doneMsgGen  	time.Time
	startEncrypt    time.Time
	doneEncrypt     time.Time
	startDecrypt    time.Time
	doneDecrypt     time.Time

	message 	string
	publicKey	string
	priv		string
	encrypted   string
	decrypted   string
}

type Queue struct {
	q [](*Message)
	l int
}