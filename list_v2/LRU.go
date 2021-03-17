/**
lru 实现
*/
package list_v2

import (
	"errors"
	"fmt"
)

type LRU struct {
	head    *lruNode
	tail    *lruNode
	nodeMap map[string]*lruNode
	size    int
	cap     int
}

func NewLRU(cap int) (*LRU, error) {
	if cap <= 0 {
		return nil, errors.New("cap must be greater than 0! ")
	}
	return &LRU{
		head:    nil,
		tail:    nil,
		nodeMap: make(map[string]*lruNode, cap),
		size:    0,
		cap:     cap,
	}, nil
}
func (l *LRU) PrintAll() {
	t := l.head
	for t != nil {
		fmt.Print("[", t.key, "=", t.val, "] ")
		t = t.next
	}
	fmt.Println()
}

func (l *LRU) Set(key string, val int) bool {
	if _, has := l.nodeMap[key]; has {
		return false
	}
	newNode := &lruNode{
		key: key,
		val: val,
	}
	if l.size == 0 {
		l.head, l.tail = newNode, newNode
		l.size++
	} else if l.size == l.cap { //size超过cap，移除队尾
		l.insertHead(newNode)
		removeNode := l.tail
		l.tail = l.tail.prev
		removeNode.remove()
		delete(l.nodeMap, key)
	} else {
		l.insertHead(newNode)
		l.size++
	}
	l.nodeMap[key] = newNode
	return true
}

func (l *LRU) Get(key string) (int, bool) {
	if node, has := l.nodeMap[key]; has {
		//当前节点插入到队头
		node.remove()
		l.insertHead(node)
		return node.val, true
	} else {
		return -1, false
	}
}

func (l *LRU) insertHead(node *lruNode) {
	l.head.prev, node.next = node, l.head
	l.head = node
}

/**
双链表
*/
type lruNode struct {
	key  string
	val  int
	next *lruNode
	prev *lruNode
}

/**
当前节点脱离双链表
*/
func (n *lruNode) remove() *lruNode {
	if n.next != nil {
		n.next.prev = n.prev
	}
	if n.prev != nil {
		n.prev.next = n.next
	}
	n.prev, n.next = nil, nil
	return n
}


//func main() {
//	lru, _ := list_v2.NewLRU(5)
//	lru.Set("a1", 1)
//	lru.Set("a2", 1)
//	lru.Set("a3", 1)
//	lru.Set("a4", 1)
//	lru.Set("a5", 1)
//	lru.PrintAll()
//	lru.Get("a3")
//	lru.Get("a4")
//	lru.Set("a6", 1)
//	lru.PrintAll()
//}