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
		CacheMap: make(map[int]*CacheNode, capacity),
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
		fmt.Printf("\nNode(%d) has existed in cached. Update new value:%d", key, value)
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
		prevFoundNode := designatedNode.Previous // prevFoundNode = node8.Previous = node17
		prevFoundNode.Next = nil                 // node17.Next = nil
		this.Tail = prevFoundNode                // this.Tail = node17

		currentHead := this.Head              // currentHead = node1
		currentHead.Previous = designatedNode // node1.Previous = node8
		designatedNode.Next = currentHead     // node8.Next= node1
		designatedNode.Previous = nil         // node8.Previous = nil
		this.Head = designatedNode            // this.Head = node8

	default:
		prevFoundNode := designatedNode.Previous // node3
		nextFoundNode := designatedNode.Next     // node8

		prevFoundNode.Next = nextFoundNode     // node3.Next = node8
		nextFoundNode.Previous = prevFoundNode // node8.Prev = node3

		currentHead := this.Head // node1

		designatedNode.Previous = nil     // node17.Prev = nil
		designatedNode.Next = currentHead // node17.Next = node1
		this.Head = designatedNode        // this.Head = node17
		currentHead.Previous = this.Head  // node1.Previous = node17
	}

	return designatedNode
}

// Evict the least recently used key
// head [1,3,17,8] -> RemoveTail() -> head [1,3,17]
func (this *LRUCache) RemoveTail() {
	currentTail := this.Tail
	fmt.Printf("\nEvict the least recently used key:%v", currentTail)

	// remove key
	delete(this.CacheMap, currentTail.Key)

	if currentTail == this.Head && currentTail == this.Tail {
		this.Head = nil
		this.Tail = nil
		currentTail = nil
	} else {
		previousTail := currentTail.Previous
		previousTail.Next = nil
		this.Tail = previousTail
		currentTail = nil
	}
}

func (this *LRUCache) PrintAllListNode() {
	fmt.Println("Print all list nodes")
	node := this.Head
	fmt.Printf("nil")
	for node != nil {
		fmt.Printf("<-[%d,%d]->", node.Key, node.Val)
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
		fmt.Printf("\nFirst new node designate as head node:%v", newNode)
		return
	}
	currentNode := this.Head
	currentNode.Previous = newNode
	newNode.Next = currentNode
	this.Head = newNode
	this.CacheMap[key] = newNode

	fmt.Printf("\nAdd new node and designate as head node:%v", newNode)
}
