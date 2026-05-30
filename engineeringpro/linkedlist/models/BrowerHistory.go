package models

import "fmt"

type BrowserHistory struct {
	First *BrowserNode
	Last  *BrowserNode
}

type BrowserNode struct {
	Page     string
	Selected bool
	Next     *BrowserNode
	Previous *BrowserNode
}

func NewBrowserHistory(homepage string) BrowserHistory {
	browserNode := &BrowserNode{
		Page:     homepage,
		Next:     nil,
		Previous: nil,
		Selected: true,
	}
	return BrowserHistory{
		First: browserNode,
		Last:  browserNode,
	}
}

func (this *BrowserHistory) Visit(url string) {
	newNode := &BrowserNode{
		Page:     url,
		Next:     nil,
		Previous: nil,
		Selected: true,
	}

	currentLast := this.Last
	if currentLast.Selected {
		// Append new node and assign it the last node
		currentLast.Next = newNode
		currentLast.Selected = false
		newNode.Previous = currentLast
	} else {
		selectedNode := this.findSelectedNode()
		selectedNode.Next = newNode
		selectedNode.Selected = false
		newNode.Previous = selectedNode
		currentLast = nil
	}
	this.Last = newNode
}

func (this *BrowserHistory) Back(steps int) string {
	backedPage := ""
	selectedNode := this.findSelectedNode()
	for selectedNode != nil {
		selectedNode.Selected = false
		if selectedNode == this.First {
			this.First.Selected = true
			backedPage = this.First.Page
			break
		}
		if steps == 0 {
			selectedNode.Selected = true
			backedPage = selectedNode.Page
			break
		}
		steps--
		selectedNode = selectedNode.Previous
	}
	return backedPage
}

func (this *BrowserHistory) Forward(steps int) string {
	forwardedPage := ""
	selectedNode := this.findSelectedNode()
	for selectedNode != nil {
		selectedNode.Selected = false
		if selectedNode == this.Last {
			this.Last.Selected = true
			forwardedPage = this.Last.Page
			break
		}
		if steps == 0 {
			selectedNode.Selected = true
			forwardedPage = selectedNode.Page
			break
		}
		steps--
		selectedNode = selectedNode.Next
	}
	return forwardedPage
}

func (this *BrowserHistory) PrintAllBrowserNodes() {
	fmt.Println("***** Print all nodes *****")
	current := this.First
	fmt.Printf("nil")
	for current != nil {
		template := fmt.Sprintf("<-(%s,selected:%v)->", current.Page, current.Selected)
		fmt.Printf("%s", template)
		current = current.Next
	}
	fmt.Printf("nil")
	fmt.Println()
}

func (this *BrowserHistory) findSelectedNode() *BrowserNode {
	currentNode := this.First
	for currentNode != nil {
		if currentNode.Selected {
			break
		}
		currentNode = currentNode.Next
	}
	return currentNode
}
