package queues

// first-in, first-out, LIFO

import (
	"github.com/watson2396/algos-datastructs/datastructs/linkedlists"
)

type Queue struct{ linkedlists.SList }

// Enqueue
func (q *Queue) QueueEnqueue(name string) {
	q.SListAppend(name)
}

// Dequeue
func (q *Queue) QueueDequeue() string {
	return q.SListRemNext(nil)
}

// Peek
