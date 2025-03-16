package datastructs

// first-in, first-out, LIFO

type Queue struct{ List }

// Enqueue
func (q *Queue) QueueEnqueue(name string) {
	q.ListAppend(name)
}

// Dequeue
func (q *Queue) QueueDequeue() string {
	return q.ListRemNext(nil)
}

// Peek
