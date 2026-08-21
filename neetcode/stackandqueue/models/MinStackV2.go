package models

import "fmt"

type MinStackV2 struct {
	elements    []int // main stack
	minElements []int // sub stack contains min value

	sizeMainStack int
	sizeMinStack  int
}

func NewMinStackV2() MinStackV2 {
	return MinStackV2{}
}

/*
stack.push (5)
elements = [5]
minElements = [5]

stack.push (3)
elements = [5,3]
minElements = [5,3]

stack.push (4)
elements = [5,3,4]
minElements = [5,3]

stack.push (2)
elements = [5,3,4,2]
minElements = [5,3,2]
*/
func (this *MinStackV2) Push(value int) {
	fmt.Println("Push:", value)
	this.elements = append(this.elements, value)
	if this.sizeMinStack == 0 {
		this.minElements = append(this.minElements, value)
	} else {
		lastMinElement := this.minElements[this.sizeMinStack-1]
		if lastMinElement >= value {
			this.minElements = append(this.minElements, value)
		}
	}
	this.sizeMainStack = len(this.elements)
	this.sizeMinStack = len(this.minElements)
}

/*
elements = [5,3,4,2]
minElements = [5,3,2]

stack.pop()
elements = [5,3,4]
minElements = [5,3]
*/
func (this *MinStackV2) Pop() {
	if this.sizeMainStack > 0 {
		lastElement := this.elements[this.sizeMainStack-1]
		if this.sizeMinStack > 0 {
			lastMinElement := this.minElements[this.sizeMinStack-1]
			if lastElement == lastMinElement {
				this.minElements = this.minElements[0 : this.sizeMinStack-1]
			}
		}
		this.elements = this.elements[0 : this.sizeMainStack-1]
	}

	this.sizeMainStack = len(this.elements)
	this.sizeMinStack = len(this.minElements)
}

func (this *MinStackV2) Top() int {
	if this.sizeMainStack > 0 {
		return this.elements[this.sizeMainStack-1]
	}
	return 0
}

func (this *MinStackV2) GetMin() int {
	if this.sizeMainStack > 0 {
		if this.sizeMinStack > 0 {
			return this.minElements[this.sizeMinStack-1]
		}
		return this.elements[this.sizeMainStack-1]
	}
	return 0
}

func (this *MinStackV2) PrintAllStackVals() {
	fmt.Println("Stack vals:", this.elements)
}

func (this *MinStackV2) PrintAllMinStackVals() {
	fmt.Println("Min stack vals:", this.minElements)
}
