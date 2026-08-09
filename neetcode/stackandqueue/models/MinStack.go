package models

import "math"

type MinStack struct {
	Elements []int
	Size     int
	minValue int
}

func NewMinStack() MinStack {
	return MinStack{
		Elements: []int{},
		Size:     0,
		minValue: 0,
	}
}

// Pushes the element value onto the stack.
func (this *MinStack) Push(value int) {
	this.Elements = append(this.Elements, value)
	this.Size = len(this.Elements)
	if this.Size == 1 {
		this.minValue = value
	} else {
		currentMin := this.minValue
		if value < currentMin {
			currentMin = value
			this.minValue = currentMin
		}
	}
}

// Removes the element on the top of the stack.
func (this *MinStack) Pop() {
	if !this.isEmpty() {
		lastIndex := len(this.Elements) - 1
		lastElement := this.Elements[lastIndex]
		currentMin := this.minValue

		this.Elements = this.Elements[0:lastIndex]
		this.Size = len(this.Elements)

		if lastElement == currentMin {
			currentMin = math.MaxInt32 + 1
			// loop through all stack to find the min value
			for index := 0; index < this.Size; index++ {
				stackItem := this.Elements[index]
				if currentMin > stackItem {
					currentMin = stackItem
				}
			}
			this.minValue = currentMin
		}
	}
}

// Gets the top element of the stack.
func (this *MinStack) Top() int {
	if !this.isEmpty() {
		lastElement := this.Elements[len(this.Elements)-1]
		return lastElement
	}
	return 0
}

func (this *MinStack) GetMin() int {
	minVal := this.minValue
	return minVal
}

func (this *MinStack) isEmpty() bool {
	return this.Size == 0
}
