package models

type MyQueue struct {
	Stack1 StackInt
	Stack2 StackInt
}

func (this *MyQueue) Constructor() MyQueue {
	return MyQueue{
		Stack1: StackInt{},
		Stack2: StackInt{},
	}
}

// stack1 is the place for Push operation
// Always push item into stack1
func (this *MyQueue) Push(x int) {
	this.Stack1.Push(x)
}

// If stack2 is empty, we get all elements from stack1 and push all of them to stack2
// If stack2 is not empty. We always Pop()/Peek from stack2
func (this *MyQueue) Pop() int {
	if this.Stack2.IsEmpty() {
		for this.Stack1.Size > 0 {
			item, _ := this.Stack1.Pop()
			this.Stack2.Push(item)
		}
	}
	item, _ := this.Stack2.Pop()
	return item
}

func (this *MyQueue) Peek() int {
	if this.Stack2.IsEmpty() {
		for this.Stack1.Size > 0 {
			item, _ := this.Stack1.Pop()
			this.Stack2.Push(item)
		}
	}
	item, _ := this.Stack2.Peek()
	return item
}

func (this *MyQueue) Empty() bool {
	return this.Stack1.IsEmpty() && this.Stack2.IsEmpty()
}
