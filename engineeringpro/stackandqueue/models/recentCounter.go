package models

type RecentCounter struct {
	QueueCounter Queue
}

func Constructor() RecentCounter {
	return RecentCounter{
		QueueCounter: Queue{},
	}
}

func (this *RecentCounter) Ping(t int) int {
	requestRange := t - 3000
	this.QueueCounter.Enqueue(t)

	for this.QueueCounter.Count > 0 && this.QueueCounter.Peek() < requestRange {
		this.QueueCounter.Dequeue()
	}

	return this.QueueCounter.Count
}
