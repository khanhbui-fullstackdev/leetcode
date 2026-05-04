package models

import "errors"

type QueueInt struct {
	Elements []int
	Count    int
}

func (q *QueueInt) Enqueue(v int) {
	q.Elements = append(q.Elements, v)
	q.Count = len(q.Elements)
}

func (q *QueueInt) Peek() int {
	firstItem := q.Elements[0]
	return firstItem
}

func (q *QueueInt) Dequeue() (int, error) {
	if len(q.Elements) == 0 {
		return 0, errors.New("queue is empty")
	}
	firstItem := q.Elements[0]
	// remove first item from slice
	q.Elements = q.Elements[1:]
	q.Count = len(q.Elements)
	return firstItem, nil
}
