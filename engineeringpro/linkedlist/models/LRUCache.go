package models

import "fmt"

type LRUCache struct {
	Head     *CacheNode
	Tail     *CacheNode
	Capacity int
	CacheMap map[int]*CacheNode
}

type CacheNode struct {
	Key      int
	Val      int
	Next     *CacheNode
	Previous *CacheNode
}

func Constructor(capacity int) LRUCache {
	lruCache := LRUCache{
		Capacity: capacity,
		CacheMap: make(map[int]*CacheNode, 2),
	}
	return lruCache
}

func (this *LRUCache) Get(key int) int {
	foundNode, hasExisted := this.CacheMap[key]
	if !hasExisted {
		return -1
	}
	head := this.DesignateNodeAsHead(foundNode)
	return head.Val
}

func (this *LRUCache) Put(key int, value int) {
	foundNode, hasExisted := this.CacheMap[key]
	if hasExisted {
		// update node value
		foundNode.Val = value
		this.CacheMap[key] = foundNode
		this.DesignateNodeAsHead(foundNode)
	} else {
		// get number of keys
		numberOfKeys := len(this.CacheMap)
		// check whether length exceeds capacity
		if numberOfKeys == this.Capacity {
			// we have to evict the least recently used key.
			// the least recently used key means that we remove tails
			this.RemoveTail()
		}
		//  add the key-value pair to the cache
		// 	the newNode will be the head
		this.AddFirst(key, value)
	}
}

// Designate a specific node as head
// For example: head [1,3,17,8] -> DesignateNodeAsHead("node#17") -> head [17,1,3,8]
// head [1,3,17,8] -> DesignateNodeAsHead("node#8") -> head [8,1,3,17]
func (this *LRUCache) DesignateNodeAsHead(designatedNode *CacheNode) *CacheNode {
	switch designatedNode {
	case this.Head:
	case this.Tail:
		prevFoundNode := designatedNode.Previous
		prevFoundNode.Next = nil
		this.Tail = prevFoundNode

		currentHead := this.Head
		currentHead.Previous = designatedNode
		designatedNode.Next = currentHead
		designatedNode.Previous = nil
		this.Head = designatedNode

	default:
		prevFoundNode := designatedNode.Previous // node3
		nextFoundNode := designatedNode.Next     // node8

		prevFoundNode.Next = nextFoundNode     // node3.Next = node8
		nextFoundNode.Previous = prevFoundNode // node8.Prev = node3

		currentHead := this.Head // node1

		designatedNode.Previous = nil     // node17.Prev = nil
		designatedNode.Next = currentHead // node17.Next = node1
		this.Head = designatedNode
	}
	return designatedNode
}

// Evict the least recently used key
// head [1,3,17,8] -> RemoveTail() -> head [1,3,17]
func (this *LRUCache) RemoveTail() {
	currentTail := this.Tail
	// remove key
	delete(this.CacheMap, currentTail.Key)

	previousTail := currentTail.Previous
	previousTail.Next = nil
	this.Tail = previousTail
	currentTail = nil
}

func (this *LRUCache) PrintAllListNode() {
	node := this.Head
	fmt.Printf("nil")
	for node != nil {
		fmt.Printf("<-%d->", node.Val)
		node = node.Next
	}
	fmt.Printf("nil")
	fmt.Printf(" || Length:%d", len(this.CacheMap))
	fmt.Println()
}

func (this *LRUCache) PrintAllCacheMap() {
	fmt.Println("Map:", this.CacheMap)
}

// head = [] -> AddFirst(3,2) => head = [node(3,2)] -> AddFirst(7,90) =>  head = [node(7,90), node(3,2)] -> AddFirst(3,3) =>  head = [node(3,3),node(7,90), node(3,2)]
func (this *LRUCache) AddFirst(key int, value int) {
	newNode := &CacheNode{
		Key:      key,
		Val:      value,
		Previous: nil,
	}
	if len(this.CacheMap) == 0 {
		this.Head = newNode
		this.Tail = newNode
		this.CacheMap[key] = newNode
		return
	}
	currentNode := this.Head
	currentNode.Previous = newNode
	newNode.Next = currentNode
	this.Head = newNode
	this.CacheMap[key] = newNode
}
