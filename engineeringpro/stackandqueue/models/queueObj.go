package models

import "errors"

type QueueObj struct {
	Elements []any
	Count    int
}

func (q *QueueObj) EnqueueObj(v any) {
	q.Elements = append(q.Elements, v)
	q.Count = len(q.Elements)
}

func (q *QueueObj) DequeueObj() (any, error) {
	if len(q.Elements) == 0 {
		return nil, errors.New("queue is empty")
	}
	firstItem := q.Elements[0]
	// Tối ưu bộ nhớ: Xóa tham chiếu tới phần tử cũ
	q.Elements[0] = nil

	// remove first item from slice
	q.Elements = q.Elements[1:]
	q.Count = len(q.Elements)
	return firstItem, nil
}

func (q *QueueObj) PeekObj() any {
	firstItem := q.Elements[0]
	return firstItem
}

func (q *QueueObj) IsEmpty() bool {
	return q.Count == 0
}
