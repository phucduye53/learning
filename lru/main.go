package main

import "fmt"

//Node struct DLL
type Node struct {
	key        string
	value      int
	prev, next *Node
}

func newNode(k string, v int) *Node {
	n := &Node{k, v, nil, nil}
	return n
}

//DoubleLinkedList struct
//Store first and last data in DoubleLinkedList
type DoubleLinkedList struct {
	first, last *Node
}

func (l *DoubleLinkedList) isEmpty() bool {
	return l.last == nil
}

//Add a node to the first node of a DLL
func (l *DoubleLinkedList) addFirst(key string, value int) *Node {
	newNode := newNode(key, value)
	if l.first == nil && l.last == nil {

		l.first = newNode
		l.last = newNode
	} else {
		newNode.next = l.last
		l.first.prev = newNode
		l.first = newNode
	}
	return newNode
}

//Chance address of node to the first node in DLL
func (l *DoubleLinkedList) moveToFirst(node *Node) {
	if node == l.first {
		return
	}
	if node == l.last {
		l.last = l.last.prev
		l.last.next = nil
	} else {
		node.prev.next = node.next
		node.next.next = node.prev
	}

	node.next = l.first
	node.prev = nil
	l.first.prev = node
	l.first = node
}
func (l *DoubleLinkedList) removeLast() {
	if l.isEmpty() == true {
		return
	}
	if l.first == l.last {
		l.first = nil
		l.last = nil

	} else {
		l.last = l.last.prev
		l.last.next = nil
	}

}
func (l *DoubleLinkedList) getLast() *Node {
	return l.last
}

//LRU struct
type LRU struct {
	cap, size int
	list      *DoubleLinkedList
	m         map[string]*Node
}

func newLRU(cap int) *LRU {
	list := new(DoubleLinkedList)
	m := make(map[string]*Node)
	lru := &LRU{cap, 0, list, m}
	return lru
}

//Get the value of the key , if exist return -1
//else move key to the first of the DLL
func (lru *LRU) get(key string) int {
	var _, found = lru.m[key]
	if found == false {
		return -1
	}
	val := lru.m[key].value
	lru.list.moveToFirst(lru.m[key])
	return val

}

// Set, or insert the value if not present.
// If memory reach capacity,
//  remove the least recently used.
func (lru *LRU) put(key string, value int) {
	var _, found = lru.m[key]
	if found == true {
		lru.m[key].value = value
		lru.list.moveToFirst(lru.m[key])
		return
	}

	if lru.size == lru.cap {
		k := lru.list.getLast().key
		delete(lru.m, k)
		lru.list.removeLast()
		lru.size--
	}

	node := lru.list.addFirst(key, value)
	lru.size++
	lru.m[key] = node
}

func main() {
	lru := newLRU(2)
	lru.put("2", 2)
	fmt.Println(lru.get("2"))
	fmt.Println(lru.get("1"))
	lru.put("1", 1)
	lru.put("1", 5)
	fmt.Println(lru.get("1"))
	fmt.Println(lru.get("2"))
	lru.put("8", 8)
	fmt.Println(lru.get("1"))
	fmt.Println(lru.get("8"))
	fmt.Println(lru.get("5"))
	fmt.Println(lru.get("2"))
}
