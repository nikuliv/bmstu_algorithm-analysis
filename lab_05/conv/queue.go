package conv

func createQueue(amount int) *Queue {
	q := new(Queue)
	q.q = make([](*Message), amount, amount)
	q.l = -1

	return q
}

func (q *Queue) push(p *Message) {
	if q.l != len(q.q)-1 {
		q.q[q.l+1] = p
		q.l++
	}
}

func (q *Queue) pop() *Message {
	p := q.q[0]
	q.q = q.q[1:]
	q.l--

	return p
}
