package models

import "fmt"

type LruCache struct {
	Head     *CacheNode
	Tail     *CacheNode
	NodeMap  map[int]*CacheNode
	Capacity int
	Count    int
}

type CacheNode struct {
	Val      int
	Key      int
	Next     *CacheNode
	Previous *CacheNode
}

func Constructor(capacity int) LruCache {
	return LruCache{
		Capacity: capacity,
		NodeMap:  make(map[int]*CacheNode, capacity),
	}
}

func (this *LruCache) Get(key int) int {
	foundNode, existed := this.NodeMap[key]
	if existed {
		switch foundNode {
		case this.Head:
		case this.Tail:
			this.AssignANodeAsTail(false)
			this.AssignANodeAsHead(foundNode)
		default:
			nextFoundNode := foundNode.Next
			previousFoundNode := foundNode.Previous

			// create new chain between previousFoundNode & nextFoundNode
			previousFoundNode.Next = nextFoundNode
			nextFoundNode.Previous = previousFoundNode

			this.AssignANodeAsHead(foundNode)
		}
		return foundNode.Val
	}
	return -1
}

func (this *LruCache) Put(key int, value int) {
	foundNode, existed := this.NodeMap[key]
	if existed {
		// check if the foundnode is head or tail or nornal node
		switch foundNode {
		case this.Head:
		case this.Tail:
			this.AssignANodeAsTail(false)
			this.AssignANodeAsHead(foundNode)
		default:
			nextFoundNode := foundNode.Next
			previousFoundNode := foundNode.Previous
			// create new chain between previousFoundNode & nextFoundNode
			previousFoundNode.Next = nextFoundNode
			nextFoundNode.Previous = previousFoundNode

			this.AssignANodeAsHead(foundNode)
		}

		// update new value to existing node
		foundNode.Val = value
		this.NodeMap[key] = foundNode
		this.Head = foundNode
	} else {
		// add new node into node list
		newNode := &CacheNode{Key: key, Val: value}
		// check if the head is nill and tail is nill
		// newNode will be the head node
		if this.Count == 0 || (this.Head == nil && this.Tail == nil) {
			this.Head = newNode
			this.Tail = newNode
			this.Count++
			this.NodeMap[newNode.Key] = newNode
			return
		}

		// If the number of keys exceeds the capacity
		if this.Capacity == len(this.NodeMap) {
			// edge case when node is tail and head (capacity = 1)
			if this.Tail == this.Head {
				delete(this.NodeMap, this.Head.Key)
				this.Head = newNode
				this.Tail = newNode
				this.NodeMap[newNode.Key] = newNode
				return
			}

			// evict the least recently used -> the tail
			this.AssignANodeAsTail(true)
			this.Count--
		}
		this.AssignANodeAsHead(newNode)
		this.Count++
	}
}

func (this *LruCache) PrintAllNodes() {
	fmt.Printf("nil")
	currentHead := this.Head
	for currentHead != nil {
		fmt.Printf("<->node(%d,%d)", currentHead.Key, currentHead.Val)
		currentHead = currentHead.Next
	}
	fmt.Printf("<-> nil | count:%d", this.Count)
	fmt.Println()
}

func (this *LruCache) PrintNodeMap() {
	fmt.Printf("Node map:[")
	for _, node := range this.NodeMap {
		fmt.Printf("node(%d,%d) ", node.Key, node.Val)
	}
	fmt.Printf("]")
	fmt.Println()
}

func (this *LruCache) AssignANodeAsHead(node *CacheNode) {
	currentHead := this.Head
	currentHead.Previous = node
	node.Previous = nil
	node.Next = currentHead
	this.Head = node
	this.NodeMap[node.Key] = node
}

func (this *LruCache) AssignANodeAsTail(isEvicted bool) {
	currentTail := this.Tail
	previousTail := currentTail.Previous
	previousTail.Next = nil
	this.Tail = previousTail
	if isEvicted {
		delete(this.NodeMap, currentTail.Key)
	}
}
