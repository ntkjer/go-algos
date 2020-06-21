package queue

type node struct {
	Item interface{}
	next *node
}

type Queue struct {
	first *node
	last  *node
	N     int
}

func CopyQueue(toCopy *Queue) *Queue {
	var q *Queue = new(Queue)
	for x := toCopy.first; x != nil; x = x.next {
		q.Enqueue(x.Item)
	}
	return q
}

func (q *Queue) IsEmpty() bool {
	return q.N == 0
}

func (q *Queue) Size() int {
	return q.N
}

func (q *Queue) Peek() interface{} {
	return q.first.Item
}

func (q *Queue) Enqueue(val interface{}) {
	var last *node = new(node)
	var oldlast *node = new(node)
	oldlast = q.last
	last.Item = val
	last.next = nil
	q.last = last
	if q.IsEmpty() {
		q.first = last
	} else {
		oldlast.next = q.last
	}
	q.N++
}

func (q *Queue) Dequeue() interface{} {
	var oldfirst *node = new(node)
	oldfirst = q.first
	q.first = q.first.next
	if q.IsEmpty() {
		q.last = nil
	}
	q.N--
	return oldfirst.Item
}
