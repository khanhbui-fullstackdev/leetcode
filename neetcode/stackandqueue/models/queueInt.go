package models

import "errors"

type QueueInt struct {
	Elements []int
	Size     int
}

// queue.Enqueue(1) -> queue = [1]
// queue.Enqueue(2) -> queue = [1,2]
// queue.Enqueue(3) -> queue = [1,2,3]
func (q *QueueInt) Enqueue(v int) {
	q.Elements = append(q.Elements, v)
	q.Size = len(q.Elements)
}

// First int first out
// queue = [1,2,3]
// queue.Dequeue() ->  queue = [2,3]
func (q *QueueInt) Dequeue() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("Queue is empty")
	}
	element := q.Elements[0]
	q.Elements = q.Elements[1:]
	q.Size = len(q.Elements)
	return element, nil
}

func (q *QueueInt) Peek() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("Queue is empty")
	}
	return q.Elements[0], nil
}

func (q *QueueInt) IsEmpty() bool {
	if q.Size == 0 {
		return true
	}
	return false
}
