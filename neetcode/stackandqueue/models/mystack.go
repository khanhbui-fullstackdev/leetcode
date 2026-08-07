package models

type MyStack struct {
	Queue1 QueueInt
	Queue2 QueueInt
}

func Constructor() MyStack {
	return MyStack{
		Queue1: QueueInt{},
		Queue2: QueueInt{},
	}
}

// Time complexity = O(n)
func (this *MyStack) Push(x int) {
	if this.Queue1.IsEmpty() {
		this.Queue1.Enqueue(x)

		for this.Queue2.Size > 0 {
			item, err := this.Queue2.Dequeue()
			if err == nil {
				this.Queue1.Enqueue(item)
			}
		}

	} else if this.Queue2.IsEmpty() {
		this.Queue2.Enqueue(x)
		for this.Queue1.Size > 0 {
			item, err := this.Queue1.Dequeue()
			if err == nil {
				this.Queue2.Enqueue(item)
			}
		}
	}
}

// Time complexity O(1)
func (this *MyStack) Pop() int {
	if this.Queue1.Size > 0 {
		item, _ := this.Queue1.Dequeue()
		return item
	}
	item, _ := this.Queue2.Dequeue()
	return item
}

// Time complexity O(1)
func (this *MyStack) Top() int {
	if this.Queue1.Size > 0 {
		return this.Queue1.Elements[0]
	}
	return this.Queue2.Elements[0]
}

func (this *MyStack) Empty() bool {
	if this.Queue1.Size == 0 && this.Queue2.Size == 0 {
		return true
	}
	return false
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
